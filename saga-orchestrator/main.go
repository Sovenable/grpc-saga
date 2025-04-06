package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderPb "github.com/Sovenable/grpc-saga/proto/order"
	paymentPb "github.com/Sovenable/grpc-saga/proto/payment"
	shippingPb "github.com/Sovenable/grpc-saga/proto/shipping"
)

// sagaExecute adalah fungsi utama yang mengatur alur transaksi
func sagaExecute(orderID string) {
	// Membuat koneksi ke masing-masing service
	connOrder, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}
	defer connOrder.Close()

	connPayment, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Payment Service: %v", err)
	}
	defer connPayment.Close()

	connShipping, err := grpc.Dial(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Shipping Service: %v", err)
	}
	defer connShipping.Close()

	// Membuat client untuk masing-masing service
	orderClient := orderPb.NewOrderServiceClient(connOrder)
	paymentClient := paymentPb.NewPaymentServiceClient(connPayment)
	shippingClient := shippingPb.NewShippingServiceClient(connShipping)

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Step 1: Buat Order
	_, err = orderClient.CreateOrder(ctx, &orderPb.CreateOrderRequest{OrderId: orderID})
	if err != nil {
		log.Printf("Order creation failed: %v", err)
		return
	}
	log.Println("Order created successfully")

	// Step 2: Proses Pembayaran
	_, err = paymentClient.ProcessPayment(ctx, &paymentPb.ProcessPaymentRequest{
		OrderId: orderID,
		Amount:  100.0, // Contoh jumlah pembayaran
	})
	if err != nil {
		log.Printf("Payment processing failed, rolling back order: %v", err)
		// Kompensasi: Batalkan order jika pembayaran gagal
		orderClient.CancelOrder(ctx, &orderPb.CancelOrderRequest{OrderId: orderID})
		return
	}
	log.Println("Payment processed successfully")

	// Step 3: Mulai Pengiriman
	_, err = shippingClient.StartShipping(ctx, &shippingPb.StartShippingRequest{
		OrderId:         orderID,
		ShippingAddress: "Contoh Alamat Pengiriman",
	})
	if err != nil {
		log.Printf("Shipping failed, rolling back payment and order: %v", err)
		// Kompensasi: Batalkan pembayaran dan order
		paymentClient.RefundPayment(ctx, &paymentPb.RefundPaymentRequest{OrderId: orderID})
		orderClient.CancelOrder(ctx, &orderPb.CancelOrderRequest{OrderId: orderID})
		return
	}
	log.Println("Shipping started successfully")

	fmt.Println("Transaksi selesai! 🎉")
}

func main() {
	// Contoh penggunaan sagaExecute dengan ID order unik
	sagaExecute("order-123")
}
