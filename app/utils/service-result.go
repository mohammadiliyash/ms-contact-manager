package utils

// // ServiceResult wraps a result from the model to be rendered in the controller
// type ServiceResult struct {
// 	ID     string `json:"id,omitempty" jsonapi:"-"`
// 	errors map[string]*error
// }

// // Renderable objects can contain a list of errors to be returned with the response
// type Renderable interface {
// 	Errors() []JSONAPIError
// }

// // Errors ...
// func (r ServiceResult) Errors() []JSONAPIError {
// 	errors := []JSONAPIError{}
// 	var err JSONAPIError

// 	for k, v := range r.errors {
// 		key := AddDashes(k)
// 		if v != nil {
// 			thisError := *v
// 			err = JSONAPIError{Detail: thisError.Error(), Status: "422"}
// 			err.Source = "data/attributes/" + key
// 			errors = append(errors, err)
// 		}
// 	}

// 	return errors
// }
