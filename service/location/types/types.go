// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains error details for each geofence that failed to delete from the geofence
// collection.
type BatchDeleteGeofenceError struct {

	// Contains details associated to the batch error.
	//
	// This member is required.
	Error *BatchItemError

	// The geofence associated with the error message.
	//
	// This member is required.
	GeofenceId *string
}

// Contains error details for each device that failed to evaluate its position
// against the geofences in a given geofence collection.
type BatchEvaluateGeofencesError struct {

	// The device associated with the position evaluation error.
	//
	// This member is required.
	DeviceId *string

	// Contains details associated to the batch error.
	//
	// This member is required.
	Error *BatchItemError

	// Specifies a timestamp for when the error occurred in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	SampleTime *time.Time
}

// Contains error details for each device that didn't return a position.
type BatchGetDevicePositionError struct {

	// The ID of the device that didn't return a position.
	//
	// This member is required.
	DeviceId *string

	// Contains details related to the error code.
	//
	// This member is required.
	Error *BatchItemError
}

// Contains the batch request error details associated with the request.
type BatchItemError struct {

	// The error code associated with the batch request error.
	Code BatchItemErrorCode

	// A message with the reason for the batch request error.
	Message *string
}

// Contains error details for each geofence that failed to be stored in a given
// geofence collection.
type BatchPutGeofenceError struct {

	// Contains details associated to the batch error.
	//
	// This member is required.
	Error *BatchItemError

	// The geofence associated with the error message.
	//
	// This member is required.
	GeofenceId *string
}

// Contains geofence geometry details.
type BatchPutGeofenceRequestEntry struct {

	// The identifier for the geofence to be stored in a given geofence collection.
	//
	// This member is required.
	GeofenceId *string

	// Contains the polygon details to specify the position of the geofence. Each
	// geofence polygon
	// (https://docs.aws.amazon.com/location-geofences/latest/APIReference/API_GeofenceGeometry.html)
	// can have a maximum of 1,000 vertices.
	//
	// This member is required.
	Geometry *GeofenceGeometry
}

// Contains a summary of each geofence that was successfully stored in a given
// geofence collection.
type BatchPutGeofenceSuccess struct {

	// The timestamp for when the geofence was stored in a geofence collection in ISO
	// 8601 (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	CreateTime *time.Time

	// The geofence successfully stored in a geofence collection.
	//
	// This member is required.
	GeofenceId *string

	// The timestamp for when the geofence was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	UpdateTime *time.Time
}

// Contains error details for each device that failed to update its position.
type BatchUpdateDevicePositionError struct {

	// The device associated with the failed location update.
	//
	// This member is required.
	DeviceId *string

	// Contains details related to the error code such as the error code and error
	// message.
	//
	// This member is required.
	Error *BatchItemError

	// The timestamp at which the device position was determined. Uses  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	SampleTime *time.Time
}

// Specifies the data storage option chosen for requesting Places. By using Places,
// you agree that AWS may transmit your API queries to your selected third party
// provider for processing, which may be outside the AWS region you are currently
// using. Also, when using HERE as your data provider, you may not (a) use HERE
// Places for Asset Management, or (b) select the Storage option for the
// IntendedUse parameter when requesting Places in Japan. For more information, see
// the AWS Service Terms (https://aws.amazon.com/service-terms/) for Amazon
// Location Service.
type DataSourceConfiguration struct {

	// Specifies how the results of an operation will be stored by the caller. Valid
	// values include:
	//
	// * SingleUse specifies that the results won't be stored.
	//
	// *
	// Storage specifies that the result can be cached or stored in a
	// database.
	//
	// Default value: SingleUse
	IntendedUse IntendedUse
}

// Contains the device position details.
type DevicePosition struct {

	// The last known device position.
	//
	// This member is required.
	Position []float64

	// The timestamp for when the tracker resource received the device position in  ISO
	// 8601 (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	ReceivedTime *time.Time

	// The timestamp at which the device's position was determined. Uses  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	SampleTime *time.Time

	// The device whose position you retrieved.
	DeviceId *string
}

// Contains the position update details for a device.
type DevicePositionUpdate struct {

	// The device associated to the position update.
	//
	// This member is required.
	DeviceId *string

	// The latest device position defined in WGS 84
	// (https://earth-info.nga.mil/GandG/wgs84/index.html) format: [X or longitude, Y
	// or latitude].
	//
	// This member is required.
	Position []float64

	// The timestamp at which the device's position was determined. Uses ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	SampleTime *time.Time
}

