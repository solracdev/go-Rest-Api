package channels

func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}

	return false
}

func ERROR(err <-chan error) bool {
	select {
	case e := <-err:
		if e != nil {
			return true
		}
	}
	return false
}
