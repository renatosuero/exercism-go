package clock

import "fmt"

//Clock keeps time, limited to a day
type Clock int

//New creates a Clock set to a given time.
func New(hour, minute int) Clock {
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock(time)
}

//String returns a clock in digital form hh:mm.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

//Add adds minutes in the time
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

//Subtract removes minutes in the time
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
