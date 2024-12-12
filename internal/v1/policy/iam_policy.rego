package iam.v1

import future.keywords.in

default allow = false


# default permssions that check token, aud, account_id and permissions
default defaul_permissions_methods := []

# public method list
default public_methods := [
	"RequestOTP",
	"VerifyOTP",
]

# method with just login requried
default method_just_loogedin := []

# method with just longtoken requried
default method_just_longtoken := [
	"CreateProfile",
]


method 	:= input.path
user  	:= input.paseto
request := input.request
internal_account_id := "88888888"
longterm := "longterm"
shortterm := "shortterm"


# permissions  line 26  method == user.scopes[_]  user.user_id == request.user_id
allow {
	method == defaul_permissions_methods[_]
	shortterm == user.ClaimType
}

# permissions    method == defaul_permissions_methods[_] method == user.scopes[_]
allow {
	method == method_just_longtoken[_]
	longterm == user.ClaimType
}

# public methods
allow {
	method == public_methods[_]
}

# method with no input put user must be looged in
allow {
	method == method_just_loogedin[_]
	user.account_id # just check if logged in user exists
}

# Management endpoints
# changeVendorStatus, list_vendors, changepluginstatus, list_all_plugins, publish_plugin