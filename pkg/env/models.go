/* auto-generated */

package env

type EnvName string

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type EnvCollection []EnvVar
