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

// checks if the EmployerManagersAddEmployerManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployerManagersAddEmployerManager{}

// EmployerManagersAddEmployerManager struct for EmployerManagersAddEmployerManager
type EmployerManagersAddEmployerManager struct {
	AdditionalPhone *EmployerManagersManagerDataAdditionalPhone `json:"additional_phone,omitempty"`
	Area EmployerManagersAreaId `json:"area"`
	// Адрес электронной почты менеджера
	Email string `json:"email"`
	// Имя менеджера
	FirstName string `json:"first_name"`
	// Является ли менеджер основным контактным лицом
	IsMainContactPerson bool `json:"is_main_contact_person"`
	// Фамилия менеджера
	LastName string `json:"last_name"`
	ManagerType EmployerManagersManagerTypeId `json:"manager_type"`
	// Отчество менеджера
	MiddleName *string `json:"middle_name,omitempty"`
	// Список [прав менеджера](#tag/Menedzhery-rabotodatelya/operation/get-employer-manager-types)
	Permissions []EmployerManagersPermissions `json:"permissions,omitempty"`
	Phone EmployerManagersManagerDataPhone `json:"phone"`
	// Должность менеджера
	Position string `json:"position"`
}

type _EmployerManagersAddEmployerManager EmployerManagersAddEmployerManager

// NewEmployerManagersAddEmployerManager instantiates a new EmployerManagersAddEmployerManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerManagersAddEmployerManager(area EmployerManagersAreaId, email string, firstName string, isMainContactPerson bool, lastName string, managerType EmployerManagersManagerTypeId, phone EmployerManagersManagerDataPhone, position string) *EmployerManagersAddEmployerManager {
	this := EmployerManagersAddEmployerManager{}
	this.Area = area
	this.Email = email
	this.FirstName = firstName
	this.IsMainContactPerson = isMainContactPerson
	this.LastName = lastName
	this.ManagerType = managerType
	this.Phone = phone
	this.Position = position
	return &this
}

// NewEmployerManagersAddEmployerManagerWithDefaults instantiates a new EmployerManagersAddEmployerManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerManagersAddEmployerManagerWithDefaults() *EmployerManagersAddEmployerManager {
	this := EmployerManagersAddEmployerManager{}
	return &this
}

// GetAdditionalPhone returns the AdditionalPhone field value if set, zero value otherwise.
func (o *EmployerManagersAddEmployerManager) GetAdditionalPhone() EmployerManagersManagerDataAdditionalPhone {
	if o == nil || IsNil(o.AdditionalPhone) {
		var ret EmployerManagersManagerDataAdditionalPhone
		return ret
	}
	return *o.AdditionalPhone
}

// GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetAdditionalPhoneOk() (*EmployerManagersManagerDataAdditionalPhone, bool) {
	if o == nil || IsNil(o.AdditionalPhone) {
		return nil, false
	}
	return o.AdditionalPhone, true
}

// HasAdditionalPhone returns a boolean if a field has been set.
func (o *EmployerManagersAddEmployerManager) HasAdditionalPhone() bool {
	if o != nil && !IsNil(o.AdditionalPhone) {
		return true
	}

	return false
}

// SetAdditionalPhone gets a reference to the given EmployerManagersManagerDataAdditionalPhone and assigns it to the AdditionalPhone field.
func (o *EmployerManagersAddEmployerManager) SetAdditionalPhone(v EmployerManagersManagerDataAdditionalPhone) {
	o.AdditionalPhone = &v
}

// GetArea returns the Area field value
func (o *EmployerManagersAddEmployerManager) GetArea() EmployerManagersAreaId {
	if o == nil {
		var ret EmployerManagersAreaId
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetAreaOk() (*EmployerManagersAreaId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *EmployerManagersAddEmployerManager) SetArea(v EmployerManagersAreaId) {
	o.Area = v
}

// GetEmail returns the Email field value
func (o *EmployerManagersAddEmployerManager) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmployerManagersAddEmployerManager) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *EmployerManagersAddEmployerManager) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EmployerManagersAddEmployerManager) SetFirstName(v string) {
	o.FirstName = v
}

