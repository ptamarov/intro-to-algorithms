package main

func OneDimPeakSlow(xs []int) int {
	// This is θ(N) where N = len(xs).
	n := len(xs)

	if n == 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	// Check boundaries (includes n == 2 case which will always return something)
	if xs[0] >= xs[1] {
		return 0
	} else if xs[n-2] <= xs[n-1] {
		return n - 1
	}

	// We now know that the peak must happen inside the array,
	// and that n >= 3 so it makes sense to look at these.
	for i := 1; i < n-1; i++ {
		if xs[i-1] <= xs[i] && xs[i] >= xs[i+1] {
			return i
		}
	}

	return -1
}

func OneDimPeakFast(xs []int) int {

	n := len(xs)

	if n <= 1 {
		return n - 1
	}

	if n/2-1 >= 0 && xs[n/2] <= xs[n/2-1] {
		return OneDimPeakFast(xs[:n/2]) // Note that if there is a boundary peak here, the if condition creates an inner peak, so things work out.

	}

	if n/2+1 < n && xs[n/2] <= xs[n/2+1] {
		// Same here!
		return n/2 + 1 + OneDimPeakFast(xs[n/2+1:]) // Be sure to shift index to account for truncating the array from the left.
	} else {
		return n / 2
	}

}
