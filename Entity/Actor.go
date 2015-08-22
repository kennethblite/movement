package Entity

import()

//Actor will be used for anything that depends on a key press
type Actor interface{
Entity
func()
}
//Register and Unregister the event with the master Listener
