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

// checks if the VacancyCreateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyCreateFields{}

// VacancyCreateFields Поля, передаваемые в запросе на создание вакансии
type VacancyCreateFields struct {
	Area VacancyArea `json:"area"`
	BillingType VacancyBillingType `json:"billing_type"`
	// Описание в html, не менее 200 символов
	Description string `json:"description"`
	// Список требуемых категорий водительских прав
	DriverLicenseTypes []VacancyDriverLicenceTypeItem `json:"driver_license_types,omitempty"`
	Manager NullableVacancyManager `json:"manager,omitempty"`
	// Название
	Name string `json:"name"`
	// Если этот параметр передан, то у новой вакансии дополнительно будет создана связь с предыдущей вакансией (поле previous_id). Этот параметр не влияет на другие и не связан с ними, их всё равно необходимо передавать.  Должен быть равен только ID архивной вакансии. ID архивной вакансии можно получить, запросив [список архивных вакансий](#tag/Upravlenie-vakansiyami/operation/get-archived-vacancies) <a name='previous_id'></a> 
	PreviousId NullableString `json:"previous_id,omitempty"`
	Type VacancyType `json:"type"`
}

type _VacancyCreateFields VacancyCreateFields

// NewVacancyCreateFields instantiates a new VacancyCreateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyCreateFields(area VacancyArea, billingType VacancyBillingType, description string, name string, type_ VacancyType) *VacancyCreateFields {
	this := VacancyCreateFields{}
	this.Area = area
	this.BillingType = billingType
	this.Description = description
	this.Name = name
	this.Type = type_
	return &this
}

// NewVacancyCreateFieldsWithDefaults instantiates a new VacancyCreateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyCreateFieldsWithDefaults() *VacancyCreateFields {
	this := VacancyCreateFields{}
	return &this
}

// GetArea returns the Area field value
func (o *VacancyCreateFields) GetArea() VacancyArea {
	if o == nil {
		var ret VacancyArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetAreaOk() (*VacancyArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *VacancyCreateFields) SetArea(v VacancyArea) {
	o.Area = v
}

// GetBillingType returns the BillingType field value
func (o *VacancyCreateFields) GetBillingType() VacancyBillingType {
	if o == nil {
		var ret VacancyBillingType
		return ret
	}

	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetBillingTypeOk() (*VacancyBillingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingType, true
}

// SetBillingType sets field value
func (o *VacancyCreateFields) SetBillingType(v VacancyBillingType) {
	o.BillingType = v
}

// GetDescription returns the Description field value
func (o *VacancyCreateFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VacancyCreateFields) SetDescription(v string) {
	o.Description = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise.
func (o *VacancyCreateFields) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		var ret []VacancyDriverLicenceTypeItem
		return ret
	}
	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetDriverLicenseTypesOk() ([]VacancyDriverLicenceTypeItem, bool) {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *VacancyCreateFields) HasDriverLicenseTypes() bool {
	if o != nil && !IsNil(o.DriverLicenseTypes) {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given []VacancyDriverLicenceTypeItem and assigns it to the DriverLicenseTypes field.
func (o *VacancyCreateFields) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem) {
	o.DriverLicenseTypes = v
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreateFields) GetManager() VacancyManager {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret VacancyManager
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreateFields) GetManagerOk() (*VacancyManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *VacancyCreateFields) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableVacancyManager and assigns it to the Manager field.
func (o *VacancyCreateFields) SetManager(v VacancyManager) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *VacancyCreateFields) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *VacancyCreateFields) UnsetManager() {
	o.Manager.Unset()
}

// GetName returns the Name field value
func (o *VacancyCreateFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacancyCreateFields) SetName(v string) {
	o.Name = v
}

// GetPreviousId returns the PreviousId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreateFields) GetPreviousId() string {
	if o == nil || IsNil(o.PreviousId.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousId.Get()
}

// GetPreviousIdOk returns a tuple with the PreviousId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreateFields) GetPreviousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousId.Get(), o.PreviousId.IsSet()
}

// HasPreviousId returns a boolean if a field has been set.
func (o *VacancyCreateFields) HasPreviousId() bool {
	if o != nil && o.PreviousId.IsSet() {
		return true
	}

	return false
}

// SetPreviousId gets a reference to the given NullableString and assigns it to the PreviousId field.
func (o *VacancyCreateFields) SetPreviousId(v string) {
	o.PreviousId.Set(&v)
}
// SetPreviousIdNil sets the value for PreviousId to be an explicit nil
func (o *VacancyCreateFields) SetPreviousIdNil() {
	o.PreviousId.Set(nil)
}

// UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil
func (o *VacancyCreateFields) UnsetPreviousId() {
	o.PreviousId.Unset()
}

// GetType returns the Type field value
func (o *VacancyCreateFields) GetType() VacancyType {
	if o == nil {
		var ret VacancyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacancyCreateFields) GetTypeOk() (*VacancyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacancyCreateFields) SetType(v VacancyType) {
	o.Type = v
}

func (o VacancyCreateFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyCreateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["area"] = o.Area
	toSerialize["billing_type"] = o.BillingType
	toSerialize["description"] = o.Description
	if !IsNil(o.DriverLicenseTypes) {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	toSerialize["name"] = o.Name
	if o.PreviousId.IsSet() {
		toSerialize["previous_id"] = o.PreviousId.Get()
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *VacancyCreateFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"area",
		"billing_type",
		"description",
		"name",
		"type",
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

	varVacancyCreateFields := _VacancyCreateFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacancyCreateFields)

	if err != nil {
		return err
	}

	*o = VacancyCreateFields(varVacancyCreateFields)

	return err
}

type NullableVacancyCreateFields struct {
	value *VacancyCreateFields
	isSet bool
}

func (v NullableVacancyCreateFields) Get() *VacancyCreateFields {
	return v.value
}

func (v *NullableVacancyCreateFields) Set(val *VacancyCreateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyCreateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyCreateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyCreateFields(val *VacancyCreateFields) *NullableVacancyCreateFields {
	return &NullableVacancyCreateFields{value: val, isSet: true}
}

func (v NullableVacancyCreateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyCreateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

