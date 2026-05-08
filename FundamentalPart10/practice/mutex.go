package practice

import "sync"

type Account struct {
	mu      sync.Mutex
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

type DonationBox struct {
	Balance int
	mu      sync.Mutex
}

func (d *DonationBox) Donate(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Balance += amount
}
