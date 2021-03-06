/*
 * easypaisa_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for OAuthProviderErrorEnum enum
 */
type OAuthProviderErrorEnum int

/**
 * Value collection for OAuthProviderErrorEnum enum
 */
const (
    OAuthProviderError_INVALID_REQUEST OAuthProviderErrorEnum = 1 + iota
    OAuthProviderError_INVALID_CLIENT
    OAuthProviderError_INVALID_GRANT
    OAuthProviderError_UNAUTHORIZED_CLIENT
    OAuthProviderError_UNSUPPORTED_GRANT_TYPE
    OAuthProviderError_INVALID_SCOPE
)

func (r OAuthProviderErrorEnum) MarshalJSON() ([]byte, error) { 
    s := OAuthProviderErrorEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *OAuthProviderErrorEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  OAuthProviderErrorEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts OAuthProviderErrorEnum to its string representation
 */
func OAuthProviderErrorEnumToValue(oAuthProviderErrorEnum OAuthProviderErrorEnum) string {
    switch oAuthProviderErrorEnum {
        case OAuthProviderError_INVALID_REQUEST:
    		return "invalid_request"		
        case OAuthProviderError_INVALID_CLIENT:
    		return "invalid_client"		
        case OAuthProviderError_INVALID_GRANT:
    		return "invalid_grant"		
        case OAuthProviderError_UNAUTHORIZED_CLIENT:
    		return "unauthorized_client"		
        case OAuthProviderError_UNSUPPORTED_GRANT_TYPE:
    		return "unsupported_grant_type"		
        case OAuthProviderError_INVALID_SCOPE:
    		return "invalid_scope"		
        default:
        	return "invalid_request"
    }
}

/**
 * Converts OAuthProviderErrorEnum Array to its string Array representation
*/
func OAuthProviderErrorEnumArrayToValue(oAuthProviderErrorEnum []OAuthProviderErrorEnum) []string {
    convArray := make([]string,len( oAuthProviderErrorEnum))
    for i:=0; i<len(oAuthProviderErrorEnum);i++ {
        convArray[i] = OAuthProviderErrorEnumToValue(oAuthProviderErrorEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func OAuthProviderErrorEnumFromValue(value string) OAuthProviderErrorEnum {
    switch value {
        case "invalid_request":
            return OAuthProviderError_INVALID_REQUEST
        case "invalid_client":
            return OAuthProviderError_INVALID_CLIENT
        case "invalid_grant":
            return OAuthProviderError_INVALID_GRANT
        case "unauthorized_client":
            return OAuthProviderError_UNAUTHORIZED_CLIENT
        case "unsupported_grant_type":
            return OAuthProviderError_UNSUPPORTED_GRANT_TYPE
        case "invalid_scope":
            return OAuthProviderError_INVALID_SCOPE
        default:
            return OAuthProviderError_INVALID_REQUEST
    }
}
