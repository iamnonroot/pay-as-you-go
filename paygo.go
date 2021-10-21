/*
	Pay as you Go library to scheduler deduct user's wallet balance every X seconds
*/
package paygo

import (
	"time"

	"github.com/go-co-op/gocron"
)

/*
	Create new scheduler job by
	setting a unique name/id for job and
	user wallet id and price of product
*/
type ItemOption struct {
	UUID   string
	Wallet string
	Price  int
}

/*
	Create a new scheduler by
	setting how many every seconds jobs run and
	function to get user wallet balance and
	function to set user wallet balance (if return be false, job will removed)
*/
type Item struct {
	scheduler gocron.Scheduler
	Every     int
	DoFunc    func(option ItemOption) bool
}

/*
	Start scheduler
*/
func (item *Item) Start() {
	item.scheduler.StartAsync()
}

/*
	The function that run in scheduler job
*/
func (item *Item) doFunc(option ItemOption) {
	if !item.DoFunc(option) {
		item.Remove(option.UUID)
	}
}

/*
	Add a new scheduler job with a UUID and do it for every X seconds and a function to do
*/
func (item *Item) Add(option ItemOption) {
	item.scheduler.Every(item.Every).Seconds().Tag(option.UUID).Do(item.doFunc, option)
}

/*
	Remove a scheduler job with UUID
*/
func (item *Item) Remove(uuid string) {
	item.scheduler.RemoveByTag(uuid)
}

/*
	Create a new pay as you go with creating a new scheduler
*/
func New(item *Item) *Item {
	item.scheduler = *gocron.NewScheduler(time.UTC)
	item.scheduler.TagsUnique()
	item.scheduler.SingletonMode()

	return item
}
