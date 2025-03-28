// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gaf.proto

package firebase_api

import (
	go_android_utils "github.com/BRUHItsABunny/go-android-utils"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FirebaseAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Expires      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires,proto3" json:"expires,omitempty"`
	RefreshToken string                 `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	IdToken      string                 `protobuf:"bytes,4,opt,name=idToken,proto3" json:"idToken,omitempty"`
}

func (x *FirebaseAuthentication) Reset() {
	*x = FirebaseAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirebaseAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirebaseAuthentication) ProtoMessage() {}

func (x *FirebaseAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_gaf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirebaseAuthentication.ProtoReflect.Descriptor instead.
func (*FirebaseAuthentication) Descriptor() ([]byte, []int) {
	return file_gaf_proto_rawDescGZIP(), []int{0}
}

func (x *FirebaseAuthentication) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *FirebaseAuthentication) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *FirebaseAuthentication) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *FirebaseAuthentication) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

type FirebaseNotificationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationToken string `protobuf:"bytes,1,opt,name=notificationToken,proto3" json:"notificationToken,omitempty"`
	// The fields below are needed for Chromium push notification, not for native Android apps (typically)
	PrivateKey []byte `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	PublicKey  []byte `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Secret     []byte `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *FirebaseNotificationData) Reset() {
	*x = FirebaseNotificationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirebaseNotificationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirebaseNotificationData) ProtoMessage() {}

func (x *FirebaseNotificationData) ProtoReflect() protoreflect.Message {
	mi := &file_gaf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirebaseNotificationData.ProtoReflect.Descriptor instead.
func (*FirebaseNotificationData) Descriptor() ([]byte, []int) {
	return file_gaf_proto_rawDescGZIP(), []int{1}
}

func (x *FirebaseNotificationData) GetNotificationToken() string {
	if x != nil {
		return x.NotificationToken
	}
	return ""
}

func (x *FirebaseNotificationData) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *FirebaseNotificationData) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *FirebaseNotificationData) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

type FirebaseInstallationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirebaseInstallationID     string                    `protobuf:"bytes,1,opt,name=FirebaseInstallationID,proto3" json:"FirebaseInstallationID,omitempty"`
	NotificationData           *FirebaseNotificationData `protobuf:"bytes,2,opt,name=notificationData,proto3" json:"notificationData,omitempty"`
	InstallationAuthentication *FirebaseAuthentication   `protobuf:"bytes,3,opt,name=installationAuthentication,proto3" json:"installationAuthentication,omitempty"` // Upon NotifyInstall
}

func (x *FirebaseInstallationData) Reset() {
	*x = FirebaseInstallationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirebaseInstallationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirebaseInstallationData) ProtoMessage() {}

func (x *FirebaseInstallationData) ProtoReflect() protoreflect.Message {
	mi := &file_gaf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirebaseInstallationData.ProtoReflect.Descriptor instead.
func (*FirebaseInstallationData) Descriptor() ([]byte, []int) {
	return file_gaf_proto_rawDescGZIP(), []int{2}
}

func (x *FirebaseInstallationData) GetFirebaseInstallationID() string {
	if x != nil {
		return x.FirebaseInstallationID
	}
	return ""
}

func (x *FirebaseInstallationData) GetNotificationData() *FirebaseNotificationData {
	if x != nil {
		return x.NotificationData
	}
	return nil
}

func (x *FirebaseInstallationData) GetInstallationAuthentication() *FirebaseAuthentication {
	if x != nil {
		return x.InstallationAuthentication
	}
	return nil
}

type FirebaseAppData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageID            string `protobuf:"bytes,1,opt,name=packageID,proto3" json:"packageID,omitempty"`                       // com.app.my
	PackageCertificate   string `protobuf:"bytes,2,opt,name=packageCertificate,proto3" json:"packageCertificate,omitempty"`     // HEX string in UPPER case
	GoogleAPIKey         string `protobuf:"bytes,3,opt,name=googleAPIKey,proto3" json:"googleAPIKey,omitempty"`                 // Some string, typically starts with "AIza"
	FirebaseProjectID    string `protobuf:"bytes,4,opt,name=firebaseProjectID,proto3" json:"firebaseProjectID,omitempty"`       // some string identifier equal to what the project ID is in Firebase Console
	NotificationSenderID string `protobuf:"bytes,5,opt,name=notificationSenderID,proto3" json:"notificationSenderID,omitempty"` // int string
	GMPAppID             string `protobuf:"bytes,6,opt,name=GMPAppID,proto3" json:"GMPAppID,omitempty"`                         // 1:NotificationSenderID:android:SOME_ID
	AppVersion           string `protobuf:"bytes,7,opt,name=AppVersion,proto3" json:"AppVersion,omitempty"`                     // Major.Minor.Micro
	AppVersionWithBuild  string `protobuf:"bytes,8,opt,name=AppVersionWithBuild,proto3" json:"AppVersionWithBuild,omitempty"`   // Similar to AppVersion but without periods and also build number appended
	AuthVersion          string `protobuf:"bytes,9,opt,name=authVersion,proto3" json:"authVersion,omitempty"`                   // Always FIS_v2
	SdkVersion           string `protobuf:"bytes,10,opt,name=sdkVersion,proto3" json:"sdkVersion,omitempty"`                    // This can vary between apps
	AppNameHash          string `protobuf:"bytes,11,opt,name=appNameHash,proto3" json:"appNameHash,omitempty"`
}

