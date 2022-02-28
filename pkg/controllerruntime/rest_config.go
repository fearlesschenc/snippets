package controllerruntime

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
)

func GetRestConfig() {
	config := ctrl.GetConfigOrDie()
	fmt.Println(config.Host)
}