// GetIsMainContactPerson returns the IsMainContactPerson field value
func (o *EmployerManagersAddEmployerManager) GetIsMainContactPerson() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMainContactPerson
}

// GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetIsMainContactPersonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMainContactPerson, true
}

// SetIsMainContactPerson sets field value
func (o *EmployerManagersAddEmployerManager) SetIsMainContactPerson(v bool) {
	o.IsMainContactPerson = v
}

// GetLastName returns the LastName field value
func (o *EmployerManagersAddEmployerManager) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *EmployerManagersAddEmployerManager) SetLastName(v string) {
	o.LastName = v
}

// GetManagerType returns the ManagerType field value
func (o *EmployerManagersAddEmployerManager) GetManagerType() EmployerManagersManagerTypeId {
	if o == nil {
		var ret EmployerManagersManagerTypeId
		return ret
	}

	return o.ManagerType
}

// GetManagerTypeOk returns a tuple with the ManagerType field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetManagerTypeOk() (*EmployerManagersManagerTypeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagerType, true
}

// SetManagerType sets field value
func (o *EmployerManagersAddEmployerManager) SetManagerType(v EmployerManagersManagerTypeId) {
	o.ManagerType = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *EmployerManagersAddEmployerManager) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *EmployerManagersAddEmployerManager) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *EmployerManagersAddEmployerManager) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *EmployerManagersAddEmployerManager) GetPermissions() []EmployerManagersPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret []EmployerManagersPermissions
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetPermissionsOk() ([]EmployerManagersPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *EmployerManagersAddEmployerManager) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []EmployerManagersPermissions and assigns it to the Permissions field.
func (o *EmployerManagersAddEmployerManager) SetPermissions(v []EmployerManagersPermissions) {
	o.Permissions = v
}

// GetPhone returns the Phone field value
func (o *EmployerManagersAddEmployerManager) GetPhone() EmployerManagersManagerDataPhone {
	if o == nil {
		var ret EmployerManagersManagerDataPhone
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetPhoneOk() (*EmployerManagersManagerDataPhone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *EmployerManagersAddEmployerManager) SetPhone(v EmployerManagersManagerDataPhone) {
	o.Phone = v
}

// GetPosition returns the Position field value
func (o *EmployerManagersAddEmployerManager) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersAddEmployerManager) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmployerManagersAddEmployerManager) SetPosition(v string) {
	o.Position = v
}

func (o EmployerManagersAddEmployerManager) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployerManagersAddEmployerManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPhone) {
		toSerialize["additional_phone"] = o.AdditionalPhone
	}
	toSerialize["area"] = o.Area
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	toSerialize["is_main_contact_person"] = o.IsMainContactPerson
	toSerialize["last_name"] = o.LastName
	toSerialize["manager_type"] = o.ManagerType
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	toSerialize["phone"] = o.Phone
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

func (o *EmployerManagersAddEmployerManager) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"area",
		"email",
		"first_name",
		"is_main_contact_person",
		"last_name",
		"manager_type",
		"phone",
		"position",
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

	varEmployerManagersAddEmployerManager := _EmployerManagersAddEmployerManager{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployerManagersAddEmployerManager)

	if err != nil {
		return err
	}

	*o = EmployerManagersAddEmployerManager(varEmployerManagersAddEmployerManager)

	return err
}

type NullableEmployerManagersAddEmployerManager struct {
	value *EmployerManagersAddEmployerManager
	isSet bool
}

func (v NullableEmployerManagersAddEmployerManager) Get() *EmployerManagersAddEmployerManager {
	return v.value
}

func (v *NullableEmployerManagersAddEmployerManager) Set(val *EmployerManagersAddEmployerManager) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerManagersAddEmployerManager) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerManagersAddEmployerManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerManagersAddEmployerManager(val *EmployerManagersAddEmployerManager) *NullableEmployerManagersAddEmployerManager {
	return &NullableEmployerManagersAddEmployerManager{value: val, isSet: true}
}

func (v NullableEmployerManagersAddEmployerManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerManagersAddEmployerManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

