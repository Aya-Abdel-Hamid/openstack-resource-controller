package v1alpha1

// SubnetPoolFilter specifies a filter to select a subnet. At least one parameter must be specified.
// +kubebuilder:validation:MinProperties:=1
type SubnetPoolFilter struct {
    Name        *OpenStackName      `json:"name,omitempty"`
    Description *NeutronDescription `json:"description,omitempty"`
    IPVersion   *IPVersion          `json:"ipVersion,omitempty"`
    GatewayIP   *IPvAny             `json:"gatewayIP,omitempty"`
    CIDR        *CIDR               `json:"cidr,omitempty"`

    // Consistent with Spec: use IPv6Options as the field name and type
    Ipv6 *IPv6Options `json:"ipv6,omitempty"`

    NetworkRef KubernetesNameRef  `json:"networkRef"`
    ProjectRef *KubernetesNameRef `json:"projectRef,omitempty"`

    FilterByNeutronTags `json:",inline"`
}

type SubnetPoolResourceSpec struct {
    Name        *OpenStackName      `json:"name,omitempty"`
    Description *NeutronDescription `json:"description,omitempty"`
    NetworkRef  KubernetesNameRef   `json:"networkRef"`

    // Add listType annotations here
    // +kubebuilder:validation:MaxItems:=32
    // +listType=set
    Tags []NeutronTag `json:"tags,omitempty"`

    IPVersion IPVersion `json:"ipVersion"`
    CIDR      CIDR      `json:"cidr"`

    // +kubebuilder:validation:MaxItems:=32
    // +listType=atomic
    AllocationPools []AllocationPool `json:"allocationPools,omitempty"`

    Gateway          *SubnetPoolGateway `json:"gateway,omitempty"`
    EnableDHCP       *bool              `json:"enableDHCP,omitempty"`

    // +kubebuilder:validation:MaxItems:=16
    // +listType=set
    DNSNameservers []IPvAny `json:"dnsNameservers,omitempty"`

    DNSPublishFixedIP *bool `json:"dnsPublishFixedIP,omitempty"`

    // +kubebuilder:validation:MaxItems:=256
    // +listType=atomic
    HostRoutes []HostRoute `json:"hostRoutes,omitempty"`

    Ipv6 *IPv6Options `json:"ipv6,omitempty"`

    RouterRef  *KubernetesNameRef `json:"routerRef,omitempty"`
    ProjectRef *KubernetesNameRef `json:"projectRef,omitempty"`
}

type SubnetPoolResourceStatus struct {
    Name        string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`

    IPVersion *int32 `json:"ipVersion,omitempty"`
    CIDR      string `json:"cidr,omitempty"`
    GatewayIP string `json:"gatewayIP,omitempty"`

    // Add listType annotations here for status fields too
    // +kubebuilder:validation:MaxItems:=16
    // +listType=set
    DNSNameservers []string `json:"dnsNameservers,omitempty"`

    DNSPublishFixedIP *bool `json:"dnsPublishFixedIP,omitempty"`

    // +kubebuilder:validation:MaxItems:=32
    // +listType=atomic
    AllocationPools []AllocationPoolStatus `json:"allocationPools,omitempty"`

    // +kubebuilder:validation:MaxItems:=256
    // +listType=atomic
    HostRoutes []HostRouteStatus `json:"hostRoutes,omitempty"`

    EnableDHCP *bool `json:"enableDHCP,omitempty"`

    NetworkID string `json:"networkID,omitempty"`
    ProjectID string `json:"projectID,omitempty"`

    // For status, use the same names as in IPv6Options struct (lowercase json names),
    // but in status it's common to flatten these into strings.
    // To satisfy the API rule, rename these fields to exactly "addressMode" and "raMode"
    // OR embed IPv6Options struct here (if allowed).
    // Safer: embed IPv6Options

    Ipv6 *IPv6Options `json:"ipv6,omitempty"`

    SubnetPoolID string `json:"subnetPoolID,omitempty"`

    // +kubebuilder:validation:MaxItems:=32
    // +listType=set
    Tags []string `json:"tags,omitempty"`

    NeutronStatusMetadata `json:",inline"`
}

type SubnetPoolGatewayType string

const (
    SubnetPoolGatewayTypeNone      = "None"
    SubnetPoolGatewayTypeAutomatic = "Automatic"
    SubnetPoolGatewayTypeIP        = "IP"
)

type SubnetPoolGateway struct {
    Type SubnetPoolGatewayType `json:"type"`
    IP   *IPvAny               `json:"ip,omitempty"`
}

