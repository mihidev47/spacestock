// Package flags contains constants that is shared across multiple packages for better code readability
package flags

// App Manifest
var AppName = "KOKi Core Service"
var AppVersion = "N/A"
var AppCommitHash = "N/A"

const (
	// Prefix for environment variables
	EnvPrefix = "KOKI"

	// Content Type Headers
	HeaderKeyContentType       = "Content-Type"
	HeaderKeyKOKIAuthorization = "Authorization"
	HeaderKeyKOKIAccessToken   = "X-KOKI-Access-Token"
	HeaderKeyKOKITokenExpired  = "X-KOKI-Token-Expired"
	HeaderKeyKOKISubject       = "X-KOKI-Subject"

	// Content Type Value
	ContentTypeJSON = "application/json; charset=utf-8"
	ContentTypeXML  = "application/xml; charset=utf-8"
	ContentTypeHTML = "text/html; charset=utf-8"
	// Role Mapping Enum
	RolePrincipalTypeUser = "USER"

	// User Status
	UserStatusUnverified = 1
	UserStatusVerified   = 2
	UserStatusBanned     = 10

	// KOKi Type
	ChefTypeUnregistered = "0"
	ChefTypeBintang      = "1"
	ChefTypeLima         = "2"

	// Token Type
	TokenTypeFCM               = 2
	TokenTypeAPNS              = 3
	TokenTypeEmailVerification = 6

	// App ID
	AppFoody   = 1
	AppCourier = 2

	// ACL
	ACLAuthenticatedUser              = ""
	ACLAuthenticatedApp               = "1"
	ACLAuthenticatedAnonymous         = "2"
	ACLEveryone                       = "3"
	ACLPurposeVerifyOTP               = "verify_otp"
	ACLPurposeCourierDeviceActivation = "courier_device_activation"
	ACLPurposeRegister                = "register"
	ACLPurposeResetPassword           = "reset_password"

	// RegistrantStatus
	RegistrantStatusPendingDocument = 1
	RegistrantStatusOnProcess       = 2
	RegistrantStatusApproved        = 3
	RegistrantActivated             = 4
	RegistrantStatusRejected        = 5

	// Country
	CountryIndonesia = "100"

	// File Parameters Keys in HTTP Request
	FileKeyBikePhoto          = "document_bike_photo_file"
	FileKeyCloseUpPhoto       = "document_closeup_photo_file"
	FileKeyDriverLicenseImage = "document_driver_license_file"
	FileKeyDocumentSKCKImage  = "document_skck_file"
	FileKeyDocumentSTNKImage  = "document_stnk_file"
	FileKeyDocumentKKImage    = "document_kk_file"
	FileKeyIdentityCardImage  = "document_identity_card_file"
	FileKeyMotorbikePhoto     = "document_motorbike_photo_file"

	// Documents Type
	DocumentIdentityCard   = 1
	DocumentDriverLicense  = 2
	DocumentBikePhoto      = 3
	DocumentMotorbikePhoto = 4
	DocumentCloseUpPhoto   = 5
	DocumentSKCK           = 6
	DocumentSTNK           = 7
	DocumentKK             = 9

	// Courier Mode
	CourierFootMode      = 1
	CourierBikeMode      = 2
	CourierMotorbikeMode = 3

	// Types
	TypesBank              = "bank"
	TypesChefCategory      = "chef-category"
	TypesChefExperienceTag = "chef-experience-tag"
	TypesChefMenuSearchTag = "chef-menu-search-tag"
	TypesCookingTechnique  = "cooking-technique"
	TypesCuisine           = "cuisine"
	TypesDiet              = "diet"
	TypesChef              = "koki"
	TypesMeal              = "meal"
	TypesPackaging         = "packaging"
	TypesServing           = "serving"
	TypesMenuFlag          = "chef-menu-flag"
	TypesWalletTrxCategory = "wallet-trx-category"

	// Chef Status
	ChefStatusActive = 1
	ChefStatusBanned = 9

	// Resource Config Key
	ResourceAvatar        = "avatar"
	ResourceChefCover     = "chef_cover"
	ResourceCustomerCover = "customer_cover"
	ResourceMenu          = "menu"
	ResourceBanner        = "banner"

	// Filter
	FilterBusinessName     = "filter_business_name"
	FilterChefActive       = "filter_chef_active"
	FilterChefCityId       = "filter_chef_city_id"
	FilterChefId           = "filter_chef_id"
	FilterChefType         = "filter_chef_type_id"
	FilterCityId           = "filter_city_id"
	FilterCuisineType      = "filter_cuisine_id"
	FilterCookingTechnique = "filter_cooking_technique"
	FilterMealType         = "filter_meal_type"
	FilterDietType         = "filter_diet_type"
	FilterFullName         = "filter_full_name"
	FilterHalalState       = "filter_halal_state"
	FilterMenuName         = "filter_menu_name"
	FilterMenuRatings      = "filter_menu_ratings"
	FilterName             = "filter_name"
	FilterPriceMin         = "filter_price_min"
	FilterPriceMax         = "filter_price_max"
	FilterPublishStatus    = "filter_publish_status"
	FilterRadius           = "filter_radius"
	FilterServingType      = "filter_serving_type"
	FilterSignature        = "filter_signature"
	FilterTrxCategory      = "filter_trx_category"
	FilterTrxType          = "filter_trx_type"
	FilterWalletId         = "filter_wallet_id"
	FilterWishlisted       = "filter_wishlisted"
	FilterSince            = "last_updated"

	// Sorting Arrangement
	SortAscending  = "ASC"
	SortDescending = "DESC"

	// Recipe Scope
	RecipeScopePrivate = 0
	RecipeScopePublic  = 1

	// Halal State
	MenuUndefinedHalal = 0
	MenuHalal          = 1
	MenuNonHalal       = 2

	// Published
	MenuUnpublished = 0
	MenuPublished   = 1
	MenuSuspended   = 2
	MenuDrafted     = 3

	// Wallet Transaction Type
	WalletTrxInit = 1
	WalletTrxIn   = 2
	WalletTrxOut  = 3

	// Wallet Transaction Status
	WalletTrxSuccess = 1
	WalletTrxFailed  = 2
	WalletTrxPending = 3

	// Wallet Transfer Status
	WalletTrfRequested = 1
	WalletTrfOnProcess = 2
	WalletTrfSuccess   = 3
	WalletTrfPending   = 6
	WalletTrfFailed    = 7
	WalletTrfCanceled  = 8
	WalletTrfExpired   = 9

	// Wallet Transfer Type
	WalletTrfTopup    = 1
	WalletTrfWithdraw = 2

	// Wallet Transaction Category
	WalletTrxCategoryTopup = "1"

	// Payment Methods
	PaymentMethodCreditCard  = 2
	PaymentMethodPermataVA   = 4
	PaymentMethodKOKICredits = 5
	PaymentMethodAlfamart    = 6
	PaymentMethodDefault     = "KOiNKU"

	// Delivery Service
	DeliveryServiceKOKOCourier = 1

	// Transaction ID
	TrxMaxDigit      = 10
	OrderIDPrefix    = "OR-"
	DeliveryIDPrefix = "DL-"
)

