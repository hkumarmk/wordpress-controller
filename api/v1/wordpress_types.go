/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WordpressSpec defines the desired state of Wordpress
type WordpressSpec struct {
	//Number of wordpress instances
	//+optional
	Replicas *uint32 `json:"replicas,omitempty"`

	//List of domains
	Domains []string `json:"domains"`

	//Custom image to be used for wordpress
	//+optional
	Image string `json:"image,omitempty"`

	//custom image tag for wordpress image
	//+optional
	Tag string `json:"tag,omitempty"`

	//Wordpress website username
	//+optional
	Username string `json:"username,omitempty"`

	//Wordpress website user password
	//+optional
	Password string `json:"password,omitempty"`

	//Wordpress website user email
	//+optional
	Email string `json:"email,omitempty"`

	//Wordpress website title
	//+optional
	Title string `json:"title,omitempty"`

	//DB host
	DBHost string `json:"dbhost"`

	//DB name
	DBName string `json:"dbname"`

	//DB User
	DBUser string `json:"dbuser"`

	//DB Password
	DBPassword string `json:"dbpassword"`
}

// WordpressStatus defines the observed state of Wordpress
type WordpressStatus struct {
	//Number of Wordpress instances
	Replicas *uint32 `json:"replicas"`
}

// +kubebuilder:object:root=true

// Wordpress is the Schema for the wordpresses API
type Wordpress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WordpressSpec   `json:"spec,omitempty"`
	Status WordpressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WordpressList contains a list of Wordpress
type WordpressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wordpress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Wordpress{}, &WordpressList{})
}
