/* auto-generated */

package env

type EnvName string

type EnvVar struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type EnvCollection []EnvVar