// Contains the geofence geometry details. Amazon Location does not currently
// support polygons with holes, multipolygons, polygons that are wound clockwise,
// or that cross the antimeridian.
type GeofenceGeometry struct {

	// An array of 1 or more linear rings. A linear ring is an array of 4 or more
	// vertices, where the first and last vertex are the same to form a closed
	// boundary. Each vertex is a 2-dimensional point of the form: [longitude,
	// latitude]. The first linear ring is an outer ring, describing the polygon's
	// boundary. Subsequent linear rings may be inner or outer rings to describe holes
	// and islands. Outer rings must list their vertices in counter-clockwise order
	// around the ring's center, where the left side is the polygon's exterior. Inner
	// rings must list their vertices in clockwise order, where the left side is the
	// polygon's interior.
	Polygon [][][]float64
}

// Contains the geofence collection details.
type ListGeofenceCollectionsResponseEntry struct {

	// The name of the geofence collection.
	//
	// This member is required.
	CollectionName *string

	// The timestamp for when the geofence collection was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	CreateTime *time.Time

	// The description for the geofence collection
	//
	// This member is required.
	Description *string

	// The pricing plan for the specified geofence collection. For additional details
	// and restrictions on each pricing plan option, see the Amazon Location Service
	// pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan PricingPlan

	// Specifies a timestamp for when the resource was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	UpdateTime *time.Time

	// The data source selected for the geofence collection and associated pricing
	// plan.
	PricingPlanDataSource *string
}

// Contains a list of geofences stored in a given geofence collection.
type ListGeofenceResponseEntry struct {

	// The timestamp for when the geofence was stored in a geofence collection in ISO
	// 8601 (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	CreateTime *time.Time

	// The geofence identifier.
	//
	// This member is required.
	GeofenceId *string

	// Contains the geofence geometry details describing a polygon.
	//
	// This member is required.
	Geometry *GeofenceGeometry

	// Identifies the state of the geofence. A geofence will hold one of the following
	// states:
	//
	// * ACTIVE — The geofence has been indexed by the system.
	//
	// * PENDING —
	// The geofence is being processed by the system.
	//
	// * FAILED — The geofence failed
	// to be indexed by the system.
	//
	// * DELETED — The geofence has been deleted from the
	// system index.
	//
	// * DELETING — The geofence is being deleted from the system index.
	//
	// This member is required.
	Status *string

	// The timestamp for when the geofence was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	UpdateTime *time.Time
}

// Contains details of an existing map resource in your AWS account.
type ListMapsResponseEntry struct {

	// The timestamp for when the map resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// Specifies the data provider for the associated map tiles.
	//
	// This member is required.
	DataSource *string

	// The description for the map resource.
	//
	// This member is required.
	Description *string

	// The name of the associated map resource.
	//
	// This member is required.
	MapName *string

	// The pricing plan for the specified map resource. For additional details and
	// restrictions on each pricing plan option, see the Amazon Location Service
	// pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan PricingPlan

	// The timestamp for when the map resource was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time
}

// A Place index resource listed in your AWS account.
type ListPlaceIndexesResponseEntry struct {

	// The timestamp for when the Place index resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The data provider of geospatial data. Indicates one of the available
	// providers:
	//
	// * Esri
	//
	// * HERE
	//
	// For additional details on data providers, see the
	// Amazon Location Service data providers page
	// (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html).
	//
	// This member is required.
	DataSource *string

	// The optional description for the Place index resource.
	//
	// This member is required.
	Description *string

	// The name of the Place index resource.
	//
	// This member is required.
	IndexName *string

	// The pricing plan for the specified Place index resource. For additional details
	// and restrictions on each pricing plan option, see the Amazon Location Service
	// pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan PricingPlan

	// The timestamp for when the Place index resource was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time
}

// Contains the tracker resource details.
type ListTrackersResponseEntry struct {

	// The timestamp for when the tracker resource was created in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The description for the tracker resource.
	//
	// This member is required.
	Description *string

	// The pricing plan for the specified tracker resource. For additional details and
	// restrictions on each pricing plan option, see the Amazon Location Service
	// pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan PricingPlan

	// The name of the tracker resource.
	//
	// This member is required.
	TrackerName *string

	// The timestamp at which the device's position was determined. Uses  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// The data source selected for the tracker resource and associated pricing plan.
	PricingPlanDataSource *string
}

