// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetCwfNetworkIDLiUesReader is a Reader for the GetCwfNetworkIDLiUes structure.
type GetCwfNetworkIDLiUesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDLiUesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDLiUesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDLiUesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDLiUesOK creates a GetCwfNetworkIDLiUesOK with default headers values
func NewGetCwfNetworkIDLiUesOK() *GetCwfNetworkIDLiUesOK {
	return &GetCwfNetworkIDLiUesOK{}
}

/*GetCwfNetworkIDLiUesOK handles this case with default header values.

Monitored LI UEs
*/
type GetCwfNetworkIDLiUesOK struct {
	Payload *models.LiUes
}

func (o *GetCwfNetworkIDLiUesOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/li_ues][%d] getCwfNetworkIdLiUesOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDLiUesOK) GetPayload() *models.LiUes {
	return o.Payload
}

func (o *GetCwfNetworkIDLiUesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LiUes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDLiUesDefault creates a GetCwfNetworkIDLiUesDefault with default headers values
func NewGetCwfNetworkIDLiUesDefault(code int) *GetCwfNetworkIDLiUesDefault {
	return &GetCwfNetworkIDLiUesDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDLiUesDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDLiUesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID li ues default response
func (o *GetCwfNetworkIDLiUesDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDLiUesDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/li_ues][%d] GetCwfNetworkIDLiUes default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDLiUesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDLiUesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
