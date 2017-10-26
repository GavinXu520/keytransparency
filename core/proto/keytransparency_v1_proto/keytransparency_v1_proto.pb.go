// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/keytransparency_v1_proto/keytransparency_v1_proto.proto

/*
Package keytransparency_v1_proto is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	proto/keytransparency_v1_proto/keytransparency_v1_proto.proto
	proto/keytransparency_v1_proto/keytransparency_v1_admin_proto.proto

It has these top-level messages:
	Committed
	EntryUpdate
	Entry
	MutationProof
	MapperMetadata
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	GetMutationsRequest
	GetMutationsResponse
	GetDomainInfoRequest
	GetDomainInfoResponse
	UserProfile
	GetEpochsRequest
	GetEpochsResponse
	BatchUpdateEntriesRequest
	BatchUpdateEntriesResponse
	Domain
	ListDomainsRequest
	ListDomainsResponse
	GetDomainRequest
	GetDomainResponse
	CreateDomainRequest
	CreateDomainResponse
	DeleteDomainRequest
	UndeleteDomainRequest
*/
package keytransparency_v1_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import trillian "github.com/google/trillian"
import trillian1 "github.com/google/trillian"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Committed) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Committed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// mutation authorizes the change to entry.
	Mutation *Entry `protobuf:"bytes,2,opt,name=mutation" json:"mutation,omitempty"`
	// commitment contains the data committed to in update.commitment.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EntryUpdate) GetMutation() *Entry {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Entry is a signed change to a map entry.
// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// index is the location of this leaf in the sparse merkle tree.
	Index []byte `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,6,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*keyspb.PublicKey `protobuf:"bytes,7,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,8,opt,name=previous,proto3" json:"previous,omitempty"`
	// signatures on key_value. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[string]*sigpb.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Entry) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Entry) GetAuthorizedKeys() []*keyspb.PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

func (m *Entry) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *Entry) GetSignatures() map[string]*sigpb.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// MutationProof contains the information necessary to compute the new leaf value.
// It contains a) the old leaf value with it's inclusion proof and b) the mutation.
// The new leaf value is computed via:
//       Mutate(leaf_value, mutation)
type MutationProof struct {
	// mutation contains the information needed to modify the old leaf.
	// The format of a mutation is specific to the particular Mutate function being used.
	Mutation *Entry `protobuf:"bytes,1,opt,name=mutation" json:"mutation,omitempty"`
	// leaf_proof contains the leaf and its inclusion proof for a particular map revision.
	LeafProof *trillian1.MapLeafInclusion `protobuf:"bytes,2,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
}

