package webhook

import (
	"strings"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JsonPatchEntry struct {
	OP    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
	From  string      `json:"from,omitempty"`
}

func hasContainer(containers []corev1.Container, prefix string) (string, bool) {
	for _, container := range containers {
		if strings.HasPrefix(container.Name, prefix) {
			return container.Name, true
		}
	}
	return "", false
}

func filterMapValuesByKey(values map[string]string, key string) ([]string, bool) {

	value, ok := values[key]
	if ok {
		if len(value) > 0 {
			return strings.Split(value, ","), true
		}
	}
	return []string{}, false
}

// failureResponse sets the admission response as a failure with the provided error.
//
// Args:
//
//	ar: The admissionv1.AdmissionResponse to be updated.
//	err: The error that occurred, which will be logged and included in the response reason.
//
// Returns:
//
//	A pointer to the updated admissionv1.AdmissionResponse with Allowed set to false,
//	and the Result status set to "Failure" with a reason including the error message.
func failureResponse(ar admissionv1.AdmissionResponse, err error) *admissionv1.AdmissionResponse {
	ar.Allowed = false
	ar.Result = &metav1.Status{
		Status:  "Failure",
		Message: err.Error(),
	}
	return &ar
}

// successResponse sets the admission response as a success.
//
// Args:
//
//	ar: The admissionv1.AdmissionResponse to be updated.
//
// Returns:
//
//	A pointer to the updated admissionv1.AdmissionResponse with Allowed set to true,
//	and the Result status set to "Success".
func successResponse(ar admissionv1.AdmissionResponse) *admissionv1.AdmissionResponse {
	ar.Allowed = true
	ar.Result = &metav1.Status{
		Status: "Success",
	}
	return &ar
}
