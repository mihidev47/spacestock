package request

type RegisterFoody struct {
	Name             string `json:"name"`
	PhoneNo          string `json:"msisdn"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Otp              string `json:"otp"`
	Email_Code       string `json:"code"`
	SocialId         string `json:"social_id,omitempty"`
	SocialProviderId string `json:"social_provider_id,omitempty"`
}

type Apartment struct {
	Name             string `json:"name"`
	Address string `json:"address"`
}

type Otp struct {
	Request_Id       string `json:"request_id"`
	Code 		     int `json:"code"`
}

type ReqOtp struct {
	Type       string `json:"type"`
	Unix       string `json:"unique"`
}

type PutFoodyProfile struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   int8   `json:"gender"`
}

type PutChangePassword struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type PostLogin struct {
	Email    string `json:"email"`
	PhoneNo  string `json:"msisdn"`
	Password string `json:"password"`
	DeviceId string `json:"device_id"`
}

type PutUpdateToken struct {
	Type     int8   `json:"type"`
	OldToken string `json:"old"`
	NewToken string `json:"new"`
}

type PostLogout struct {
	TokenType int8   `json:"token_type"`
	Token     string `json:"token"`
}

type ChefProfile struct {
	BusinessName      string  `json:"business_name"`
	BusinessEmail     string  `json:"business_email"`
	Description       string  `json:"description"`
	ShortDescription  string  `json:"short_description"`
	DistrictId        string  `json:"district_id"`
	PostalCode        string  `json:"postal_code"`
	AddressLine       string  `json:"address_line"`
	LastLocationLat   float32 `json:"last_location_lat"`
	LastLocationLng   float32 `json:"last_location_lng"`
	KokiTypeId        string  `json:"koki_type_id"`
	CuisineTypeId     string  `json:"cuisine_type_id"`
	DietTypeId        string  `json:"diet_type_id"`
	MealTypeId        string  `json:"meal_type_id"`
	ChefCategoryId    string  `json:"chef_category_id"`
	ChefExperienceTag string  `json:"chef_experience_tag"`
}

type PostFoodyLoginSocial struct {
	Email            string `json:"email"`
	SocialProviderId int8   `json:"social_provider_id"`
	SocialId         string `json:"social_id"`
	SocialToken      string `json:"social_token"`
}

type PutCourierDeviceActivation struct {
	DeviceId           string `json:"device_id"`
	DeviceIMSI         string `json:"device_imsi"`
	DeviceModel        string `json:"device_model"`
	DeviceManufacturer string `json:"device_manufacturer"`
}

type MenuUpdate struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	ImageFiles           []string `json:"image_files"`
	Price                float32  `json:"price"`
	CuisineTypeId        string   `json:"cuisine_type_id"`
	DietTypeId           string   `json:"diet_type_id"`
	MealTypeId           string   `json:"meal_type_id"`
	CookingTechniqueId   string   `json:"cooking_technique_id"`
	PackagingTypeId      string   `json:"packaging_type_id"`
	ServingTypeId        string   `json:"serving_type_id"`
	ServingTime          int16    `json:"serving_time"`
	HalalState           int8     `json:"halal_state"`
	RecipeDetail         string   `json:"recipe_detail"`
	RecipeScope          int8     `json:"recipe_scope"`
	RecipeTag            string   `json:"recipe_tag"`
	IsSignature          bool     `json:"signature_menu"`
	SignatureDescription string   `json:"signature_description"`
	SearchTag            string   `json:"search_tag"`
}

type MenuReportFlag struct {
	MenuFlagIds []string `json:"chef_menu_flag_ids"`
	Reason      string   `json:"reason"`
}

type ResendVerificationEmail struct {
	Email string `json:"email"`
	Role  int8   `json:"role"`
}

type BankAccount struct {
	Label            string `json:"label"`
	BankId           string `json:"bank_id"`
	BankLocation     string `json:"bank_location"`
	AccountNo        string `json:"account_no"`
	AccountOwnerName string `json:"account_owner_name"`
	Default          bool   `json:"default"`
}

type PushNotification struct {
	TargetId string                  `json:"target_id"`
	RoleId   string                  `json:"role_id"`
	Payload  PushNotificationPayload `json:"payload"`
}

type PushNotificationPayload struct {
	Data map[string]string `json:"data"`
}