// Courier Discovery Status
const (
	_ = iota
	CourierDiscoveryWaiting
	CourierDiscoveryNotified
	CourierDiscoveryAccepted
	CourierDiscoveryRejected
	CourierDiscoveryExpired
	CourierDiscoveryBusy
	CourierDiscoveryIgnored
)

// Courier Delivery Status
const (
	DeliveryStatusRequest = iota + 1
	DeliveryStatusFindCourier
	DeliveryStatusCourierAccept
	DeliveryStatusPreparePickUp
	DeliveryStatusPickUp
	DeliveryStatusDeliver
	DeliveryStatusDelivered
	_
	_
	DeliveryStatusSuccess
	DeliveryCreated
)

// Courier Delivery Status for Request. Starts from 20
const (
	DeliveryStatusCourierRequestCancel = iota + 20
	DeliveryStatusSenderRequestCancel
)

// Courier Delivery Status Canceled. Starts from 30
const (
	DeliveryAdminCancel = iota + 30
	DeliverySenderCancel
)

// Courier Delivery Status for Internal. Starts from 90
const (
	DeliveryStatusCourierNotFound = iota + 90
	DeliveryStatusCourierNotRespond
	DeliveryReceiverCancel
	DeliveryStatusInternalServerError = 99
)

// Courier Bid Status
const (
	BidInactive = iota
	BidActive
	BidBusy
)

