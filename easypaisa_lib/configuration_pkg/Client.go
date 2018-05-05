/*
 * easypaisa_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package configuration_pkg

import(
	"easypaisa_lib/apihelper_pkg"

/* Setting up enums for Environments and Servers 
*/
type Environments int

type Servers int

const (
     PRODUCTION Environments = 1 + iota
)
const (
 	 DEFAULT Servers = 1 + iota
   	 AUTH
)

type CONFIGURATION_IMPL struct {
/** OAuth 2 Client ID */
    /** Replace the value of oauthclientid with SetOAuthClientId function */
    oauthclientid string
/** OAuth 2 Client Secret */
    /** Replace the value of oauthclientsecret with SetOAuthClientSecret function */
    oauthclientsecret string
}
 
  
/*
 * Getter function returning oauthclientid
 */
func (me *CONFIGURATION_IMPL) OAuthClientId() string{
     return me.oauthclientid
}

/*
 * Getter function returning oauthclientsecret
 */
func (me *CONFIGURATION_IMPL) OAuthClientSecret() string{
     return me.oauthclientsecret
}

/*
 * Setter function setting up oauthclientid
 */
func (me *CONFIGURATION_IMPL) SetOAuthClientId(oAuthClientId string) {
      me.oauthclientid = oAuthClientId
}

/*
 * Setter function setting up oauthclientsecret
 */
func (me *CONFIGURATION_IMPL) SetOAuthClientSecret(oAuthClientSecret string) {
      me.oauthclientsecret = oAuthClientSecret
}

// Setting up Default Environment
var Environment = PRODUCTION

//A map of environments and their corresponding servers/baseurls
 var EnvironmentsMap = map[Environments](map[Servers]string){

    PRODUCTION : map[Servers]string{
        DEFAULT:"https://api.telenor.com.pk",
        AUTH:"https://api.telenor.com.pk",
    },
}
 
// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper_pkg.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
