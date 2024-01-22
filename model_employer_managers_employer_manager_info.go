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

// checks if the EmployerManagersEmployerManagerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployerManagersEmployerManagerInfo{}

// EmployerManagersEmployerManagerInfo struct for EmployerManagersEmployerManagerInfo
type EmployerManagersEmployerManagerInfo struct {
	AdditionalPhone *EmployerManagersManagerDataAdditionalPhone `json:"additional_phone,omitempty"`
	Area *EmployerManagersArea `json:"area,omitempty"`
	// Адрес электронной почты менеджера
	Email string `json:"email"`
	// Имя менеджера
	FirstName string `json:"first_name"`
	// Полное имя менеджера
	FullName *string `json:"full_name,omitempty"`
	// Идентификатор менеджера
	Id string `json:"id"`
	// Является ли менеджер основным контактным лицом
	IsMainContactPerson bool `json:"is_main_contact_person"`
	// Фамилия менеджера
	LastName string `json:"last_name"`
	ManagerType *EmployerManagersManagerType `json:"manager_type,omitempty"`
	// Отчество менеджера
	MiddleName *string `json:"middle_name,omitempty"`
	// Полное имя менеджера
	// Deprecated
	Name *string `json:"name,omitempty"`
	// Список [прав менеджера](#tag/Menedzhery-rabotodatelya/operation/get-employer-manager-types)
	Permissions []EmployerManagerTypesAvailablePermissions `json:"permissions"`
	Phone EmployerManagersManagerDataPhone `json:"phone"`
	// Должность менеджера
	Position string `json:"position"`
	// Количество опубликованных (активных) вакансий у данного менеджера. `null` — если у пользователя нет прав на просмотр вакансий этого менеджера
	VacanciesCount NullableFloat32 `json:"vacancies_count,omitempty"`
}

type _EmployerManagersEmployerManagerInfo EmployerManagersEmployerManagerInfo

// NewEmployerManagersEmployerManagerInfo instantiates a new EmployerManagersEmployerManagerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerManagersEmployerManagerInfo(email string, firstName string, id string, isMainContactPerson bool, lastName string, permissions []EmployerManagerTypesAvailablePermissions, phone EmployerManagersManagerDataPhone, position string) *EmployerManagersEmployerManagerInfo {
	this := EmployerManagersEmployerManagerInfo{}
	this.Email = email
	this.FirstName = firstName
	this.Id = id
	this.IsMainContactPerson = isMainContactPerson
	this.LastName = lastName
	this.Permissions = permissions
	this.Phone = phone
	this.Position = position
	return &this
}

// NewEmployerManagersEmployerManagerInfoWithDefaults instantiates a new EmployerManagersEmployerManagerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerManagersEmployerManagerInfoWithDefaults() *EmployerManagersEmployerManagerInfo {
	this := EmployerManagersEmployerManagerInfo{}
	return &this
}

// GetAdditionalPhone returns the AdditionalPhone field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerInfo) GetAdditionalPhone() EmployerManagersManagerDataAdditionalPhone {
	if o == nil || IsNil(o.AdditionalPhone) {
		var ret EmployerManagersManagerDataAdditionalPhone
		return ret
	}
	return *o.AdditionalPhone
}

// GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetAdditionalPhoneOk() (*EmployerManagersManagerDataAdditionalPhone, bool) {
	if o == nil || IsNil(o.AdditionalPhone) {
		return nil, false
	}
	return o.AdditionalPhone, true
}

// HasAdditionalPhone returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasAdditionalPhone() bool {
	if o != nil && !IsNil(o.AdditionalPhone) {
		return true
	}

	return false
}

// SetAdditionalPhone gets a reference to the given EmployerManagersManagerDataAdditionalPhone and assigns it to the AdditionalPhone field.
func (o *EmployerManagersEmployerManagerInfo) SetAdditionalPhone(v EmployerManagersManagerDataAdditionalPhone) {
	o.AdditionalPhone = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerInfo) GetArea() EmployerManagersArea {
	if o == nil || IsNil(o.Area) {
		var ret EmployerManagersArea
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetAreaOk() (*EmployerManagersArea, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given EmployerManagersArea and assigns it to the Area field.
func (o *EmployerManagersEmployerManagerInfo) SetArea(v EmployerManagersArea) {
	o.Area = &v
}

// GetEmail returns the Email field value
func (o *EmployerManagersEmployerManagerInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmployerManagersEmployerManagerInfo) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *EmployerManagersEmployerManagerInfo) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EmployerManagersEmployerManagerInfo) SetFirstName(v string) {
	o.FirstName = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerInfo) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *EmployerManagersEmployerManagerInfo) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value
func (o *EmployerManagersEmployerManagerInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmployerManagersEmployerManagerInfo) SetId(v string) {
	o.Id = v
}

// GetIsMainContactPerson returns the IsMainContactPerson field value
func (o *EmployerManagersEmployerManagerInfo) GetIsMainContactPerson() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMainContactPerson
}

// GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetIsMainContactPersonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMainContactPerson, true
}

// SetIsMainContactPerson sets field value
func (o *EmployerManagersEmployerManagerInfo) SetIsMainContactPerson(v bool) {
	o.IsMainContactPerson = v
}

// GetLastName returns the LastName field value
func (o *EmployerManagersEmployerManagerInfo) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *EmployerManagersEmployerManagerInfo) SetLastName(v string) {
	o.LastName = v
}