// Queue
const (
	// Sync Jobs
	QueueSyncAllStats = "cron.sync_all_stats"
	// Sync Event
	QueueEventDeliveryReview = "event.delivery_review"
	QueueEventOrderDone      = "event.order_done"
	QueueEventOrderReview    = "event.order_review"
	// Courier
	QueueCourierDiscover       = "courier.discover"
	QueueCourierNewDelivery    = "courier.new_delivery"
	QueueCourierEndRequest     = "courier.end_delivery_request"
	QueueCourierSettlementSync = "courier.settlement.sync"
	QueueCourierStatsSync      = "couriers.sync.stats"
	// Dish
	QueueDishLikesSync = "dish.likes.sync"
	QueueDishStatsSync = "dish.sync.all"
	// Exchange
	ExchangeCourierDelayed = "courier.delayed"
	// Order Queue
	QueueOrderCancelCourier = "order.cancel_by_courier"
	// Settlement Queue
	QueueMerchantOrderAutoAccept = "merchant.order.auto_accept"
	QueueMerchantSettlementSync  = "merchant.settlement.sync"
	QueueMerchantFollowersSync   = "merchant.followers.sync"
	QueueMerchantStatsSync       = "merchant.sync.stats"
	// User
	QueueUserNotifyFCM = "user.notify_fcm"
	// Routing Key
	RoutingNewDelivery = "new_delivery"
	RoutingEndRequest  = "end_request"
)

// Order Status
const (
	OrderPlaced = iota + 1
	OrderFindCourier
	OrderCourierAccept
	OrderMerchantAccept
	OrderCourierPickUp
	OrderCourierDeliver
	OrderCourierDelivered
	_
	_
	OrderSuccess
)

// Order Status for Cancel
const (
	OrderAdminCancel = iota + 30
	OrderMerchantReject
)

// Order Status for Internal
const (
	OrderCancelNoCourier = iota + 90
	OrderCourierNotRespond
	OrderCancelFindCourier
	OrderCancelInternalError = 99
)

// User role
const (
	RoleFoody = iota + 10
	RoleMerchant
	_
	RoleCourier
	RoleAdministrator = 21
)

// city type
const (
	_ = iota
	LocationTypeCity
	LocationTypeRegency
)

// Wallet Transaction Category
const (
	TrxUncategorized = iota
	TrxCategoryTopup
	TrxCategoryWithdraw
	TrxCategoryTransfer
	TrxCategoryFoodBeverages
)

// chat const
const (
	ChatWsTypeTracking            = 1
	ChatWsTypeChatting            = 2
	ChatStatusSending             = 1
	ChatStatusSent                = 2
	ChatStatusRead                = 3
	ChatStatusFailed              = 4
	ChatEventTyping               = 1
	ChatEventChatting             = 2
	ChatBetweenMerchantAndCourier = "merchant_courier_chat"
	ChatBetweenCustomerAndCourier = "customer_courier_chat"
)

const (
	SyncBatchLimit      int64 = 500
	SettlementItemLimit int8  = 50
)

// Settlement Status
const (
	SettlementRequested = iota + 1
	SettlementOnProcess
	SettlementApproved
	SettlementPaymentRequested
	SettlementPaymentPending
	SettlementPaymentFailed
	SettlementSuccess  = 10
	SettlementRejected = 20
)

const (
	SettlementItemRequested = iota + 1
	SettlementItemApproved
	SettlementItemRejected
)

// Stats Context
const (
	StatsOrderDone                 = 10
	StatsOrderSuccess              = 11
	StatsOrderFailed               = 12
	StatsOrderUniqueCustomer       = 14
	StatsOrderSuccessQty           = 15
	StatsDeliveryDone              = 20
	StatsDeliverySuccess           = 21
	StatsDeliveryFailed            = 22
	StatsMerchantFollowers         = 30
	StatsMerchantRating            = 31
	StatsMerchantRatedDish         = 32
	StatsDishReview                = 42
	StatsDishAvgRating             = 43
	StatsCourierAvgRating          = 60
	StatsCourierFeedbackByCustomer = 61
)

const (
	// Sync All Context
	SyncAllDishes = iota + 1
	SyncAllMerchants
	SyncAllCouriers
)

var CourierVehicle = map[string]string{
	"1": "KOKO Kaki",
	"2": "KOKO Speda",
	"3": "KOKO Motor",
}