func (x *FirebaseAppData) Reset() {
	*x = FirebaseAppData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirebaseAppData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirebaseAppData) ProtoMessage() {}

func (x *FirebaseAppData) ProtoReflect() protoreflect.Message {
	mi := &file_gaf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirebaseAppData.ProtoReflect.Descriptor instead.
func (*FirebaseAppData) Descriptor() ([]byte, []int) {
	return file_gaf_proto_rawDescGZIP(), []int{3}
}

func (x *FirebaseAppData) GetPackageID() string {
	if x != nil {
		return x.PackageID
	}
	return ""
}

func (x *FirebaseAppData) GetPackageCertificate() string {
	if x != nil {
		return x.PackageCertificate
	}
	return ""
}

func (x *FirebaseAppData) GetGoogleAPIKey() string {
	if x != nil {
		return x.GoogleAPIKey
	}
	return ""
}

func (x *FirebaseAppData) GetFirebaseProjectID() string {
	if x != nil {
		return x.FirebaseProjectID
	}
	return ""
}

func (x *FirebaseAppData) GetNotificationSenderID() string {
	if x != nil {
		return x.NotificationSenderID
	}
	return ""
}

func (x *FirebaseAppData) GetGMPAppID() string {
	if x != nil {
		return x.GMPAppID
	}
	return ""
}

func (x *FirebaseAppData) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *FirebaseAppData) GetAppVersionWithBuild() string {
	if x != nil {
		return x.AppVersionWithBuild
	}
	return ""
}

func (x *FirebaseAppData) GetAuthVersion() string {
	if x != nil {
		return x.AuthVersion
	}
	return ""
}

func (x *FirebaseAppData) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *FirebaseAppData) GetAppNameHash() string {
	if x != nil {
		return x.AppNameHash
	}
	return ""
}

type FirebaseDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device               *go_android_utils.Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	CheckinAndroidID     int64                    `protobuf:"varint,2,opt,name=checkinAndroidID,proto3" json:"checkinAndroidID,omitempty"`
	CheckinSecurityToken uint64                   `protobuf:"varint,3,opt,name=checkinSecurityToken,proto3" json:"checkinSecurityToken,omitempty"`
	// app Package ID : auth Obj
	FirebaseInstallations map[string]*FirebaseInstallationData `protobuf:"bytes,4,rep,name=firebaseInstallations,proto3" json:"firebaseInstallations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MTalkLastPersistentId string                               `protobuf:"bytes,5,opt,name=mTalkLastPersistentId,proto3" json:"mTalkLastPersistentId,omitempty"`
	FirebaseClientVersion string                               `protobuf:"bytes,6,opt,name=firebaseClientVersion,proto3" json:"firebaseClientVersion,omitempty"`
	GmsVersion            string                               `protobuf:"bytes,7,opt,name=gmsVersion,proto3" json:"gmsVersion,omitempty"`
	MTalkPrivateKey       string                               `protobuf:"bytes,8,opt,name=mTalkPrivateKey,proto3" json:"mTalkPrivateKey,omitempty"`
	MTalkPublicKey        string                               `protobuf:"bytes,9,opt,name=mTalkPublicKey,proto3" json:"mTalkPublicKey,omitempty"`
	MTalkAuthSecret       string                               `protobuf:"bytes,10,opt,name=mTalkAuthSecret,proto3" json:"mTalkAuthSecret,omitempty"`
}

func (x *FirebaseDevice) Reset() {
	*x = FirebaseDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirebaseDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirebaseDevice) ProtoMessage() {}

func (x *FirebaseDevice) ProtoReflect() protoreflect.Message {
	mi := &file_gaf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirebaseDevice.ProtoReflect.Descriptor instead.
func (*FirebaseDevice) Descriptor() ([]byte, []int) {
	return file_gaf_proto_rawDescGZIP(), []int{4}
}

func (x *FirebaseDevice) GetDevice() *go_android_utils.Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *FirebaseDevice) GetCheckinAndroidID() int64 {
	if x != nil {
		return x.CheckinAndroidID
	}
	return 0
}

func (x *FirebaseDevice) GetCheckinSecurityToken() uint64 {
	if x != nil {
		return x.CheckinSecurityToken
	}
	return 0
}

func (x *FirebaseDevice) GetFirebaseInstallations() map[string]*FirebaseInstallationData {
	if x != nil {
		return x.FirebaseInstallations
	}
	return nil
}

func (x *FirebaseDevice) GetMTalkLastPersistentId() string {
	if x != nil {
		return x.MTalkLastPersistentId
	}
	return ""
}

func (x *FirebaseDevice) GetFirebaseClientVersion() string {
	if x != nil {
		return x.FirebaseClientVersion
	}
	return ""
}

func (x *FirebaseDevice) GetGmsVersion() string {
	if x != nil {
		return x.GmsVersion
	}
	return ""
}

func (x *FirebaseDevice) GetMTalkPrivateKey() string {
	if x != nil {
		return x.MTalkPrivateKey
	}
	return ""
}

func (x *FirebaseDevice) GetMTalkPublicKey() string {
	if x != nil {
		return x.MTalkPublicKey
	}
	return ""
}

func (x *FirebaseDevice) GetMTalkAuthSecret() string {
	if x != nil {
		return x.MTalkAuthSecret
	}
	return ""
}

var File_gaf_proto protoreflect.FileDescriptor

var file_gaf_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x16, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x94, 0x02, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x36, 0x0a, 0x16, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x56, 0x0a, 0x10, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x66,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x68, 0x0a, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x5f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb7, 0x03, 0x0a,
	0x0f, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x2e,
	0x0a, 0x12, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x32, 0x0a, 0x14, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x4d, 0x50, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x4d, 0x50, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41,
	0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0x90, 0x05, 0x0a, 0x0e, 0x46, 0x69, 0x72, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x71, 0x0a, 0x15, 0x66, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x6d,
	0x54, 0x61, 0x6c, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x54, 0x61, 0x6c,
	0x6b, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x15, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6d, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6d, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x54, 0x61, 0x6c, 0x6b,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6d, 0x54, 0x61, 0x6c, 0x6b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x54, 0x61, 0x6c, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x54, 0x61, 0x6c, 0x6b,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x54, 0x61,
	0x6c, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x54, 0x61, 0x6c, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0x74, 0x0a, 0x1a, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x52, 0x55, 0x48, 0x49, 0x74, 0x73, 0x41,
	0x42, 0x75, 0x6e, 0x6e, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x2d, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gaf_proto_rawDescOnce sync.Once
	file_gaf_proto_rawDescData = file_gaf_proto_rawDesc
)

func file_gaf_proto_rawDescGZIP() []byte {
	file_gaf_proto_rawDescOnce.Do(func() {
		file_gaf_proto_rawDescData = protoimpl.X.CompressGZIP(file_gaf_proto_rawDescData)
	})
	return file_gaf_proto_rawDescData
}

var file_gaf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gaf_proto_goTypes = []interface{}{
	(*FirebaseAuthentication)(nil),   // 0: android_firebase.FirebaseAuthentication
	(*FirebaseNotificationData)(nil), // 1: android_firebase.FirebaseNotificationData
	(*FirebaseInstallationData)(nil), // 2: android_firebase.FirebaseInstallationData
	(*FirebaseAppData)(nil),          // 3: android_firebase.FirebaseAppData
	(*FirebaseDevice)(nil),           // 4: android_firebase.FirebaseDevice
	nil,                              // 5: android_firebase.FirebaseDevice.FirebaseInstallationsEntry
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*go_android_utils.Device)(nil),  // 7: android_utils.Device
}
var file_gaf_proto_depIdxs = []int32{
	6, // 0: android_firebase.FirebaseAuthentication.expires:type_name -> google.protobuf.Timestamp
	1, // 1: android_firebase.FirebaseInstallationData.notificationData:type_name -> android_firebase.FirebaseNotificationData
	0, // 2: android_firebase.FirebaseInstallationData.installationAuthentication:type_name -> android_firebase.FirebaseAuthentication
	7, // 3: android_firebase.FirebaseDevice.device:type_name -> android_utils.Device
	5, // 4: android_firebase.FirebaseDevice.firebaseInstallations:type_name -> android_firebase.FirebaseDevice.FirebaseInstallationsEntry
	2, // 5: android_firebase.FirebaseDevice.FirebaseInstallationsEntry.value:type_name -> android_firebase.FirebaseInstallationData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gaf_proto_init() }
func file_gaf_proto_init() {
	if File_gaf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gaf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirebaseAuthentication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gaf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirebaseNotificationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gaf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirebaseInstallationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gaf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirebaseAppData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gaf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirebaseDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gaf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gaf_proto_goTypes,
		DependencyIndexes: file_gaf_proto_depIdxs,
		MessageInfos:      file_gaf_proto_msgTypes,
	}.Build()
	File_gaf_proto = out.File
	file_gaf_proto_rawDesc = nil
	file_gaf_proto_goTypes = nil
	file_gaf_proto_depIdxs = nil
}
