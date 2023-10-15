package calculations

import "testing"

func TestAlphanumericValues(t *testing.T) {
	v := AlphanumericCharacters("Target", 1)
	if v != 6 {
		t.Fatalf("alphanumeric values dont match. expected=%d, got=%d", 6, v)
	}
}

type testPrice struct {
	Price         string
	ExpectedValue int
	ExpectedError bool
}

func TestIsZeroCents(t *testing.T) {
	testCases := []testPrice{
		{Price: "1.00", ExpectedValue: 1, ExpectedError: false},
		{Price: "123.20", ExpectedValue: 0, ExpectedError: false},
		{Price: "asdf.20", ExpectedValue: 0, ExpectedError: true},
	}

	for _, tt := range testCases {
		p, err := IsZeroCents(tt.Price, 1)
		if p != tt.ExpectedValue {
			t.Fatalf("expected different point value, expected=%d, got=%d", tt.ExpectedValue, p)
		}

		if err != nil && !tt.ExpectedError {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}
}

type testItemsLength struct {
	Lenght        int
	ExpectedValue int
}

func TestTwoItemsPoints(t *testing.T) {
	testCases := []testItemsLength{
		{Lenght: 1, ExpectedValue: 0},
		{Lenght: 4, ExpectedValue: 2},
		{Lenght: 5, ExpectedValue: 2},
	}

	for _, tt := range testCases {
		p := TwoItemsPoints(tt.Lenght, 1)
		if p != tt.ExpectedValue {
			t.Fatalf("problem with items lenght of %d. expected=%d, got=%d", tt.Lenght, tt.ExpectedValue, p)
		}
	}
}

func TestIsMultipleOf25(t *testing.T) {
	const POINTS = 1
	testCases := []testPrice{
		{Price: "100.00", ExpectedValue: POINTS, ExpectedError: false},
		{Price: "100.10", ExpectedValue: 0, ExpectedError: false},
		{Price: "asdf9.00", ExpectedValue: 0, ExpectedError: true},
	}

	for _, tt := range testCases {
		p, err := IsMultipleOf25(tt.Price, POINTS)
		if p != tt.ExpectedValue {
			t.Fatalf("expected different point value, expected=%d, got=%d", tt.ExpectedValue, p)
		}

		if err != nil && !tt.ExpectedError {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}
}

type descriptionTest struct {
	Description   string
	Price         string
	ExpectedValue int
	ExpectedError bool
}

func TestTrimmedDescription(t *testing.T) {
	testCases := []descriptionTest{
		{Description: "Emils Cheese Pizza", Price: "12.25", ExpectedValue: 3, ExpectedError: false},
		{Description: "  Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00", ExpectedValue: 3, ExpectedError: false},
		{Description: "asds", Price: "10.00", ExpectedValue: 0, ExpectedError: false},
		{Description: "asd", Price: "10s.00", ExpectedValue: 0, ExpectedError: true},
	}

	for _, tt := range testCases {
		p, err := TrimmedDescriptionPoints(tt.Description, tt.Price)
		if p != tt.ExpectedValue {
			t.Fatalf("expected different point value, expected=%d, got=%d", tt.ExpectedValue, p)
		}

		if err != nil && !tt.ExpectedError {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}
}

type dateOddsTest struct {
	Date          string
	ExpectedValue int
	ExpectedError bool
}

func TestDateIsOddPOint(t *testing.T) {
	const POINTS = 1
	testCases := []dateOddsTest{
		{Date: "2023-01-03", ExpectedValue: POINTS, ExpectedError: false},
		{Date: "2023-01-02", ExpectedValue: 0, ExpectedError: false},
		{Date: "20230103", ExpectedValue: 0, ExpectedError: true},
	}

	for _, tt := range testCases {
		p, err := DateIsOddPoints(tt.Date, POINTS)
		if p != tt.ExpectedValue {
			t.Fatalf("expected different point value, expected=%d, got=%d", tt.ExpectedValue, p)
		}

		if err != nil && !tt.ExpectedError {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}
}

type betweenTimesTest struct {
	Time          string
	InitialTime   int
	FinalTime     int
	ExpectedValue int
	ExpectedError bool
}

func TestBetweenTimePoints(t *testing.T) {
	const POINTS = 1
	testCases := []betweenTimesTest{
		{Time: "14:01", InitialTime: 14, FinalTime: 16, ExpectedValue: POINTS, ExpectedError: false},
		{Time: "13:01", InitialTime: 14, FinalTime: 16, ExpectedValue: 0, ExpectedError: false},
		{Time: "14:01:0", InitialTime: 14, FinalTime: 16, ExpectedValue: 0, ExpectedError: true},
	}

	for _, tt := range testCases {
		p, err := BetweenTimePoints(tt.Time, tt.InitialTime, tt.FinalTime, POINTS)
		if p != tt.ExpectedValue {
			t.Fatalf("expected different point value for time %s, expected=%d, got=%d", tt.Time, tt.ExpectedValue, p)
		}

		if err != nil && !tt.ExpectedError {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}

}
