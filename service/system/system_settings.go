package system

type SettingsService struct{}

var SettingsServiceApp = new(SettingsService)

func (settingsService *SettingsService) GetSettings() (err error) {
	return nil
}
