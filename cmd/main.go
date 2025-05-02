package main

import (
	"context"
	"fmt"
	"time"

	"github.com/williamkoller/bank-transfer/internal/account/application"
	"github.com/williamkoller/bank-transfer/internal/account/domain"
	"github.com/williamkoller/bank-transfer/internal/account/infra"
	"github.com/williamkoller/bank-transfer/pkg/uuid"
)

func main() {
	now := time.Now()
	ctx := context.Background()
	repo := infra.NewInMemoryAccountRepo()
	svc := application.NewTransferService(repo)

	acc1, _ := domain.NewAccount(uuid.New(), "Alice", 1000)
	acc2, _ := domain.NewAccount(uuid.New(), "Bob", 500)

	repo.Save(ctx, acc1)
	repo.Save(ctx, acc2)

	err := svc.Transfer(ctx, acc1.ID, acc2.ID, 300)
	if err != nil {
		panic(err)
	}

	a1, _ := repo.FindByID(ctx, acc1.ID)
	a2, _ := repo.FindByID(ctx, acc2.ID)

	fmt.Printf("Saldo Alice: %.2f, ID: %s, balance: %.2f\n", a1.Balance, a1.ID, a1.Balance)
	fmt.Printf("Saldo Bob:   %.2f ID: %s, balance: %.2f\n", a2.Balance, a2.ID, a2.Balance)
	fmt.Printf("Took: %dms\n", time.Since(now).Milliseconds())
}