// Specifies the map tile style selected from an available provider.
type MapConfiguration struct {

	// Specifies the map style selected from an available data provider. Valid styles:
	// VectorEsriStreets, VectorEsriTopographic, VectorEsriNavigation,
	// VectorEsriDarkGrayCanvas, VectorEsriLightGrayCanvas, VectorHereBerlin. When
	// using HERE as your data provider, and selecting the Style VectorHereBerlin, you
	// may not use HERE Maps for Asset Management. See the AWS Service Terms
	// (https://aws.amazon.com/service-terms/) for Amazon Location Service.
	//
	// This member is required.
	Style *string
}

// Contains details about addresses or points of interest that match the search
// criteria.
type Place struct {

	// Places uses a point geometry to specify a location or a Place.
	//
	// This member is required.
	Geometry *PlaceGeometry

	// The numerical portion of an address, such as a building number.
	AddressNumber *string

	// A country/region specified using ISO 3166
	// (https://www.iso.org/iso-3166-country-codes.html) 3-digit country/region code.
	// For example, CAN.
	Country *string

	// The full name and address of the point of interest such as a city, region, or
	// country. For example, 123 Any Street, Any Town, USA.
	Label *string

	// A name for a local area, such as a city or town name. For example, Toronto.
	Municipality *string

	// The name of a community district. For example, Downtown.
	Neighborhood *string

	// A group of numbers and letters in a country-specific format, which accompanies
	// the address for the purpose of identifying a location.
	PostalCode *string

	// A name for an area or geographical division, such as a province or state name.
	// For example, British Columbia.
	Region *string

	// The name for a street or a road to identify a location. For example, Main
	// Street.
	Street *string

	// A country, or an area that's part of a larger region . For example, Metro
	// Vancouver.
	SubRegion *string
}

// Places uses a point geometry to specify a location or a Place.
type PlaceGeometry struct {

	// A single point geometry specifies a location for a Place using WGS 84
	// (https://gisgeography.com/wgs84-world-geodetic-system/) coordinates:
	//
	// * x —
	// Specifies the x coordinate or longitude.
	//
	// * y — Specifies the y coordinate or
	// latitude.
	Point []float64
}

// Specifies a single point of interest, or Place as a result of a search query
// obtained from a dataset configured in the Place index Resource.
type SearchForPositionResult struct {

	// Contains details about the relevant point of interest.
	//
	// This member is required.
	Place *Place
}

// Contains relevant Places returned by calling SearchPlaceIndexForText.
type SearchForTextResult struct {

	// Contains details about the relevant point of interest.
	//
	// This member is required.
	Place *Place
}

// A summary of the reverse geocoding request sent using
// SearchPlaceIndexForPosition.
type SearchPlaceIndexForPositionSummary struct {

	// The data provider of geospatial data. Indicates one of the available
	// providers:
	//
	// * Esri
	//
	// * HERE
	//
	// For additional details on data providers, see the
	// Amazon Location Service data providers page
	// (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html).
	//
	// This member is required.
	DataSource *string

	// The position given in the reverse geocoding request.
	//
	// This member is required.
	Position []float64

	// An optional parameter. The maximum number of results returned per request.
	// Default value: 50
	MaxResults int32
}

// A summary of the geocoding request sent using SearchPlaceIndexForText.
type SearchPlaceIndexForTextSummary struct {

	// The data provider of geospatial data. Indicates one of the available
	// providers:
	//
	// * Esri
	//
	// * HERE
	//
	// For additional details on data providers, see the
	// Amazon Location Service data providers page
	// (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html).
	//
	// This member is required.
	DataSource *string

	// The address, name, city or region to be used in the geocoding request. In
	// free-form text format. For example, Vancouver.
	//
	// This member is required.
	Text *string

	// Contains the coordinates for the bias position entered in the geocoding request.
	BiasPosition []float64

	// Contains the coordinates for the optional bounding box coordinated entered in
	// the geocoding request.
	FilterBBox []float64

	// Contains the country filter entered in the geocoding request.
	FilterCountries []string

	// Contains the maximum number of results indicated for the request.
	MaxResults int32

	// A bounding box that contains the search results within the specified area
	// indicated by FilterBBox. A subset of bounding box specified using FilterBBox.
	ResultBBox []float64
}

// The input failed to meet the constraints specified by the AWS service in a
// specified field.
type ValidationExceptionField struct {

	// A message with the reason for the validation exception error.
	//
	// This member is required.
	Message *string

	// The field name where the invalid entry was detected.
	//
	// This member is required.
	Name *string
}