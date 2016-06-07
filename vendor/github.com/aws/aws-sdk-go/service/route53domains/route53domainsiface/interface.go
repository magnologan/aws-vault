// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53domainsiface provides an interface for the Amazon Route 53 Domains.
package route53domainsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

// Route53DomainsAPI is the interface type for route53domains.Route53Domains.
type Route53DomainsAPI interface {
	CheckDomainAvailabilityRequest(*route53domains.CheckDomainAvailabilityInput) (*request.Request, *route53domains.CheckDomainAvailabilityOutput)

	CheckDomainAvailability(*route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)

	DeleteTagsForDomainRequest(*route53domains.DeleteTagsForDomainInput) (*request.Request, *route53domains.DeleteTagsForDomainOutput)

	DeleteTagsForDomain(*route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)

	DisableDomainAutoRenewRequest(*route53domains.DisableDomainAutoRenewInput) (*request.Request, *route53domains.DisableDomainAutoRenewOutput)

	DisableDomainAutoRenew(*route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)

	DisableDomainTransferLockRequest(*route53domains.DisableDomainTransferLockInput) (*request.Request, *route53domains.DisableDomainTransferLockOutput)

	DisableDomainTransferLock(*route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)

	EnableDomainAutoRenewRequest(*route53domains.EnableDomainAutoRenewInput) (*request.Request, *route53domains.EnableDomainAutoRenewOutput)

	EnableDomainAutoRenew(*route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)

	EnableDomainTransferLockRequest(*route53domains.EnableDomainTransferLockInput) (*request.Request, *route53domains.EnableDomainTransferLockOutput)

	EnableDomainTransferLock(*route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)

	GetContactReachabilityStatusRequest(*route53domains.GetContactReachabilityStatusInput) (*request.Request, *route53domains.GetContactReachabilityStatusOutput)

	GetContactReachabilityStatus(*route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error)

	GetDomainDetailRequest(*route53domains.GetDomainDetailInput) (*request.Request, *route53domains.GetDomainDetailOutput)

	GetDomainDetail(*route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)

	GetOperationDetailRequest(*route53domains.GetOperationDetailInput) (*request.Request, *route53domains.GetOperationDetailOutput)

	GetOperationDetail(*route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)

	ListDomainsRequest(*route53domains.ListDomainsInput) (*request.Request, *route53domains.ListDomainsOutput)

	ListDomains(*route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)

	ListDomainsPages(*route53domains.ListDomainsInput, func(*route53domains.ListDomainsOutput, bool) bool) error

	ListOperationsRequest(*route53domains.ListOperationsInput) (*request.Request, *route53domains.ListOperationsOutput)

	ListOperations(*route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)

	ListOperationsPages(*route53domains.ListOperationsInput, func(*route53domains.ListOperationsOutput, bool) bool) error

	ListTagsForDomainRequest(*route53domains.ListTagsForDomainInput) (*request.Request, *route53domains.ListTagsForDomainOutput)

	ListTagsForDomain(*route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)

	RegisterDomainRequest(*route53domains.RegisterDomainInput) (*request.Request, *route53domains.RegisterDomainOutput)

	RegisterDomain(*route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)

	ResendContactReachabilityEmailRequest(*route53domains.ResendContactReachabilityEmailInput) (*request.Request, *route53domains.ResendContactReachabilityEmailOutput)

	ResendContactReachabilityEmail(*route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error)

	RetrieveDomainAuthCodeRequest(*route53domains.RetrieveDomainAuthCodeInput) (*request.Request, *route53domains.RetrieveDomainAuthCodeOutput)

	RetrieveDomainAuthCode(*route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)

	TransferDomainRequest(*route53domains.TransferDomainInput) (*request.Request, *route53domains.TransferDomainOutput)

	TransferDomain(*route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)

	UpdateDomainContactRequest(*route53domains.UpdateDomainContactInput) (*request.Request, *route53domains.UpdateDomainContactOutput)

	UpdateDomainContact(*route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)

	UpdateDomainContactPrivacyRequest(*route53domains.UpdateDomainContactPrivacyInput) (*request.Request, *route53domains.UpdateDomainContactPrivacyOutput)

	UpdateDomainContactPrivacy(*route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)

	UpdateDomainNameserversRequest(*route53domains.UpdateDomainNameserversInput) (*request.Request, *route53domains.UpdateDomainNameserversOutput)

	UpdateDomainNameservers(*route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)

	UpdateTagsForDomainRequest(*route53domains.UpdateTagsForDomainInput) (*request.Request, *route53domains.UpdateTagsForDomainOutput)

	UpdateTagsForDomain(*route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
}

var _ Route53DomainsAPI = (*route53domains.Route53Domains)(nil)
