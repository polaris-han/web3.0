package sql_practice

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance float64
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func (a *Account) InitTable(db *gorm.DB) {
	db.AutoMigrate(a)
}

func (t *Transaction) InitTable(db *gorm.DB) {
	db.AutoMigrate(t)
}

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，
// 如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务
func Transfer(db *gorm.DB, fromAccountID uint, toAccountID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 查询转出账户
		var fromAccount Account
		tx.Debug().First(&fromAccount, fromAccountID)
		fmt.Println("转出账户信息:", fromAccount)

		// 检查账户余额
		if fromAccount.Balance < amount {
			tx.Rollback()
			return fmt.Errorf("转出账户余额不足")
		}
		// 从转出账户扣除金额
		tx.Debug().Model(&fromAccount).Update("balance", fromAccount.Balance-amount)

		var toAccount Account
		// 查询转入账户
		tx.Debug().First(&toAccount, toAccountID)
		fmt.Println("转入账户信息:", toAccount)
		// 转入账户增加金额
		tx.Debug().Model(&toAccount).Update("balance", toAccount.Balance+amount)

		transaction := Transaction{
			FromAccountId: fromAccountID,
			ToAccountId:   toAccountID,
			Amount:        amount,
		}

		tx.Debug().Save(&transaction)

		return nil
	})
}

func InsertRandomAccounts(db *gorm.DB, numAccounts int) error {
	a := Account{}
	a.InitTable(db)

	t := Transaction{}
	t.InitTable(db)

	// 设置随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成并插入随机账户数据
	for range numAccounts {
		// 生成随机余额，范围为 0 到 10000
		balance := r.Float64() * 10000

		// 创建 Account 实例
		account := Account{
			Balance: balance,
		}

		// 插入到数据库
		if err := db.Create(&account).Error; err != nil {
			return fmt.Errorf("failed to insert account: %w", err)
		}
	}

	return nil
}
