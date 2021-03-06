/*
 * easypaisa_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package configuration_pkg



type CONFIGURATION interface {
        OAuthClientId()  string
        OAuthClientSecret()  string
        SetOAuthClientId(oAuthClientId   string)
        SetOAuthClientSecret(oAuthClientSecret   string)
}   
/*
 * Factory for the CONFIGURATION interaface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION() CONFIGURATION{
    configuration := new(CONFIGURATION_IMPL)
    return configuration
}