// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package genericattributeprofile

type GattProtectionLevel int32

const SignatureGattProtectionLevel string = "enum(Windows.Devices.Bluetooth.GenericAttributeProfile.GattProtectionLevel;i4)"

const (
	GattProtectionLevelPlain                               GattProtectionLevel = 0
	GattProtectionLevelAuthenticationRequired              GattProtectionLevel = 1
	GattProtectionLevelEncryptionRequired                  GattProtectionLevel = 2
	GattProtectionLevelEncryptionAndAuthenticationRequired GattProtectionLevel = 3
)