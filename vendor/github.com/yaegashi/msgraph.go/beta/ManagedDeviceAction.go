// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedDeviceCollectionExecuteActionRequestParameter undocumented
type ManagedDeviceCollectionExecuteActionRequestParameter struct {
	// ActionName undocumented
	ActionName *ManagedDeviceRemoteAction `json:"actionName,omitempty"`
	// KeepEnrollmentData undocumented
	KeepEnrollmentData *bool `json:"keepEnrollmentData,omitempty"`
	// KeepUserData undocumented
	KeepUserData *bool `json:"keepUserData,omitempty"`
	// DeviceIDs undocumented
	DeviceIDs []string `json:"deviceIds,omitempty"`
	// NotificationTitle undocumented
	NotificationTitle *string `json:"notificationTitle,omitempty"`
	// NotificationBody undocumented
	NotificationBody *string `json:"notificationBody,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
}

// ManagedDeviceOverrideComplianceStateRequestParameter undocumented
type ManagedDeviceOverrideComplianceStateRequestParameter struct {
	// ComplianceState undocumented
	ComplianceState *AdministratorConfiguredDeviceComplianceState `json:"complianceState,omitempty"`
	// RemediationURL undocumented
	RemediationURL *string `json:"remediationUrl,omitempty"`
}

// ManagedDeviceEnableLostModeRequestParameter undocumented
type ManagedDeviceEnableLostModeRequestParameter struct {
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Footer undocumented
	Footer *string `json:"footer,omitempty"`
}

// ManagedDevicePlayLostModeSoundRequestParameter undocumented
type ManagedDevicePlayLostModeSoundRequestParameter struct {
}

// ManagedDeviceSetDeviceNameRequestParameter undocumented
type ManagedDeviceSetDeviceNameRequestParameter struct {
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
}

// ManagedDeviceRotateFileVaultKeyRequestParameter undocumented
type ManagedDeviceRotateFileVaultKeyRequestParameter struct {
}

// ManagedDeviceRetireRequestParameter undocumented
type ManagedDeviceRetireRequestParameter struct {
}

// ManagedDeviceWipeRequestParameter undocumented
type ManagedDeviceWipeRequestParameter struct {
	// KeepEnrollmentData undocumented
	KeepEnrollmentData *bool `json:"keepEnrollmentData,omitempty"`
	// KeepUserData undocumented
	KeepUserData *bool `json:"keepUserData,omitempty"`
	// MacOsUnlockCode undocumented
	MacOsUnlockCode *string `json:"macOsUnlockCode,omitempty"`
}

// ManagedDeviceResetPasscodeRequestParameter undocumented
type ManagedDeviceResetPasscodeRequestParameter struct {
}

// ManagedDeviceRemoteLockRequestParameter undocumented
type ManagedDeviceRemoteLockRequestParameter struct {
}

// ManagedDeviceRequestRemoteAssistanceRequestParameter undocumented
type ManagedDeviceRequestRemoteAssistanceRequestParameter struct {
}

// ManagedDeviceDisableLostModeRequestParameter undocumented
type ManagedDeviceDisableLostModeRequestParameter struct {
}

// ManagedDeviceLocateDeviceRequestParameter undocumented
type ManagedDeviceLocateDeviceRequestParameter struct {
}

// ManagedDeviceBypassActivationLockRequestParameter undocumented
type ManagedDeviceBypassActivationLockRequestParameter struct {
}

// ManagedDeviceRebootNowRequestParameter undocumented
type ManagedDeviceRebootNowRequestParameter struct {
}

// ManagedDeviceShutDownRequestParameter undocumented
type ManagedDeviceShutDownRequestParameter struct {
}

// ManagedDeviceRecoverPasscodeRequestParameter undocumented
type ManagedDeviceRecoverPasscodeRequestParameter struct {
}

// ManagedDeviceCleanWindowsDeviceRequestParameter undocumented
type ManagedDeviceCleanWindowsDeviceRequestParameter struct {
	// KeepUserData undocumented
	KeepUserData *bool `json:"keepUserData,omitempty"`
}

// ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestParameter undocumented
type ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestParameter struct {
}

// ManagedDeviceDeleteUserFromSharedAppleDeviceRequestParameter undocumented
type ManagedDeviceDeleteUserFromSharedAppleDeviceRequestParameter struct {
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// ManagedDeviceSyncDeviceRequestParameter undocumented
type ManagedDeviceSyncDeviceRequestParameter struct {
}

// ManagedDeviceWindowsDefenderScanRequestParameter undocumented
type ManagedDeviceWindowsDefenderScanRequestParameter struct {
	// QuickScan undocumented
	QuickScan *bool `json:"quickScan,omitempty"`
}

// ManagedDeviceWindowsDefenderUpdateSignaturesRequestParameter undocumented
type ManagedDeviceWindowsDefenderUpdateSignaturesRequestParameter struct {
}

// ManagedDeviceUpdateWindowsDeviceAccountRequestParameter undocumented
type ManagedDeviceUpdateWindowsDeviceAccountRequestParameter struct {
	// UpdateWindowsDeviceAccountActionParameter undocumented
	UpdateWindowsDeviceAccountActionParameter *UpdateWindowsDeviceAccountActionParameter `json:"updateWindowsDeviceAccountActionParameter,omitempty"`
}

// ManagedDeviceRevokeAppleVppLicensesRequestParameter undocumented
type ManagedDeviceRevokeAppleVppLicensesRequestParameter struct {
}

// ManagedDeviceRotateBitLockerKeysRequestParameter undocumented
type ManagedDeviceRotateBitLockerKeysRequestParameter struct {
}

// ManagedDeviceSendCustomNotificationToCompanyPortalRequestParameter undocumented
type ManagedDeviceSendCustomNotificationToCompanyPortalRequestParameter struct {
	// NotificationTitle undocumented
	NotificationTitle *string `json:"notificationTitle,omitempty"`
	// NotificationBody undocumented
	NotificationBody *string `json:"notificationBody,omitempty"`
}

// ManagedDeviceTriggerConfigurationManagerActionRequestParameter undocumented
type ManagedDeviceTriggerConfigurationManagerActionRequestParameter struct {
	// ConfigurationManagerAction undocumented
	ConfigurationManagerAction *ConfigurationManagerAction `json:"configurationManagerAction,omitempty"`
}

//
type ManagedDeviceCollectionExecuteActionRequestBuilder struct{ BaseRequestBuilder }

// ExecuteAction action undocumented
func (b *DetectedAppManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ExecuteAction action undocumented
func (b *DeviceManagementManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ExecuteAction action undocumented
func (b *UserManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceCollectionExecuteActionRequest struct{ BaseRequest }

//
func (b *ManagedDeviceCollectionExecuteActionRequestBuilder) Request() *ManagedDeviceCollectionExecuteActionRequest {
	return &ManagedDeviceCollectionExecuteActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceCollectionExecuteActionRequest) Post(ctx context.Context) (resObj *BulkManagedDeviceActionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ManagedDeviceOverrideComplianceStateRequestBuilder struct{ BaseRequestBuilder }

// OverrideComplianceState action undocumented
func (b *ManagedDeviceRequestBuilder) OverrideComplianceState(reqObj *ManagedDeviceOverrideComplianceStateRequestParameter) *ManagedDeviceOverrideComplianceStateRequestBuilder {
	bb := &ManagedDeviceOverrideComplianceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/overrideComplianceState"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceOverrideComplianceStateRequest struct{ BaseRequest }

//
func (b *ManagedDeviceOverrideComplianceStateRequestBuilder) Request() *ManagedDeviceOverrideComplianceStateRequest {
	return &ManagedDeviceOverrideComplianceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceOverrideComplianceStateRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceEnableLostModeRequestBuilder struct{ BaseRequestBuilder }

// EnableLostMode action undocumented
func (b *ManagedDeviceRequestBuilder) EnableLostMode(reqObj *ManagedDeviceEnableLostModeRequestParameter) *ManagedDeviceEnableLostModeRequestBuilder {
	bb := &ManagedDeviceEnableLostModeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enableLostMode"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceEnableLostModeRequest struct{ BaseRequest }

//
func (b *ManagedDeviceEnableLostModeRequestBuilder) Request() *ManagedDeviceEnableLostModeRequest {
	return &ManagedDeviceEnableLostModeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceEnableLostModeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDevicePlayLostModeSoundRequestBuilder struct{ BaseRequestBuilder }

// PlayLostModeSound action undocumented
func (b *ManagedDeviceRequestBuilder) PlayLostModeSound(reqObj *ManagedDevicePlayLostModeSoundRequestParameter) *ManagedDevicePlayLostModeSoundRequestBuilder {
	bb := &ManagedDevicePlayLostModeSoundRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/playLostModeSound"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDevicePlayLostModeSoundRequest struct{ BaseRequest }

//
func (b *ManagedDevicePlayLostModeSoundRequestBuilder) Request() *ManagedDevicePlayLostModeSoundRequest {
	return &ManagedDevicePlayLostModeSoundRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDevicePlayLostModeSoundRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceSetDeviceNameRequestBuilder struct{ BaseRequestBuilder }

// SetDeviceName action undocumented
func (b *ManagedDeviceRequestBuilder) SetDeviceName(reqObj *ManagedDeviceSetDeviceNameRequestParameter) *ManagedDeviceSetDeviceNameRequestBuilder {
	bb := &ManagedDeviceSetDeviceNameRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setDeviceName"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceSetDeviceNameRequest struct{ BaseRequest }

//
func (b *ManagedDeviceSetDeviceNameRequestBuilder) Request() *ManagedDeviceSetDeviceNameRequest {
	return &ManagedDeviceSetDeviceNameRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceSetDeviceNameRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRotateFileVaultKeyRequestBuilder struct{ BaseRequestBuilder }

// RotateFileVaultKey action undocumented
func (b *ManagedDeviceRequestBuilder) RotateFileVaultKey(reqObj *ManagedDeviceRotateFileVaultKeyRequestParameter) *ManagedDeviceRotateFileVaultKeyRequestBuilder {
	bb := &ManagedDeviceRotateFileVaultKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/rotateFileVaultKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRotateFileVaultKeyRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRotateFileVaultKeyRequestBuilder) Request() *ManagedDeviceRotateFileVaultKeyRequest {
	return &ManagedDeviceRotateFileVaultKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRotateFileVaultKeyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRetireRequestBuilder struct{ BaseRequestBuilder }

// Retire action undocumented
func (b *ManagedDeviceRequestBuilder) Retire(reqObj *ManagedDeviceRetireRequestParameter) *ManagedDeviceRetireRequestBuilder {
	bb := &ManagedDeviceRetireRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/retire"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRetireRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRetireRequestBuilder) Request() *ManagedDeviceRetireRequest {
	return &ManagedDeviceRetireRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRetireRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceWipeRequestBuilder struct{ BaseRequestBuilder }

// Wipe action undocumented
func (b *ManagedDeviceRequestBuilder) Wipe(reqObj *ManagedDeviceWipeRequestParameter) *ManagedDeviceWipeRequestBuilder {
	bb := &ManagedDeviceWipeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipe"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceWipeRequest struct{ BaseRequest }

//
func (b *ManagedDeviceWipeRequestBuilder) Request() *ManagedDeviceWipeRequest {
	return &ManagedDeviceWipeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceWipeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceResetPasscodeRequestBuilder struct{ BaseRequestBuilder }

// ResetPasscode action undocumented
func (b *ManagedDeviceRequestBuilder) ResetPasscode(reqObj *ManagedDeviceResetPasscodeRequestParameter) *ManagedDeviceResetPasscodeRequestBuilder {
	bb := &ManagedDeviceResetPasscodeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resetPasscode"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceResetPasscodeRequest struct{ BaseRequest }

//
func (b *ManagedDeviceResetPasscodeRequestBuilder) Request() *ManagedDeviceResetPasscodeRequest {
	return &ManagedDeviceResetPasscodeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceResetPasscodeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRemoteLockRequestBuilder struct{ BaseRequestBuilder }

// RemoteLock action undocumented
func (b *ManagedDeviceRequestBuilder) RemoteLock(reqObj *ManagedDeviceRemoteLockRequestParameter) *ManagedDeviceRemoteLockRequestBuilder {
	bb := &ManagedDeviceRemoteLockRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/remoteLock"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRemoteLockRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRemoteLockRequestBuilder) Request() *ManagedDeviceRemoteLockRequest {
	return &ManagedDeviceRemoteLockRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRemoteLockRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRequestRemoteAssistanceRequestBuilder struct{ BaseRequestBuilder }

// RequestRemoteAssistance action undocumented
func (b *ManagedDeviceRequestBuilder) RequestRemoteAssistance(reqObj *ManagedDeviceRequestRemoteAssistanceRequestParameter) *ManagedDeviceRequestRemoteAssistanceRequestBuilder {
	bb := &ManagedDeviceRequestRemoteAssistanceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/requestRemoteAssistance"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRequestRemoteAssistanceRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRequestRemoteAssistanceRequestBuilder) Request() *ManagedDeviceRequestRemoteAssistanceRequest {
	return &ManagedDeviceRequestRemoteAssistanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRequestRemoteAssistanceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceDisableLostModeRequestBuilder struct{ BaseRequestBuilder }

// DisableLostMode action undocumented
func (b *ManagedDeviceRequestBuilder) DisableLostMode(reqObj *ManagedDeviceDisableLostModeRequestParameter) *ManagedDeviceDisableLostModeRequestBuilder {
	bb := &ManagedDeviceDisableLostModeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/disableLostMode"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceDisableLostModeRequest struct{ BaseRequest }

//
func (b *ManagedDeviceDisableLostModeRequestBuilder) Request() *ManagedDeviceDisableLostModeRequest {
	return &ManagedDeviceDisableLostModeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceDisableLostModeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceLocateDeviceRequestBuilder struct{ BaseRequestBuilder }

// LocateDevice action undocumented
func (b *ManagedDeviceRequestBuilder) LocateDevice(reqObj *ManagedDeviceLocateDeviceRequestParameter) *ManagedDeviceLocateDeviceRequestBuilder {
	bb := &ManagedDeviceLocateDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/locateDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceLocateDeviceRequest struct{ BaseRequest }

//
func (b *ManagedDeviceLocateDeviceRequestBuilder) Request() *ManagedDeviceLocateDeviceRequest {
	return &ManagedDeviceLocateDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceLocateDeviceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceBypassActivationLockRequestBuilder struct{ BaseRequestBuilder }

// BypassActivationLock action undocumented
func (b *ManagedDeviceRequestBuilder) BypassActivationLock(reqObj *ManagedDeviceBypassActivationLockRequestParameter) *ManagedDeviceBypassActivationLockRequestBuilder {
	bb := &ManagedDeviceBypassActivationLockRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bypassActivationLock"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceBypassActivationLockRequest struct{ BaseRequest }

//
func (b *ManagedDeviceBypassActivationLockRequestBuilder) Request() *ManagedDeviceBypassActivationLockRequest {
	return &ManagedDeviceBypassActivationLockRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceBypassActivationLockRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRebootNowRequestBuilder struct{ BaseRequestBuilder }

// RebootNow action undocumented
func (b *ManagedDeviceRequestBuilder) RebootNow(reqObj *ManagedDeviceRebootNowRequestParameter) *ManagedDeviceRebootNowRequestBuilder {
	bb := &ManagedDeviceRebootNowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/rebootNow"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRebootNowRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRebootNowRequestBuilder) Request() *ManagedDeviceRebootNowRequest {
	return &ManagedDeviceRebootNowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRebootNowRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceShutDownRequestBuilder struct{ BaseRequestBuilder }

// ShutDown action undocumented
func (b *ManagedDeviceRequestBuilder) ShutDown(reqObj *ManagedDeviceShutDownRequestParameter) *ManagedDeviceShutDownRequestBuilder {
	bb := &ManagedDeviceShutDownRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/shutDown"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceShutDownRequest struct{ BaseRequest }

//
func (b *ManagedDeviceShutDownRequestBuilder) Request() *ManagedDeviceShutDownRequest {
	return &ManagedDeviceShutDownRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceShutDownRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRecoverPasscodeRequestBuilder struct{ BaseRequestBuilder }

// RecoverPasscode action undocumented
func (b *ManagedDeviceRequestBuilder) RecoverPasscode(reqObj *ManagedDeviceRecoverPasscodeRequestParameter) *ManagedDeviceRecoverPasscodeRequestBuilder {
	bb := &ManagedDeviceRecoverPasscodeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/recoverPasscode"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRecoverPasscodeRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRecoverPasscodeRequestBuilder) Request() *ManagedDeviceRecoverPasscodeRequest {
	return &ManagedDeviceRecoverPasscodeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRecoverPasscodeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceCleanWindowsDeviceRequestBuilder struct{ BaseRequestBuilder }

// CleanWindowsDevice action undocumented
func (b *ManagedDeviceRequestBuilder) CleanWindowsDevice(reqObj *ManagedDeviceCleanWindowsDeviceRequestParameter) *ManagedDeviceCleanWindowsDeviceRequestBuilder {
	bb := &ManagedDeviceCleanWindowsDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cleanWindowsDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceCleanWindowsDeviceRequest struct{ BaseRequest }

//
func (b *ManagedDeviceCleanWindowsDeviceRequestBuilder) Request() *ManagedDeviceCleanWindowsDeviceRequest {
	return &ManagedDeviceCleanWindowsDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceCleanWindowsDeviceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestBuilder struct{ BaseRequestBuilder }

// LogoutSharedAppleDeviceActiveUser action undocumented
func (b *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser(reqObj *ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestParameter) *ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestBuilder {
	bb := &ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/logoutSharedAppleDeviceActiveUser"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceLogoutSharedAppleDeviceActiveUserRequest struct{ BaseRequest }

//
func (b *ManagedDeviceLogoutSharedAppleDeviceActiveUserRequestBuilder) Request() *ManagedDeviceLogoutSharedAppleDeviceActiveUserRequest {
	return &ManagedDeviceLogoutSharedAppleDeviceActiveUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceLogoutSharedAppleDeviceActiveUserRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceDeleteUserFromSharedAppleDeviceRequestBuilder struct{ BaseRequestBuilder }

// DeleteUserFromSharedAppleDevice action undocumented
func (b *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice(reqObj *ManagedDeviceDeleteUserFromSharedAppleDeviceRequestParameter) *ManagedDeviceDeleteUserFromSharedAppleDeviceRequestBuilder {
	bb := &ManagedDeviceDeleteUserFromSharedAppleDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deleteUserFromSharedAppleDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceDeleteUserFromSharedAppleDeviceRequest struct{ BaseRequest }

//
func (b *ManagedDeviceDeleteUserFromSharedAppleDeviceRequestBuilder) Request() *ManagedDeviceDeleteUserFromSharedAppleDeviceRequest {
	return &ManagedDeviceDeleteUserFromSharedAppleDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceDeleteUserFromSharedAppleDeviceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceSyncDeviceRequestBuilder struct{ BaseRequestBuilder }

// SyncDevice action undocumented
func (b *ManagedDeviceRequestBuilder) SyncDevice(reqObj *ManagedDeviceSyncDeviceRequestParameter) *ManagedDeviceSyncDeviceRequestBuilder {
	bb := &ManagedDeviceSyncDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceSyncDeviceRequest struct{ BaseRequest }

//
func (b *ManagedDeviceSyncDeviceRequestBuilder) Request() *ManagedDeviceSyncDeviceRequest {
	return &ManagedDeviceSyncDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceSyncDeviceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceWindowsDefenderScanRequestBuilder struct{ BaseRequestBuilder }

// WindowsDefenderScan action undocumented
func (b *ManagedDeviceRequestBuilder) WindowsDefenderScan(reqObj *ManagedDeviceWindowsDefenderScanRequestParameter) *ManagedDeviceWindowsDefenderScanRequestBuilder {
	bb := &ManagedDeviceWindowsDefenderScanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/windowsDefenderScan"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceWindowsDefenderScanRequest struct{ BaseRequest }

//
func (b *ManagedDeviceWindowsDefenderScanRequestBuilder) Request() *ManagedDeviceWindowsDefenderScanRequest {
	return &ManagedDeviceWindowsDefenderScanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceWindowsDefenderScanRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceWindowsDefenderUpdateSignaturesRequestBuilder struct{ BaseRequestBuilder }

// WindowsDefenderUpdateSignatures action undocumented
func (b *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures(reqObj *ManagedDeviceWindowsDefenderUpdateSignaturesRequestParameter) *ManagedDeviceWindowsDefenderUpdateSignaturesRequestBuilder {
	bb := &ManagedDeviceWindowsDefenderUpdateSignaturesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/windowsDefenderUpdateSignatures"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceWindowsDefenderUpdateSignaturesRequest struct{ BaseRequest }

//
func (b *ManagedDeviceWindowsDefenderUpdateSignaturesRequestBuilder) Request() *ManagedDeviceWindowsDefenderUpdateSignaturesRequest {
	return &ManagedDeviceWindowsDefenderUpdateSignaturesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceWindowsDefenderUpdateSignaturesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceUpdateWindowsDeviceAccountRequestBuilder struct{ BaseRequestBuilder }

// UpdateWindowsDeviceAccount action undocumented
func (b *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount(reqObj *ManagedDeviceUpdateWindowsDeviceAccountRequestParameter) *ManagedDeviceUpdateWindowsDeviceAccountRequestBuilder {
	bb := &ManagedDeviceUpdateWindowsDeviceAccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateWindowsDeviceAccount"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceUpdateWindowsDeviceAccountRequest struct{ BaseRequest }

//
func (b *ManagedDeviceUpdateWindowsDeviceAccountRequestBuilder) Request() *ManagedDeviceUpdateWindowsDeviceAccountRequest {
	return &ManagedDeviceUpdateWindowsDeviceAccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceUpdateWindowsDeviceAccountRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRevokeAppleVppLicensesRequestBuilder struct{ BaseRequestBuilder }

// RevokeAppleVppLicenses action undocumented
func (b *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses(reqObj *ManagedDeviceRevokeAppleVppLicensesRequestParameter) *ManagedDeviceRevokeAppleVppLicensesRequestBuilder {
	bb := &ManagedDeviceRevokeAppleVppLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeAppleVppLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRevokeAppleVppLicensesRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRevokeAppleVppLicensesRequestBuilder) Request() *ManagedDeviceRevokeAppleVppLicensesRequest {
	return &ManagedDeviceRevokeAppleVppLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRevokeAppleVppLicensesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceRotateBitLockerKeysRequestBuilder struct{ BaseRequestBuilder }

// RotateBitLockerKeys action undocumented
func (b *ManagedDeviceRequestBuilder) RotateBitLockerKeys(reqObj *ManagedDeviceRotateBitLockerKeysRequestParameter) *ManagedDeviceRotateBitLockerKeysRequestBuilder {
	bb := &ManagedDeviceRotateBitLockerKeysRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/rotateBitLockerKeys"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceRotateBitLockerKeysRequest struct{ BaseRequest }

//
func (b *ManagedDeviceRotateBitLockerKeysRequestBuilder) Request() *ManagedDeviceRotateBitLockerKeysRequest {
	return &ManagedDeviceRotateBitLockerKeysRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceRotateBitLockerKeysRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceSendCustomNotificationToCompanyPortalRequestBuilder struct{ BaseRequestBuilder }

// SendCustomNotificationToCompanyPortal action undocumented
func (b *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal(reqObj *ManagedDeviceSendCustomNotificationToCompanyPortalRequestParameter) *ManagedDeviceSendCustomNotificationToCompanyPortalRequestBuilder {
	bb := &ManagedDeviceSendCustomNotificationToCompanyPortalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendCustomNotificationToCompanyPortal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceSendCustomNotificationToCompanyPortalRequest struct{ BaseRequest }

//
func (b *ManagedDeviceSendCustomNotificationToCompanyPortalRequestBuilder) Request() *ManagedDeviceSendCustomNotificationToCompanyPortalRequest {
	return &ManagedDeviceSendCustomNotificationToCompanyPortalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceSendCustomNotificationToCompanyPortalRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceTriggerConfigurationManagerActionRequestBuilder struct{ BaseRequestBuilder }

// TriggerConfigurationManagerAction action undocumented
func (b *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction(reqObj *ManagedDeviceTriggerConfigurationManagerActionRequestParameter) *ManagedDeviceTriggerConfigurationManagerActionRequestBuilder {
	bb := &ManagedDeviceTriggerConfigurationManagerActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/triggerConfigurationManagerAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceTriggerConfigurationManagerActionRequest struct{ BaseRequest }

//
func (b *ManagedDeviceTriggerConfigurationManagerActionRequestBuilder) Request() *ManagedDeviceTriggerConfigurationManagerActionRequest {
	return &ManagedDeviceTriggerConfigurationManagerActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceTriggerConfigurationManagerActionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}