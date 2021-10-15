// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetLTENetworkIDGatewaysGatewayIDCellularNonEpsReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDCellularNonEps structure.
type GetLTENetworkIDGatewaysGatewayIDCellularNonEpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent creates a GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent() *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent {
	return &GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent{}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent handles this case with default header values.

Non-EPS configuration
*/
type GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent struct {
	Payload *models.GatewayNonEpsConfigs
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/non_eps][%d] getLteNetworkIdGatewaysGatewayIdCellularNonEpsNoContent  %+v", 204, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent) GetPayload() *models.GatewayNonEpsConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayNonEpsConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault creates a GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault(code int) *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault {
	return &GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID cellular non eps default response
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/non_eps][%d] GetLTENetworkIDGatewaysGatewayIDCellularNonEps default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
