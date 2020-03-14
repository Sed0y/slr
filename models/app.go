package models

import (
	conf "dpm/conf"
	logger "dpm/models/activity"
	en "dpm/models/entities"
	migration "dpm/models/migration"
	solaris "dpm/models/solaris"
)

type Application struct {
	Filials     en.FilialList
	Emloyers    en.EmployerList
	Migration   migration.Migration
	Permissions conf.Permissions
	Logger      logger.ActionLogger
	Solaris     solaris.Solaris
}
