// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceapiDeviceSwagger deviceapi device swagger
//
// swagger:model deviceapi.DeviceSwagger
type DeviceapiDeviceSwagger struct {

	// agent load flags
	AgentLoadFlags string `json:"agent_load_flags,omitempty"`

	// agent local time
	AgentLocalTime string `json:"agent_local_time,omitempty"`

	// agent version
	AgentVersion string `json:"agent_version,omitempty"`

	// base image version
	BaseImageVersion string `json:"base_image_version,omitempty"`

	// bios manufacturer
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`

	// bios version
	BiosVersion string `json:"bios_version,omitempty"`

	// build number
	BuildNumber string `json:"build_number,omitempty"`

	// chassis type
	ChassisType string `json:"chassis_type,omitempty"`

	// chassis type desc
	ChassisTypeDesc string `json:"chassis_type_desc,omitempty"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// config id base
	ConfigIDBase string `json:"config_id_base,omitempty"`

	// config id build
	ConfigIDBuild string `json:"config_id_build,omitempty"`

	// config id platform
	ConfigIDPlatform string `json:"config_id_platform,omitempty"`

	// connection ip
	ConnectionIP string `json:"connection_ip,omitempty"`

	// connection mac address
	ConnectionMacAddress string `json:"connection_mac_address,omitempty"`

	// cpu signature
	CPUSignature string `json:"cpu_signature,omitempty"`

	// cpu vendor
	CPUVendor string `json:"cpu_vendor,omitempty"`

	// default gateway ip
	DefaultGatewayIP string `json:"default_gateway_ip,omitempty"`

	// deployment type
	DeploymentType string `json:"deployment_type,omitempty"`

	// detection suppression status
	DetectionSuppressionStatus string `json:"detection_suppression_status,omitempty"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// device policies
	DevicePolicies *DeviceapiMappedDevicePoliciesSwagger `json:"device_policies,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// external ip
	ExternalIP string `json:"external_ip,omitempty"`

	// filesystem containment status
	FilesystemContainmentStatus string `json:"filesystem_containment_status,omitempty"`

	// first login timestamp
	FirstLoginTimestamp string `json:"first_login_timestamp,omitempty"`

	// first seen
	FirstSeen string `json:"first_seen,omitempty"`

	// group hash
	GroupHash string `json:"group_hash,omitempty"`

	// groups
	Groups []string `json:"groups"`

	// host hidden status
	HostHiddenStatus string `json:"host_hidden_status,omitempty"`

	// host utc offset
	HostUtcOffset string `json:"host_utc_offset,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// instance id
	InstanceID string `json:"instance_id,omitempty"`

	// internet exposure
	InternetExposure string `json:"internet_exposure,omitempty"`

	// k8s cluster git version
	K8sClusterGitVersion string `json:"k8s_cluster_git_version,omitempty"`

	// k8s cluster id
	K8sClusterID string `json:"k8s_cluster_id,omitempty"`

	// k8s cluster version
	K8sClusterVersion string `json:"k8s_cluster_version,omitempty"`

	// kernel version
	KernelVersion string `json:"kernel_version,omitempty"`

	// last login timestamp
	LastLoginTimestamp string `json:"last_login_timestamp,omitempty"`

	// last login uid
	LastLoginUID string `json:"last_login_uid,omitempty"`

	// last login user
	LastLoginUser string `json:"last_login_user,omitempty"`

	// last login user sid
	LastLoginUserSid string `json:"last_login_user_sid,omitempty"`

	// last reboot
	LastReboot string `json:"last_reboot,omitempty"`

	// last seen
	LastSeen string `json:"last_seen,omitempty"`

	// linux sensor mode
	LinuxSensorMode string `json:"linux_sensor_mode,omitempty"`

	// local ip
	LocalIP string `json:"local_ip,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// machine domain
	MachineDomain string `json:"machine_domain,omitempty"`

	// major version
	MajorVersion string `json:"major_version,omitempty"`

	// managed apps
	ManagedApps *DeviceManagedApps `json:"managed_apps,omitempty"`

	// meta
	Meta *DeviceDeviceMeta `json:"meta,omitempty"`

	// migration completed time
	MigrationCompletedTime string `json:"migration_completed_time,omitempty"`

	// minor version
	MinorVersion string `json:"minor_version,omitempty"`

	// modified timestamp
	ModifiedTimestamp string `json:"modified_timestamp,omitempty"`

	// notes
	Notes []string `json:"notes"`

	// os build
	OsBuild string `json:"os_build,omitempty"`

	// os product name
	OsProductName string `json:"os_product_name,omitempty"`

	// os version
	OsVersion string `json:"os_version,omitempty"`

	// ou
	Ou []string `json:"ou"`

	// platform id
	PlatformID string `json:"platform_id,omitempty"`

	// platform name
	PlatformName string `json:"platform_name,omitempty"`

	// pod annotations
	PodAnnotations []string `json:"pod_annotations"`

	// pod host ip4
	PodHostIp4 string `json:"pod_host_ip4,omitempty"`

	// pod host ip6
	PodHostIp6 string `json:"pod_host_ip6,omitempty"`

	// pod hostname
	PodHostname string `json:"pod_hostname,omitempty"`

	// pod id
	PodID string `json:"pod_id,omitempty"`

	// pod ip4
	PodIp4 string `json:"pod_ip4,omitempty"`

	// pod ip6
	PodIp6 string `json:"pod_ip6,omitempty"`

	// pod labels
	PodLabels []string `json:"pod_labels"`

	// pod name
	PodName string `json:"pod_name,omitempty"`

	// pod namespace
	PodNamespace string `json:"pod_namespace,omitempty"`

	// pod service account name
	PodServiceAccountName string `json:"pod_service_account_name,omitempty"`

	// pointer size
	PointerSize string `json:"pointer_size,omitempty"`

	// policies
	Policies []*DeviceDevicePolicy `json:"policies"`

	// product type
	ProductType string `json:"product_type,omitempty"`

	// product type desc
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// provision status
	ProvisionStatus string `json:"provision_status,omitempty"`

	// reduced functionality mode
	ReducedFunctionalityMode string `json:"reduced_functionality_mode,omitempty"`

	// release group
	ReleaseGroup string `json:"release_group,omitempty"`

	// rtr state
	RtrState string `json:"rtr_state,omitempty"`

	// serial number
	SerialNumber string `json:"serial_number,omitempty"`

	// service pack major
	ServicePackMajor string `json:"service_pack_major,omitempty"`

	// service pack minor
	ServicePackMinor string `json:"service_pack_minor,omitempty"`

	// service provider
	ServiceProvider string `json:"service_provider,omitempty"`

	// service provider account id
	ServiceProviderAccountID string `json:"service_provider_account_id,omitempty"`

	// site name
	SiteName string `json:"site_name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// system manufacturer
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// system product name
	SystemProductName string `json:"system_product_name,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// zone group
	ZoneGroup string `json:"zone_group,omitempty"`
}

