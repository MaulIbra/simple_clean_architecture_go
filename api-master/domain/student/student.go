/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package student

type Student struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
