package blockchain

import "net/http"

// Error data structure describing the error.
type Error struct {
	// Address wrong address.
	Address *string
	// ErrMain error information from the standard package error set.
	ErrMain error
	// ErrExec information about the error that occurred during
	// the operation of the standard library or external packages.
	ErrExec error
	// Response http response.
	Response *http.Response
}

// Error compatibility with error interface.
func (e Error) Error() string {
	return e.ErrMain.Error()
}

// NewError creates a new Error instance.
func NewError(errorMain error, errorExec error, response *http.Response, address *string) *Error {
	if errorMain == nil {
		return nil
	}

	return &Error{
		ErrMain:  errorMain,
		ErrExec:  errorExec,
		Response: response,
		Address:  address,
	}
}

// err build error helper.
func (c *Client) err(errorMain error) error {
	return NewError(errorMain, nil, nil, nil)
}

// err2 build error helper.
func (c *Client) err2(errorMain error, errorExec error) error {
	return NewError(errorMain, errorExec, nil, nil)
}

// err3 build error helper.
func (c *Client) err3(errorMain error, errorExec error, response *http.Response) error {
	return NewError(errorMain, errorExec, response, nil)
}

// err4 build error helper.
func (c *Client) err4(errorMain error, response string) error {
	return NewError(errorMain, nil, nil, &response)
}
