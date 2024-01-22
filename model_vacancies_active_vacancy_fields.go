/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/zaboal/hh-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VacanciesActiveVacancyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesActiveVacancyFields{}

// VacanciesActiveVacancyFields struct for VacanciesActiveVacancyFields
type VacanciesActiveVacancyFields struct {
	Address NullableVacanciesAddress `json:"address,omitempty"`
	// Можно ли улучшить биллинговый тип вакансии
	CanUpgradeBillingType bool `json:"can_upgrade_billing_type"`
	Counters VacancyCountersForActive `json:"counters"`
	// Дата и время публикации вакансии
	CreatedAt string `json:"created_at"`
	// Дата и время окончания публикации вакансии
	ExpiresAt string `json:"expires_at"`
	// Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания
	HasUpdates bool `json:"has_updates"`
	Manager VacancyManagerOutput `json:"manager"`
	// Расстояние в метрах между центром сортировки (заданной параметрами `sort_point_lat`, `sort_point_lng`) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение `sort_point_distance` выдается только в случае, если заданы параметры `sort_point_lat`, `sort_point_lng`, `order_by=distance` 
	SortPointDistance NullableFloat32 `json:"sort_point_distance,omitempty"`
}

type _VacanciesActiveVacancyFields VacanciesActiveVacancyFields

// NewVacanciesActiveVacancyFields instantiates a new VacanciesActiveVacancyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesActiveVacancyFields(canUpgradeBillingType bool, counters VacancyCountersForActive, createdAt string, expiresAt string, hasUpdates bool, manager VacancyManagerOutput) *VacanciesActiveVacancyFields {
	this := VacanciesActiveVacancyFields{}
	this.CanUpgradeBillingType = canUpgradeBillingType
	this.Counters = counters
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.HasUpdates = hasUpdates
	this.Manager = manager
	return &this
}

// NewVacanciesActiveVacancyFieldsWithDefaults instantiates a new VacanciesActiveVacancyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesActiveVacancyFieldsWithDefaults() *VacanciesActiveVacancyFields {
	this := VacanciesActiveVacancyFields{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesActiveVacancyFields) GetAddress() VacanciesAddress {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacanciesAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesActiveVacancyFields) GetAddressOk() (*VacanciesAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesActiveVacancyFields) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacanciesAddress and assigns it to the Address field.
func (o *VacanciesActiveVacancyFields) SetAddress(v VacanciesAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesActiveVacancyFields) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesActiveVacancyFields) UnsetAddress() {
	o.Address.Unset()
}

// GetCanUpgradeBillingType returns the CanUpgradeBillingType field value
func (o *VacanciesActiveVacancyFields) GetCanUpgradeBillingType() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanUpgradeBillingType
}

// GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetCanUpgradeBillingTypeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanUpgradeBillingType, true
}

// SetCanUpgradeBillingType sets field value
func (o *VacanciesActiveVacancyFields) SetCanUpgradeBillingType(v bool) {
	o.CanUpgradeBillingType = v
}

// GetCounters returns the Counters field value
func (o *VacanciesActiveVacancyFields) GetCounters() VacancyCountersForActive {
	if o == nil {
		var ret VacancyCountersForActive
		return ret
	}

	return o.Counters
}

// GetCountersOk returns a tuple with the Counters field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetCountersOk() (*VacancyCountersForActive, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counters, true
}

// SetCounters sets field value
func (o *VacanciesActiveVacancyFields) SetCounters(v VacancyCountersForActive) {
	o.Counters = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VacanciesActiveVacancyFields) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VacanciesActiveVacancyFields) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *VacanciesActiveVacancyFields) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *VacanciesActiveVacancyFields) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetHasUpdates returns the HasUpdates field value
func (o *VacanciesActiveVacancyFields) GetHasUpdates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasUpdates
}

// GetHasUpdatesOk returns a tuple with the HasUpdates field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetHasUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasUpdates, true
}

// SetHasUpdates sets field value
func (o *VacanciesActiveVacancyFields) SetHasUpdates(v bool) {
	o.HasUpdates = v
}

// GetManager returns the Manager field value
func (o *VacanciesActiveVacancyFields) GetManager() VacancyManagerOutput {
	if o == nil {
		var ret VacancyManagerOutput
		return ret
	}

	return o.Manager
}

// GetManagerOk returns a tuple with the Manager field value
// and a boolean to check if the value has been set.
func (o *VacanciesActiveVacancyFields) GetManagerOk() (*VacancyManagerOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manager, true
}

// SetManager sets field value
func (o *VacanciesActiveVacancyFields) SetManager(v VacancyManagerOutput) {
	o.Manager = v
}

// GetSortPointDistance returns the SortPointDistance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesActiveVacancyFields) GetSortPointDistance() float32 {
	if o == nil || IsNil(o.SortPointDistance.Get()) {
		var ret float32
		return ret
	}
	return *o.SortPointDistance.Get()
}

// GetSortPointDistanceOk returns a tuple with the SortPointDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesActiveVacancyFields) GetSortPointDistanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortPointDistance.Get(), o.SortPointDistance.IsSet()
}

// HasSortPointDistance returns a boolean if a field has been set.
func (o *VacanciesActiveVacancyFields) HasSortPointDistance() bool {
	if o != nil && o.SortPointDistance.IsSet() {
		return true
	}

	return false
}

// SetSortPointDistance gets a reference to the given NullableFloat32 and assigns it to the SortPointDistance field.
func (o *VacanciesActiveVacancyFields) SetSortPointDistance(v float32) {
	o.SortPointDistance.Set(&v)
}
// SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil
func (o *VacanciesActiveVacancyFields) SetSortPointDistanceNil() {
	o.SortPointDistance.Set(nil)
}

// UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
func (o *VacanciesActiveVacancyFields) UnsetSortPointDistance() {
	o.SortPointDistance.Unset()
}

func (o VacanciesActiveVacancyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesActiveVacancyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["can_upgrade_billing_type"] = o.CanUpgradeBillingType
	toSerialize["counters"] = o.Counters
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["has_updates"] = o.HasUpdates
	toSerialize["manager"] = o.Manager
	if o.SortPointDistance.IsSet() {
		toSerialize["sort_point_distance"] = o.SortPointDistance.Get()
	}
	return toSerialize, nil
}

func (o *VacanciesActiveVacancyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"can_upgrade_billing_type",
		"counters",
		"created_at",
		"expires_at",
		"has_updates",
		"manager",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVacanciesActiveVacancyFields := _VacanciesActiveVacancyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesActiveVacancyFields)

	if err != nil {
		return err
	}

	*o = VacanciesActiveVacancyFields(varVacanciesActiveVacancyFields)

	return err
}

type NullableVacanciesActiveVacancyFields struct {
	value *VacanciesActiveVacancyFields
	isSet bool
}

func (v NullableVacanciesActiveVacancyFields) Get() *VacanciesActiveVacancyFields {
	return v.value
}

func (v *NullableVacanciesActiveVacancyFields) Set(val *VacanciesActiveVacancyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesActiveVacancyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesActiveVacancyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesActiveVacancyFields(val *VacanciesActiveVacancyFields) *NullableVacanciesActiveVacancyFields {
	return &NullableVacanciesActiveVacancyFields{value: val, isSet: true}
}

func (v NullableVacanciesActiveVacancyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesActiveVacancyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

