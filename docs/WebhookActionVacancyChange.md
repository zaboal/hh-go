# WebhookActionVacancyChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**NullableWebhookActionVacancyOnlyMineSettings**](WebhookActionVacancyOnlyMineSettings.md) |  | [optional] 
**Type** | **string** | Изменение вакансии. Аккумулирует изменения, внесенные за несколько последних секунд, и отправляет вебхук, содержащий время последнего изменения.  Если вы внесете два изменения с разницей в одну секунду, сервис отправит только один вебхук, который будет содержать время последнего изменения. Если изменение одно, сервис отправит вебхук с задержкой в несколько секунд  | 

## Methods

### NewWebhookActionVacancyChange

`func NewWebhookActionVacancyChange(type_ string, ) *WebhookActionVacancyChange`

NewWebhookActionVacancyChange instantiates a new WebhookActionVacancyChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookActionVacancyChangeWithDefaults

`func NewWebhookActionVacancyChangeWithDefaults() *WebhookActionVacancyChange`

NewWebhookActionVacancyChangeWithDefaults instantiates a new WebhookActionVacancyChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *WebhookActionVacancyChange) GetSettings() WebhookActionVacancyOnlyMineSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebhookActionVacancyChange) GetSettingsOk() (*WebhookActionVacancyOnlyMineSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebhookActionVacancyChange) SetSettings(v WebhookActionVacancyOnlyMineSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WebhookActionVacancyChange) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *WebhookActionVacancyChange) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *WebhookActionVacancyChange) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetType

`func (o *WebhookActionVacancyChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookActionVacancyChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookActionVacancyChange) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


