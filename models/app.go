package models

import (
	conf "solaris/conf"
	logger "solaris/models/activity"
	en "solaris/models/entities"
	migration "solaris/models/migration"
	solaris "solaris/models/solaris"
)

type Application struct {
	Filials     en.FilialList
	Emloyers    en.EmployerList
	Migration   migration.Migration
	Permissions conf.Permissions
	Logger      logger.ActionLogger
	Solaris     solaris.Solaris
}
