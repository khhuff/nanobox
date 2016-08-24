package app

import (
  "fmt"
  
  "github.com/jcelliott/lumber"
  
  "github.com/nanobox-io/nanobox/models"
  "github.com/nanobox-io/nanobox/processors/component"
)

// Start will start all services associated with an app
func Start(a *models.App) error {
  locker.LocalLock()
  defer locker.LocalUnlock()
  
  // start all the app components
  if err := component.StartAll(a); err != nil {
    return fmt.Errorf("failed to start app components: %s", err.Error())
  }
  
  // set the status to up
  a.status = "up"
  if err := a.Save(); err != nil {
    lumber.Error("app:Start:models.App.Save(): %s", err.Error())
    return fmt.Errorf("failed to persist app status: %s", err.Error())
  }
  
  return nil
}