// GetManagerType returns the ManagerType field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerInfo) GetManagerType() EmployerManagersManagerType {
	if o == nil || IsNil(o.ManagerType) {
		var ret EmployerManagersManagerType
		return ret
	}
	return *o.ManagerType
}

// GetManagerTypeOk returns a tuple with the ManagerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetManagerTypeOk() (*EmployerManagersManagerType, bool) {
	if o == nil || IsNil(o.ManagerType) {
		return nil, false
	}
	return o.ManagerType, true
}

// HasManagerType returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasManagerType() bool {
	if o != nil && !IsNil(o.ManagerType) {
		return true
	}

	return false
}

// SetManagerType gets a reference to the given EmployerManagersManagerType and assigns it to the ManagerType field.
func (o *EmployerManagersEmployerManagerInfo) SetManagerType(v EmployerManagersManagerType) {
	o.ManagerType = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerInfo) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *EmployerManagersEmployerManagerInfo) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
// Deprecated
func (o *EmployerManagersEmployerManagerInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EmployerManagersEmployerManagerInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
// Deprecated
func (o *EmployerManagersEmployerManagerInfo) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value
func (o *EmployerManagersEmployerManagerInfo) GetPermissions() []EmployerManagerTypesAvailablePermissions {
	if o == nil {
		var ret []EmployerManagerTypesAvailablePermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetPermissionsOk() ([]EmployerManagerTypesAvailablePermissions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *EmployerManagersEmployerManagerInfo) SetPermissions(v []EmployerManagerTypesAvailablePermissions) {
	o.Permissions = v
}

// GetPhone returns the Phone field value
func (o *EmployerManagersEmployerManagerInfo) GetPhone() EmployerManagersManagerDataPhone {
	if o == nil {
		var ret EmployerManagersManagerDataPhone
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetPhoneOk() (*EmployerManagersManagerDataPhone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *EmployerManagersEmployerManagerInfo) SetPhone(v EmployerManagersManagerDataPhone) {
	o.Phone = v
}

// GetPosition returns the Position field value
func (o *EmployerManagersEmployerManagerInfo) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerInfo) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmployerManagersEmployerManagerInfo) SetPosition(v string) {
	o.Position = v
}

// GetVacanciesCount returns the VacanciesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerManagersEmployerManagerInfo) GetVacanciesCount() float32 {
	if o == nil || IsNil(o.VacanciesCount.Get()) {
		var ret float32
		return ret
	}
	return *o.VacanciesCount.Get()
}

// GetVacanciesCountOk returns a tuple with the VacanciesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerManagersEmployerManagerInfo) GetVacanciesCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacanciesCount.Get(), o.VacanciesCount.IsSet()
}

// HasVacanciesCount returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerInfo) HasVacanciesCount() bool {
	if o != nil && o.VacanciesCount.IsSet() {
		return true
	}

	return false
}

// SetVacanciesCount gets a reference to the given NullableFloat32 and assigns it to the VacanciesCount field.
func (o *EmployerManagersEmployerManagerInfo) SetVacanciesCount(v float32) {
	o.VacanciesCount.Set(&v)
}
// SetVacanciesCountNil sets the value for VacanciesCount to be an explicit nil
func (o *EmployerManagersEmployerManagerInfo) SetVacanciesCountNil() {
	o.VacanciesCount.Set(nil)
}

// UnsetVacanciesCount ensures that no value is present for VacanciesCount, not even an explicit nil
func (o *EmployerManagersEmployerManagerInfo) UnsetVacanciesCount() {
	o.VacanciesCount.Unset()
}

func (o EmployerManagersEmployerManagerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployerManagersEmployerManagerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPhone) {
		toSerialize["additional_phone"] = o.AdditionalPhone
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	toSerialize["id"] = o.Id
	toSerialize["is_main_contact_person"] = o.IsMainContactPerson
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.ManagerType) {
		toSerialize["manager_type"] = o.ManagerType
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["permissions"] = o.Permissions
	toSerialize["phone"] = o.Phone
	toSerialize["position"] = o.Position
	if o.VacanciesCount.IsSet() {
		toSerialize["vacancies_count"] = o.VacanciesCount.Get()
	}
	return toSerialize, nil
}

func (o *EmployerManagersEmployerManagerInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"first_name",
		"id",
		"is_main_contact_person",
		"last_name",
		"permissions",
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

	varEmployerManagersEmployerManagerInfo := _EmployerManagersEmployerManagerInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployerManagersEmployerManagerInfo)

	if err != nil {
		return err
	}

	*o = EmployerManagersEmployerManagerInfo(varEmployerManagersEmployerManagerInfo)

	return err
}

type NullableEmployerManagersEmployerManagerInfo struct {
	value *EmployerManagersEmployerManagerInfo
	isSet bool
}

func (v NullableEmployerManagersEmployerManagerInfo) Get() *EmployerManagersEmployerManagerInfo {
	return v.value
}

func (v *NullableEmployerManagersEmployerManagerInfo) Set(val *EmployerManagersEmployerManagerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerManagersEmployerManagerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerManagersEmployerManagerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerManagersEmployerManagerInfo(val *EmployerManagersEmployerManagerInfo) *NullableEmployerManagersEmployerManagerInfo {
	return &NullableEmployerManagersEmployerManagerInfo{value: val, isSet: true}
}

func (v NullableEmployerManagersEmployerManagerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerManagersEmployerManagerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