func (m *MutationProof) Reset()                    { *m = MutationProof{} }
func (m *MutationProof) String() string            { return proto.CompactTextString(m) }
func (*MutationProof) ProtoMessage()               {}
func (*MutationProof) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MutationProof) GetMutation() *Entry {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (m *MutationProof) GetLeafProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

// MapperMetadata tracks the mutations that have been mapped so far. It is
// embedded in the Trillian SignedMapHead.
type MapperMetadata struct {
	HighestFullyCompletedSeq int64 `protobuf:"varint,1,opt,name=highest_fully_completed_seq,json=highestFullyCompletedSeq" json:"highest_fully_completed_seq,omitempty"`
}

func (m *MapperMetadata) Reset()                    { *m = MapperMetadata{} }
func (m *MapperMetadata) String() string            { return proto.CompactTextString(m) }
func (*MapperMetadata) ProtoMessage()               {}
func (*MapperMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MapperMetadata) GetHighestFullyCompletedSeq() int64 {
	if m != nil {
		return m.HighestFullyCompletedSeq
	}
	return 0
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// domain_id identifies the domain in which the user and application live.
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// user_id is the user identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetEntryRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *GetEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,1,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,2,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse Merkle
	// Tree.
	LeafProof *trillian1.MapLeafInclusion `protobuf:"bytes,3,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smr contains the signed map head for the sparse Merkle Tree.
	// smr is also stored in the append only log.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,4,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	// TODO: gossip the log root to verify global consistency.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,5,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,6,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,7,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetEntryResponse) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetEntryResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetEntryResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetEntryResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

// ListEntryHistoryRequest gets a list of historical keys for a user.
type ListEntryHistoryRequest struct {
	// domain_id identifies the domain in which the user and application live.
	DomainId string `protobuf:"bytes,6,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epoch.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,5,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListEntryHistoryRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// ListEntryHistoryResponse requests a paginated history of keys for a user.
type ListEntryHistoryResponse struct {
	// values represents the list of keys this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// next_start is the next page token to query for pagination.
	// next_start is 0 when there are no more results to fetch.
	NextStart int64 `protobuf:"varint,2,opt,name=next_start,json=nextStart" json:"next_start,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ListEntryHistoryResponse) GetNextStart() int64 {
	if m != nil {
		return m.NextStart
	}
	return 0
}

// UpdateEntryRequest updates a user's profile.
type UpdateEntryRequest struct {
	// domain_id identifies the domain in which the user and application live.
	DomainId string `protobuf:"bytes,5,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// user_id specifies the id for the user who's profile is being updated.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// entry_update contains the user submitted update.
	EntryUpdate *EntryUpdate `protobuf:"bytes,4,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateEntryRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *UpdateEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UpdateEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkle Tree.
type UpdateEntryResponse struct {
	// proof contains a proof that the update has been included in the tree.
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// GetMutationsRequest contains the input parameters of the GetMutation APIs.
type GetMutationsRequest struct {
	// domain_id is the domain for which epochs are being requested.
	DomainId string `protobuf:"bytes,5,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// epoch specifies the epoch number in which mutations will be returned.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,2,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// page_token defines the starting point for pagination. An empty
	// value means start from the beginning. A non-empty value requests the next
	// page of values.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// page_size is the maximum number of epochs to return.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *GetMutationsRequest) Reset()                    { *m = GetMutationsRequest{} }
func (m *GetMutationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsRequest) ProtoMessage()               {}
func (*GetMutationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetMutationsRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetMutationsRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *GetMutationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *GetMutationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// GetMutationsResponse contains the results of GetMutation APIs.
type GetMutationsResponse struct {
	// epoch specifies the epoch number of the returned mutations.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// smr contains the signed map root for the sparse Merkle Tree.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,2,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,3,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,4,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,5,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
	// mutation contains mutation information.
	Mutations []*MutationProof `protobuf:"bytes,6,rep,name=mutations" json:"mutations,omitempty"`
	// next_page_token is the next page token to query for pagination.
	// An empty value means there are no more results to fetch.
	// A non-zero value may be used by the client to fetch the next page of
	// results.
	NextPageToken string `protobuf:"bytes,7,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *GetMutationsResponse) Reset()                    { *m = GetMutationsResponse{} }
func (m *GetMutationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsResponse) ProtoMessage()               {}
func (*GetMutationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetMutationsResponse) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetMutationsResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetMutationsResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetMutationsResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

func (m *GetMutationsResponse) GetMutations() []*MutationProof {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func (m *GetMutationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// GetDomainInfoRequest contains an empty request to query the GetDomainInfo
// APIs.
type GetDomainInfoRequest struct {
	// domain_id identifies the domain in which the user and application live.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *GetDomainInfoRequest) Reset()                    { *m = GetDomainInfoRequest{} }
func (m *GetDomainInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoRequest) ProtoMessage()               {}
func (*GetDomainInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetDomainInfoRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

// GetDomainInfoResponse contains the results of GetDomainInfo APIs.
type GetDomainInfoResponse struct {
	// domain_id identifies the domain in which the user and application live.
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,2,opt,name=map" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,3,opt,name=vrf" json:"vrf,omitempty"`
}

func (m *GetDomainInfoResponse) Reset()                    { *m = GetDomainInfoResponse{} }
func (m *GetDomainInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoResponse) ProtoMessage()               {}
func (*GetDomainInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetDomainInfoResponse) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetDomainInfoResponse) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *GetDomainInfoResponse) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *GetDomainInfoResponse) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

// UserProfile is the data that a client would like to store on the server.
type UserProfile struct {
	// data is the public key data for the user.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UserProfile) Reset()                    { *m = UserProfile{} }
func (m *UserProfile) String() string            { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()               {}
func (*UserProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UserProfile) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// GetEpochsRequest is an empty proto message used as input to GetEpochs API.
type GetEpochsRequest struct {
}

func (m *GetEpochsRequest) Reset()                    { *m = GetEpochsRequest{} }
func (m *GetEpochsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEpochsRequest) ProtoMessage()               {}
func (*GetEpochsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// GetEpochsResponse contains mutations of a newly created epoch.
type GetEpochsResponse struct {
	// mutations contains all mutations information of a newly created epoch.
	Mutations *GetMutationsResponse `protobuf:"bytes,1,opt,name=mutations" json:"mutations,omitempty"`
}

func (m *GetEpochsResponse) Reset()                    { *m = GetEpochsResponse{} }
func (m *GetEpochsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEpochsResponse) ProtoMessage()               {}
func (*GetEpochsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetEpochsResponse) GetMutations() *GetMutationsResponse {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.proto.Committed")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.proto.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.proto.Entry")
	proto.RegisterType((*MutationProof)(nil), "keytransparency.v1.proto.MutationProof")
	proto.RegisterType((*MapperMetadata)(nil), "keytransparency.v1.proto.MapperMetadata")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.proto.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.proto.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.proto.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.proto.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.proto.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.proto.UpdateEntryResponse")
	proto.RegisterType((*GetMutationsRequest)(nil), "keytransparency.v1.proto.GetMutationsRequest")
	proto.RegisterType((*GetMutationsResponse)(nil), "keytransparency.v1.proto.GetMutationsResponse")
	proto.RegisterType((*GetDomainInfoRequest)(nil), "keytransparency.v1.proto.GetDomainInfoRequest")
	proto.RegisterType((*GetDomainInfoResponse)(nil), "keytransparency.v1.proto.GetDomainInfoResponse")
	proto.RegisterType((*UserProfile)(nil), "keytransparency.v1.proto.UserProfile")
	proto.RegisterType((*GetEpochsRequest)(nil), "keytransparency.v1.proto.GetEpochsRequest")
	proto.RegisterType((*GetEpochsResponse)(nil), "keytransparency.v1.proto.GetEpochsResponse")
}

func init() {
	proto.RegisterFile("proto/keytransparency_v1_proto/keytransparency_v1_proto.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1098 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4d, 0x6f, 0xdb, 0x36,
	0x18, 0x86, 0xac, 0xc8, 0xb1, 0x5f, 0x3b, 0x49, 0xcb, 0xa6, 0x8d, 0x90, 0xa2, 0x5b, 0xa6, 0x60,
	0x6b, 0x36, 0x0c, 0x0e, 0x92, 0x5e, 0xb6, 0x0e, 0x05, 0xb6, 0xa6, 0x59, 0x1a, 0x34, 0x41, 0x03,
	0xa5, 0xc1, 0x80, 0x5d, 0x04, 0xc6, 0xa2, 0x15, 0x22, 0xb2, 0xc8, 0x90, 0x94, 0x51, 0x07, 0xd8,
	0x75, 0x3b, 0xed, 0x27, 0xec, 0x1f, 0xec, 0x2f, 0xec, 0xb0, 0xdb, 0xee, 0xbb, 0xed, 0xd7, 0x0c,
	0x24, 0x65, 0x5b, 0x76, 0xec, 0x7c, 0xf4, 0xd0, 0x8b, 0x24, 0xbe, 0x1f, 0xd4, 0xcb, 0xe7, 0x79,
	0x5e, 0x92, 0xf0, 0x82, 0x0b, 0xa6, 0xd8, 0xe6, 0x39, 0xe9, 0x2b, 0x81, 0x33, 0xc9, 0xb1, 0x20,
	0x59, 0xbb, 0x1f, 0xf5, 0xb6, 0xa2, 0xeb, 0x1d, 0x2d, 0xf3, 0x44, 0xfe, 0x84, 0xbf, 0xd5, 0xdb,
	0xb2, 0x9e, 0xd5, 0xd5, 0xb6, 0xe8, 0x73, 0x3b, 0x81, 0xe4, 0xa7, 0xc5, 0xab, 0xf0, 0xf9, 0x85,
	0x4f, 0xd2, 0x84, 0x9f, 0xda, 0x67, 0xe1, 0x59, 0x54, 0x82, 0xa6, 0x29, 0xc5, 0x59, 0x31, 0x7e,
	0x34, 0x18, 0x47, 0x5d, 0xcc, 0x23, 0xcc, 0xa9, 0xb5, 0x07, 0x5b, 0x50, 0xdf, 0x61, 0xdd, 0x2e,
	0x55, 0x8a, 0xc4, 0xe8, 0x1e, 0xb8, 0xe7, 0xa4, 0xef, 0x3b, 0x6b, 0xce, 0x46, 0x33, 0xd4, 0x9f,
	0x08, 0xc1, 0x5c, 0x8c, 0x15, 0xf6, 0x2b, 0xc6, 0x64, 0xbe, 0x83, 0xdf, 0x1d, 0x68, 0xec, 0x66,
	0x4a, 0xf4, 0x4f, 0x78, 0x8c, 0x15, 0x41, 0xdf, 0x41, 0xad, 0x9b, 0x2b, 0xac, 0x28, 0xcb, 0x4c,
	0x5c, 0x63, 0xfb, 0xd3, 0xd6, 0xac, 0xd5, 0xb4, 0x4c, 0x62, 0x38, 0x4c, 0x40, 0x3f, 0x40, 0xbd,
	0x3d, 0xf8, 0xbf, 0xef, 0x9a, 0xec, 0xf5, 0xd9, 0xd9, 0xc3, 0x52, 0xc3, 0x51, 0x56, 0xf0, 0x57,
	0x05, 0x3c, 0x33, 0x2d, 0x5a, 0x06, 0x8f, 0x66, 0x31, 0x79, 0x6f, 0x26, 0x6a, 0x86, 0x76, 0x80,
	0x3e, 0x01, 0xb0, 0xc1, 0x5d, 0x92, 0x29, 0xbf, 0x6a, 0x5c, 0x25, 0x0b, 0x7a, 0x0e, 0x4b, 0x38,
	0x57, 0x67, 0x4c, 0xd0, 0x4b, 0x12, 0x47, 0x1a, 0x5f, 0x7f, 0x7e, 0xcd, 0xdd, 0x68, 0x6c, 0xdf,
	0x6f, 0x15, 0x60, 0x1f, 0xe5, 0xa7, 0x29, 0x6d, 0xbf, 0x21, 0xfd, 0x70, 0x71, 0x14, 0xf9, 0x86,
	0xf4, 0x25, 0x5a, 0x85, 0x1a, 0x17, 0xa4, 0x47, 0x59, 0x2e, 0xfd, 0x9a, 0x99, 0x79, 0x38, 0x46,
	0x6f, 0x01, 0x24, 0x4d, 0x32, 0xac, 0x72, 0x41, 0xa4, 0x5f, 0x31, 0x53, 0x6e, 0xde, 0x80, 0x4c,
	0xeb, 0x78, 0x98, 0x61, 0x91, 0x2a, 0x4d, 0xb1, 0x7a, 0x02, 0x4b, 0x13, 0xee, 0x32, 0x63, 0x75,
	0xcb, 0xd8, 0xd7, 0xe0, 0xf5, 0x70, 0x9a, 0x93, 0x82, 0x8a, 0x47, 0x2d, 0xab, 0x8a, 0x57, 0x34,
	0xa1, 0x0a, 0xa7, 0x69, 0x5f, 0xcf, 0x40, 0xe2, 0xd0, 0x06, 0x3d, 0xaf, 0x7c, 0xe3, 0x04, 0xbf,
	0x39, 0xb0, 0x70, 0x58, 0xf0, 0x71, 0x24, 0x18, 0xeb, 0x8c, 0x31, 0xea, 0xdc, 0x95, 0xd1, 0x6f,
	0x01, 0x52, 0x82, 0x3b, 0x5a, 0xdd, 0xac, 0x53, 0x54, 0xb1, 0xda, 0x1a, 0xca, 0xf1, 0x10, 0xf3,
	0x03, 0x82, 0x3b, 0xfb, 0x59, 0x3b, 0xcd, 0x25, 0x65, 0x59, 0x58, 0xd7, 0xd1, 0xe6, 0xbf, 0xc1,
	0x5b, 0x58, 0x3c, 0xc4, 0x9c, 0x13, 0x71, 0x48, 0x14, 0xd6, 0x5a, 0x43, 0x2f, 0xe0, 0xf1, 0x19,
	0x4d, 0xce, 0x88, 0x54, 0x51, 0x27, 0x4f, 0xd3, 0x7e, 0xd4, 0x66, 0x5d, 0x9e, 0x12, 0x45, 0xe2,
	0x48, 0x92, 0x0b, 0x53, 0x9c, 0x1b, 0xfa, 0x45, 0xc8, 0x8f, 0x3a, 0x62, 0x67, 0x10, 0x70, 0x4c,
	0x2e, 0x82, 0x5f, 0x1d, 0x58, 0xda, 0x23, 0xca, 0x96, 0x48, 0x2e, 0x72, 0x22, 0x15, 0x7a, 0x0c,
	0xf5, 0x98, 0x75, 0x31, 0xcd, 0x22, 0x1a, 0xfb, 0x73, 0x06, 0xb8, 0x9a, 0x35, 0xec, 0xc7, 0x68,
	0x05, 0xe6, 0x73, 0x49, 0x84, 0x76, 0x59, 0x4c, 0xab, 0x7a, 0xb8, 0x1f, 0xa3, 0x87, 0x50, 0xc5,
	0x9c, 0x6b, 0x7b, 0xc5, 0xd8, 0x3d, 0xcc, 0xf9, 0x7e, 0x8c, 0xbe, 0x80, 0xa5, 0x0e, 0x15, 0x52,
	0x45, 0x4a, 0x10, 0x12, 0x49, 0x7a, 0x49, 0x8c, 0xf6, 0xdc, 0x70, 0xc1, 0x98, 0xdf, 0x09, 0x42,
	0x8e, 0xe9, 0x25, 0x09, 0xfe, 0xab, 0xc0, 0xbd, 0x51, 0x21, 0x92, 0xb3, 0x4c, 0x12, 0x5d, 0x49,
	0x4f, 0x0c, 0x80, 0xb2, 0x4d, 0x57, 0xeb, 0x09, 0x8b, 0xc5, 0x78, 0x63, 0x54, 0x3e, 0xa4, 0x31,
	0x26, 0x98, 0x70, 0xef, 0xc0, 0x04, 0xfa, 0x12, 0x5c, 0xd9, 0x15, 0x06, 0x9e, 0xc6, 0xf6, 0xca,
	0x28, 0xc7, 0xaa, 0xe7, 0x10, 0xf3, 0x90, 0x31, 0x15, 0xea, 0x18, 0xb4, 0x0d, 0xb5, 0x94, 0x25,
	0x91, 0x60, 0x4c, 0xf9, 0xde, 0xf4, 0xf8, 0x03, 0x96, 0x98, 0xf8, 0xf9, 0xd4, 0x7e, 0xa0, 0xa7,
	0xb0, 0xa4, 0x73, 0xda, 0x2c, 0x93, 0x54, 0x2a, 0xbd, 0x14, 0xbf, 0xba, 0xe6, 0x6e, 0x34, 0xc3,
	0xc5, 0x94, 0x25, 0x3b, 0x23, 0x2b, 0x5a, 0x87, 0x05, 0x1d, 0x48, 0x07, 0x35, 0x9a, 0xce, 0x6c,
	0x86, 0xcd, 0x94, 0x25, 0xc3, 0xba, 0x83, 0xbf, 0x1d, 0x58, 0x39, 0xa0, 0xd2, 0xa2, 0xfb, 0x9a,
	0x4a, 0xc5, 0x66, 0xb0, 0x5d, 0xbd, 0x2d, 0xdb, 0xcb, 0xe0, 0x49, 0x85, 0x85, 0x32, 0xc0, 0xbb,
	0xa1, 0x1d, 0xe8, 0xb9, 0x38, 0x4e, 0x4a, 0x34, 0x7b, 0x61, 0x4d, 0x1b, 0x34, 0xc3, 0x25, 0x81,
	0xcc, 0xdd, 0x20, 0x10, 0x6f, 0x9a, 0x40, 0x7e, 0x01, 0xff, 0xea, 0x12, 0x0a, 0x9d, 0xbc, 0x84,
	0xaa, 0xe9, 0x56, 0xe9, 0x3b, 0x66, 0x13, 0xf9, 0x6a, 0xb6, 0x0e, 0x26, 0x35, 0x16, 0x16, 0x99,
	0xe8, 0x09, 0x40, 0x46, 0xde, 0xab, 0xa8, 0xbc, 0xac, 0xba, 0xb6, 0x1c, 0x6b, 0x43, 0xf0, 0xaf,
	0x03, 0xc8, 0x6e, 0xe7, 0xb3, 0x7b, 0xc5, 0xfb, 0x38, 0xbd, 0x82, 0x5e, 0x43, 0x93, 0xe8, 0x22,
	0xa2, 0xdc, 0x14, 0x54, 0x88, 0xf0, 0xf3, 0x1b, 0x76, 0x20, 0x5b, 0x7d, 0xd8, 0x20, 0xa3, 0x41,
	0xf0, 0x13, 0x3c, 0x18, 0x5b, 0x54, 0x81, 0xe7, 0xf7, 0xe0, 0x8d, 0x7a, 0xee, 0x6e, 0x70, 0xda,
	0xc4, 0xe0, 0x4f, 0x07, 0x1e, 0xec, 0x11, 0x35, 0xd8, 0x35, 0xe5, 0xad, 0xf0, 0x5a, 0x06, 0x8f,
	0x70, 0xd6, 0x3e, 0x2b, 0x76, 0x2d, 0x3b, 0x98, 0x86, 0x4a, 0x65, 0x1a, 0x2a, 0x4f, 0x00, 0x8c,
	0xf8, 0x14, 0x3b, 0x27, 0x99, 0x01, 0xae, 0x1e, 0x1a, 0x39, 0xbe, 0xd3, 0x86, 0x71, 0x6d, 0xce,
	0x8d, 0x6b, 0x33, 0xf8, 0xa7, 0x02, 0xcb, 0xe3, 0xe5, 0x16, 0x48, 0x4c, 0x2f, 0xa9, 0x68, 0xfe,
	0xca, 0x1d, 0x9b, 0xdf, 0xfd, 0xf0, 0xe6, 0x9f, 0xbb, 0x5d, 0xf3, 0x7b, 0x57, 0x9b, 0x1f, 0xed,
	0x42, 0x7d, 0x70, 0xf4, 0x48, 0xb3, 0x89, 0x34, 0xb6, 0x9f, 0xce, 0x26, 0x74, 0xec, 0x9c, 0x0b,
	0x47, 0x99, 0x9a, 0x06, 0xd3, 0x1f, 0x25, 0x8c, 0xe7, 0x0d, 0xc6, 0x0b, 0xda, 0x7c, 0x34, 0xc0,
	0x39, 0x78, 0x66, 0x90, 0x7c, 0x65, 0x39, 0xcd, 0x3a, 0x6c, 0x2a, 0xf3, 0xce, 0x38, 0xf3, 0xc1,
	0x1f, 0x0e, 0x3c, 0x9c, 0xc8, 0x1a, 0x1d, 0x01, 0xb3, 0x0f, 0xa3, 0x35, 0x70, 0x53, 0x96, 0x14,
	0x2a, 0x5d, 0x1c, 0xe1, 0xaa, 0x35, 0x11, 0x6a, 0x97, 0x8e, 0xe8, 0x62, 0x5e, 0x30, 0x75, 0x25,
	0xa2, 0x8b, 0x39, 0x5a, 0x07, 0xb7, 0x27, 0x06, 0x9b, 0xff, 0x94, 0x0b, 0x8d, 0xf6, 0x06, 0x9f,
	0x41, 0xe3, 0x44, 0x12, 0x71, 0x24, 0x58, 0x87, 0xa6, 0x64, 0x78, 0xe9, 0x73, 0x4a, 0x97, 0x3e,
	0x64, 0xcf, 0x2f, 0xad, 0x8f, 0x81, 0xda, 0x03, 0x0c, 0xf7, 0x4b, 0xb6, 0x62, 0x45, 0x07, 0x65,
	0x3e, 0x6c, 0xe9, 0xad, 0x6b, 0x1b, 0xec, 0x8a, 0x2a, 0x4b, 0xb4, 0xbc, 0xdc, 0xfb, 0x79, 0x37,
	0xa1, 0xea, 0x2c, 0x3f, 0x6d, 0xb5, 0x59, 0x77, 0x33, 0x61, 0x2c, 0x49, 0xc9, 0xe4, 0x55, 0x7a,
	0xb3, 0xcd, 0x04, 0xd9, 0xbc, 0xfe, 0x96, 0x7d, 0x5a, 0x35, 0xaf, 0x67, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0xf4, 0x30, 0x20, 0xa7, 0x0b, 0x00, 0x00,
}
