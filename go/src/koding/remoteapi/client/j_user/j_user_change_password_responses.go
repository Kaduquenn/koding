package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JUserChangePasswordReader is a Reader for the JUserChangePassword structure.
type JUserChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserChangePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserChangePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserChangePasswordOK creates a JUserChangePasswordOK with default headers values
func NewJUserChangePasswordOK() *JUserChangePasswordOK {
	return &JUserChangePasswordOK{}
}

/*JUserChangePasswordOK handles this case with default header values.

Request processed successfully
*/
type JUserChangePasswordOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserChangePasswordOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.changePassword][%d] jUserChangePasswordOK  %+v", 200, o.Payload)
}

func (o *JUserChangePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserChangePasswordUnauthorized creates a JUserChangePasswordUnauthorized with default headers values
func NewJUserChangePasswordUnauthorized() *JUserChangePasswordUnauthorized {
	return &JUserChangePasswordUnauthorized{}
}

/*JUserChangePasswordUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserChangePasswordUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserChangePasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.changePassword][%d] jUserChangePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserChangePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
