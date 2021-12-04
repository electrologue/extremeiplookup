package extremeiplookup

type IPInfo struct {
	// Client IP address or IP address specified.
	Query string `json:"query,omitempty"`

	// Business, Education or Residential.
	// (Residential is an IP address from an Internet, Hosting or Cloud provider)
	IPType string `json:"ipType,omitempty"`

	// Name of Business / Education Organization.
	BusinessName string `json:"businessName,omitempty"`

	// Website domain of Business / Education Organization.
	BusinessWebsite string `json:"businessWebsite,omitempty"`

	// Name of the continent.
	Continent string `json:"continent,omitempty"`

	// Two-letter ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"countryCode,omitempty"`

	// Name of the country.
	Country string `json:"country,omitempty"`

	// Name of the region.
	Region string `json:"region,omitempty"`

	// Name of the city.
	City string `json:"city,omitempty"`

	// Latitude.
	Latitude string `json:"lat,omitempty"`

	// Longitude.
	Longitude string `json:"lon,omitempty"`

	// Resolved IP Name.
	IPName string `json:"ipName,omitempty"`

	// Organization Name
	Organization string `json:"org,omitempty"`

	// ISP Name
	ISP string `json:"isp,omitempty"`

	// Pro Feature! Samples: America/Chicago, Europe/London, see all on Wikipedia
	Timezone string `json:"timezone,omitempty"`

	// Pro Feature! Samples: -10:00, +02:00
	UTCOffset string `json:"utcOffset,omitempty"`

	// Success or fail
	Status string `json:"status,omitempty"`

	// Extra status message if fail
	Message string `json:"message,omitempty"`
}
