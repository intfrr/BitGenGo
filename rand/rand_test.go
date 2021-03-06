package rand

import (
	"testing"
)

func TestBasicBit(t *testing.T) {
	rnd, err := NewRandomizer(50)
	if err != nil {
		t.Errorf("Failed to create a new Randomizer")
	}
	err = rnd.Powerup()
	if err != nil {
		t.Errorf("Failed to powerup the Randomizer")
	}
	_, err = rnd.GetBit()
	if err != nil {
		t.Errorf("Failed to get a bit from the Randomizer")
	}
	err = rnd.Shutdown()
	if err != nil {
		t.Errorf("Failed to shutdown the Randomizer")
	}
}

func TestBasicInt(t *testing.T) {
	rnd, err := NewRandomizer(44)
	if err != nil {
		t.Errorf("Failed to create a new Randomizer")
	}
	err = rnd.Powerup()
	if err != nil {
		t.Errorf("Failed to powerup the Randomizer")
	}
	_, err = rnd.GetInt(20, 100)
	if err != nil {
		t.Errorf("Failed to get an int from the Randomizer")
	}
	err = rnd.Shutdown()
	if err != nil {
		t.Errorf("Failed to shutdown the Randomizer")
	}
}

func Test1000MeanVariance(t *testing.T) {
	rnd, _ := NewRandomizer(20)
	rnd.Powerup()
	sum := 0
	for i := 1; i < 1000; i++ {
		b, _ := rnd.GetBit()
		sum = sum + b
	}
	rnd.Shutdown()
	var mean float64 = float64(sum) / 1000
	if mean < 0.45 || mean > 0.55 {
		t.Errorf("Error in variance of the bits")
	}
}

func TestDoublePowerup(t *testing.T) {
	rnd, _ := NewRandomizer(30)
	rnd.Powerup()
	err := rnd.Powerup()
	if err == nil {
		t.Fatalf("Double Powerup should crash!")
	}
}

func TestDoubleShutdown(t *testing.T) {
	rnd, _ := NewRandomizer(34)
	rnd.Powerup()
	rnd.Shutdown()
	err := rnd.Shutdown()
	if err == nil {
		t.Fatalf("Double Shutdown should crash!")
	}
}

func TestGetBitOfShutdownRandomizer(t *testing.T) {
	rnd, _ := NewRandomizer(40)
	_, err := rnd.GetBit()
	if err == nil {
		t.Fatalf("Getting bit from shutdown randomizer should crash!")
	}
}

func TestMinimumInterval(t *testing.T) {
	_, err := NewRandomizer(19)
	if err == nil {
		t.Fatalf("Minimum interval is 20 milliseconds")
	}
}
