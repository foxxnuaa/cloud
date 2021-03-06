package uri

// Resource Service URIs.
const (
	API     string = "/api"
	Version string = API + "/v1"
	Devices string = Version + "/devices"
	Device  string = Devices + "/{{ .DeviceId }}"

	ResourceValues string = Devices + "/{{ .DeviceId }}/{{ .ResourceLinkHref }}"

	DevicesSubscriptions string = Devices + "/subscriptions"
	DevicesSubscription  string = Devices + "/subscriptions/{{ .SubscriptionID }}"

	DeviceSubscriptions string = Devices + "​/{{ .DeviceId }}/subscriptions"
	DeviceSubscription  string = Devices + "/{{ .DeviceId }}/subscriptions/{{ .SubscriptionID }}"

	ResourceSubscriptions string = Devices + "/{{ .DeviceId }}/{{ .ResourceLinkHref }}/subscriptions"
	ResourceSubscription  string = Devices + "/{{ .DeviceId }}/{{ .ResourceLinkHref }}/subscriptions/{{ .SubscriptionID }}"
)
