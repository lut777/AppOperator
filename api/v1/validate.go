package v1

import (
	"errors"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"regexp"
)

var _maxreplica int32 = 10

func Validation(r *AppOperator, log logr.Logger) error {

	if !checkUsrname(r.Spec.User) {
		log.Info("misformatted user name", r.Spec.User, r.Name)
		return errors.New("misformatted user name")
	}

	if !checkType(r.Spec.Type) {
		log.Info("misformatted type", r.Spec.Type, r.Name)
		return errors.New("misformatted type")
	}

	if !checkMaxBackups(r.Spec.MaxBackups) {
		log.Info("misformatted maxBackups", r.Spec.MaxBackups, r.Name)
		return errors.New("misformatted maxBackups")
	}

	if !checkAddress(r.Spec.Address) {
		log.Info("misformatted address", r.Spec.Address, r.Name)
		return errors.New("misformatted address")
	}

	ok, err := checkResourceRequirements(r.Spec.ResourceRequirements)
	if !ok {
		return err
	}

	if ok, err := checkStorageSize(r.Spec.StorageSize, r.Spec.ResourceRequirements.Requests[corev1.ResourceMemory]); !ok {
		return err
	}
	return nil
}

func checkUsrname(usr string) bool {
	if len(usr) == 0 {
		return true
	}
	match, _ := regexp.MatchString("^([0-9a-zA-Z]+?)_([0-9a-zA-Z]+?)$", usr)
	return match
}

func checkType(typ string) bool {
	switch typ {
	case "full":
		return true
	case "partial":
		return true
	case "":
		return true
	default:
		return false
	}
}

func checkMaxBackups(maxb int32) bool {
	return 1 <= maxb && maxb <= _maxreplica
}

func checkAddress(address string) bool {
	match, _ := regexp.MatchString("^(http|https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$", address)
	return match
}

func checkResourceRequirements(req corev1.ResourceRequirements) (bool, error) {
	reqstore, ok := req.Requests[corev1.ResourceMemory]
	if !ok {
		return false, errors.New("request memory not requested")
	}
	if _, ok := req.Limits[corev1.ResourceMemory]; ok && req.Limits.Memory().Cmp(reqstore) < 0 {
		return false, errors.New("limit memory smaller than request")
	}

	return true, nil
}

func checkStorageSize(size string, cmp resource.Quantity) (bool, error) {
	if len(size) == 0 {
		return true, nil
	}

	base, err := resource.ParseQuantity(size)
	if err != nil {
		return false, errors.New("misformatted storage")
	}

	if base.Cmp(cmp) < 0 {
		return false, errors.New("storageSize smaller than request memory")
	}

	return true, nil
}