// Validate validates this deviceapi device swagger
func (m *DeviceapiDeviceSwagger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevicePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceapiDeviceSwagger) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) validateDevicePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.DevicePolicies) { // not required
		return nil
	}

	if m.DevicePolicies != nil {
		if err := m.DevicePolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_policies")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) validateManagedApps(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagedApps) { // not required
		return nil
	}

	if m.ManagedApps != nil {
		if err := m.ManagedApps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_apps")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this deviceapi device swagger based on the context it is used
func (m *DeviceapiDeviceSwagger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevicePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagedApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceapiDeviceSwagger) contextValidateDevicePolicies(ctx context.Context, formats strfmt.Registry) error {

	if m.DevicePolicies != nil {

		if swag.IsZero(m.DevicePolicies) { // not required
			return nil
		}

		if err := m.DevicePolicies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_policies")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) contextValidateManagedApps(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagedApps != nil {

		if swag.IsZero(m.ManagedApps) { // not required
			return nil
		}

		if err := m.ManagedApps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_apps")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if swag.IsZero(m.Meta) { // not required
			return nil
		}

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceapiDeviceSwagger) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {

			if swag.IsZero(m.Policies[i]) { // not required
				return nil
			}

			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceapiDeviceSwagger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceapiDeviceSwagger) UnmarshalBinary(b []byte) error {
	var res DeviceapiDeviceSwagger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
