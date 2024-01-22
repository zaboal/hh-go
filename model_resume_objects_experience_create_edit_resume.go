/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResumeObjectsExperienceCreateEditResume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeObjectsExperienceCreateEditResume{}

// ResumeObjectsExperienceCreateEditResume struct for ResumeObjectsExperienceCreateEditResume
type ResumeObjectsExperienceCreateEditResume struct {
	Area *IncludesIdNameUrl `json:"area,omitempty"`
	// Название организации
	Company NullableString `json:"company"`
	// Уникальный идентификатор организации
	CompanyId NullableString `json:"company_id,omitempty"`
	// Сайт компании
	CompanyUrl NullableString `json:"company_url,omitempty"`
	// Обязанности, функции, достижения
	Description NullableString `json:"description"`
	Employer *EmployersEmployerInfoShort `json:"employer,omitempty"`
	// Окончание работы (дата в формате `ГГГГ-ММ-ДД`)
	End NullableString `json:"end,omitempty"`
	// Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries)
	Industries []IncludesIdName `json:"industries,omitempty"`
	// Deprecated
	Industry *ResumeObjectsIndustry `json:"industry,omitempty"`
	// Должность
	Position string `json:"position"`
	// Начало работы (дата в формате `ГГГГ-ММ-ДД`)
	Start string `json:"start"`
}

type _ResumeObjectsExperienceCreateEditResume ResumeObjectsExperienceCreateEditResume

// NewResumeObjectsExperienceCreateEditResume instantiates a new ResumeObjectsExperienceCreateEditResume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeObjectsExperienceCreateEditResume(company NullableString, description NullableString, position string, start string) *ResumeObjectsExperienceCreateEditResume {
	this := ResumeObjectsExperienceCreateEditResume{}
	this.Company = company
	this.Description = description
	this.Position = position
	this.Start = start
	return &this
}

// NewResumeObjectsExperienceCreateEditResumeWithDefaults instantiates a new ResumeObjectsExperienceCreateEditResume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeObjectsExperienceCreateEditResumeWithDefaults() *ResumeObjectsExperienceCreateEditResume {
	this := ResumeObjectsExperienceCreateEditResume{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceCreateEditResume) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceCreateEditResume) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given IncludesIdNameUrl and assigns it to the Area field.
func (o *ResumeObjectsExperienceCreateEditResume) SetArea(v IncludesIdNameUrl) {
	o.Area = &v
}

// GetCompany returns the Company field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}

	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// SetCompany sets field value
func (o *ResumeObjectsExperienceCreateEditResume) SetCompany(v string) {
	o.Company.Set(&v)
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}
// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyUrl.Get()
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyUrl.Get(), o.CompanyUrl.IsSet()
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasCompanyUrl() bool {
	if o != nil && o.CompanyUrl.IsSet() {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given NullableString and assigns it to the CompanyUrl field.
func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyUrl(v string) {
	o.CompanyUrl.Set(&v)
}
// SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyUrlNil() {
	o.CompanyUrl.Set(nil)
}

// UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) UnsetCompanyUrl() {
	o.CompanyUrl.Unset()
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ResumeObjectsExperienceCreateEditResume) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceCreateEditResume) GetEmployer() EmployersEmployerInfoShort {
	if o == nil || IsNil(o.Employer) {
		var ret EmployersEmployerInfoShort
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceCreateEditResume) GetEmployerOk() (*EmployersEmployerInfoShort, bool) {
	if o == nil || IsNil(o.Employer) {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasEmployer() bool {
	if o != nil && !IsNil(o.Employer) {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given EmployersEmployerInfoShort and assigns it to the Employer field.
func (o *ResumeObjectsExperienceCreateEditResume) SetEmployer(v EmployersEmployerInfoShort) {
	o.Employer = &v
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceCreateEditResume) GetEnd() string {
	if o == nil || IsNil(o.End.Get()) {
		var ret string
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceCreateEditResume) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableString and assigns it to the End field.
func (o *ResumeObjectsExperienceCreateEditResume) SetEnd(v string) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *ResumeObjectsExperienceCreateEditResume) UnsetEnd() {
	o.End.Unset()
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceCreateEditResume) GetIndustries() []IncludesIdName {
	if o == nil || IsNil(o.Industries) {
		var ret []IncludesIdName
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceCreateEditResume) GetIndustriesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []IncludesIdName and assigns it to the Industries field.
func (o *ResumeObjectsExperienceCreateEditResume) SetIndustries(v []IncludesIdName) {
	o.Industries = v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
// Deprecated
func (o *ResumeObjectsExperienceCreateEditResume) GetIndustry() ResumeObjectsIndustry {
	if o == nil || IsNil(o.Industry) {
		var ret ResumeObjectsIndustry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeObjectsExperienceCreateEditResume) GetIndustryOk() (*ResumeObjectsIndustry, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceCreateEditResume) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given ResumeObjectsIndustry and assigns it to the Industry field.
// Deprecated
func (o *ResumeObjectsExperienceCreateEditResume) SetIndustry(v ResumeObjectsIndustry) {
	o.Industry = &v
}

// GetPosition returns the Position field value
func (o *ResumeObjectsExperienceCreateEditResume) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceCreateEditResume) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *ResumeObjectsExperienceCreateEditResume) SetPosition(v string) {
	o.Position = v
}

// GetStart returns the Start field value
func (o *ResumeObjectsExperienceCreateEditResume) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceCreateEditResume) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResumeObjectsExperienceCreateEditResume) SetStart(v string) {
	o.Start = v
}

func (o ResumeObjectsExperienceCreateEditResume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeObjectsExperienceCreateEditResume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	toSerialize["company"] = o.Company.Get()
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.CompanyUrl.IsSet() {
		toSerialize["company_url"] = o.CompanyUrl.Get()
	}
	toSerialize["description"] = o.Description.Get()
	if !IsNil(o.Employer) {
		toSerialize["employer"] = o.Employer
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	toSerialize["position"] = o.Position
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

func (o *ResumeObjectsExperienceCreateEditResume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
		"description",
		"position",
		"start",
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

	varResumeObjectsExperienceCreateEditResume := _ResumeObjectsExperienceCreateEditResume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeObjectsExperienceCreateEditResume)

	if err != nil {
		return err
	}

	*o = ResumeObjectsExperienceCreateEditResume(varResumeObjectsExperienceCreateEditResume)

	return err
}

type NullableResumeObjectsExperienceCreateEditResume struct {
	value *ResumeObjectsExperienceCreateEditResume
	isSet bool
}

func (v NullableResumeObjectsExperienceCreateEditResume) Get() *ResumeObjectsExperienceCreateEditResume {
	return v.value
}

func (v *NullableResumeObjectsExperienceCreateEditResume) Set(val *ResumeObjectsExperienceCreateEditResume) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeObjectsExperienceCreateEditResume) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeObjectsExperienceCreateEditResume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeObjectsExperienceCreateEditResume(val *ResumeObjectsExperienceCreateEditResume) *NullableResumeObjectsExperienceCreateEditResume {
	return &NullableResumeObjectsExperienceCreateEditResume{value: val, isSet: true}
}

func (v NullableResumeObjectsExperienceCreateEditResume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeObjectsExperienceCreateEditResume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


