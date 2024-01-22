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

// checks if the VacanciesVacancyEmployer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyEmployer{}

// VacanciesVacancyEmployer struct for VacanciesVacancyEmployer
type VacanciesVacancyEmployer struct {
	// Флаг, показывающий, прошла ли компания IT аккредитацию
	AccreditedItEmployer *bool `json:"accredited_it_employer,omitempty"`
	// Ссылка на представление компании на сайте
	AlternateUrl NullableString `json:"alternate_url,omitempty"`
	// Идентификатор компании
	Id NullableString `json:"id,omitempty"`
	LogoUrls *IncludesLogoUrls `json:"logo_urls,omitempty"`
	// Название компании
	Name string `json:"name"`
	// Флаг, показывающий, прошла ли компания проверку на сайте
	Trusted bool `json:"trusted"`
	// URL, на который нужно сделать GET-запрос, чтобы получить информацию о компании
	Url NullableString `json:"url,omitempty"`
	// Ссылка на поисковую выдачу вакансий данной компании
	VacanciesUrl NullableString `json:"vacancies_url,omitempty"`
	// Добавлены ли все вакансии работодателя в [список скрытых](#tag/Skrytye-rabotodateli)
	Blacklisted *bool `json:"blacklisted,omitempty"`
	ApplicantServices *IncludesEmployerApplicantServices `json:"applicant_services,omitempty"`
}

type _VacanciesVacancyEmployer VacanciesVacancyEmployer

// NewVacanciesVacancyEmployer instantiates a new VacanciesVacancyEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyEmployer(name string, trusted bool) *VacanciesVacancyEmployer {
	this := VacanciesVacancyEmployer{}
	this.Name = name
	this.Trusted = trusted
	return &this
}

// NewVacanciesVacancyEmployerWithDefaults instantiates a new VacanciesVacancyEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyEmployerWithDefaults() *VacanciesVacancyEmployer {
	this := VacanciesVacancyEmployer{}
	return &this
}

// GetAccreditedItEmployer returns the AccreditedItEmployer field value if set, zero value otherwise.
func (o *VacanciesVacancyEmployer) GetAccreditedItEmployer() bool {
	if o == nil || IsNil(o.AccreditedItEmployer) {
		var ret bool
		return ret
	}
	return *o.AccreditedItEmployer
}

// GetAccreditedItEmployerOk returns a tuple with the AccreditedItEmployer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetAccreditedItEmployerOk() (*bool, bool) {
	if o == nil || IsNil(o.AccreditedItEmployer) {
		return nil, false
	}
	return o.AccreditedItEmployer, true
}

// HasAccreditedItEmployer returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasAccreditedItEmployer() bool {
	if o != nil && !IsNil(o.AccreditedItEmployer) {
		return true
	}

	return false
}

// SetAccreditedItEmployer gets a reference to the given bool and assigns it to the AccreditedItEmployer field.
func (o *VacanciesVacancyEmployer) SetAccreditedItEmployer(v bool) {
	o.AccreditedItEmployer = &v
}

// GetAlternateUrl returns the AlternateUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyEmployer) GetAlternateUrl() string {
	if o == nil || IsNil(o.AlternateUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AlternateUrl.Get()
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyEmployer) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlternateUrl.Get(), o.AlternateUrl.IsSet()
}

// HasAlternateUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasAlternateUrl() bool {
	if o != nil && o.AlternateUrl.IsSet() {
		return true
	}

	return false
}

// SetAlternateUrl gets a reference to the given NullableString and assigns it to the AlternateUrl field.
func (o *VacanciesVacancyEmployer) SetAlternateUrl(v string) {
	o.AlternateUrl.Set(&v)
}
// SetAlternateUrlNil sets the value for AlternateUrl to be an explicit nil
func (o *VacanciesVacancyEmployer) SetAlternateUrlNil() {
	o.AlternateUrl.Set(nil)
}

// UnsetAlternateUrl ensures that no value is present for AlternateUrl, not even an explicit nil
func (o *VacanciesVacancyEmployer) UnsetAlternateUrl() {
	o.AlternateUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyEmployer) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyEmployer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *VacanciesVacancyEmployer) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *VacanciesVacancyEmployer) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *VacanciesVacancyEmployer) UnsetId() {
	o.Id.Unset()
}

// GetLogoUrls returns the LogoUrls field value if set, zero value otherwise.
func (o *VacanciesVacancyEmployer) GetLogoUrls() IncludesLogoUrls {
	if o == nil || IsNil(o.LogoUrls) {
		var ret IncludesLogoUrls
		return ret
	}
	return *o.LogoUrls
}

// GetLogoUrlsOk returns a tuple with the LogoUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetLogoUrlsOk() (*IncludesLogoUrls, bool) {
	if o == nil || IsNil(o.LogoUrls) {
		return nil, false
	}
	return o.LogoUrls, true
}

// HasLogoUrls returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasLogoUrls() bool {
	if o != nil && !IsNil(o.LogoUrls) {
		return true
	}

	return false
}

// SetLogoUrls gets a reference to the given IncludesLogoUrls and assigns it to the LogoUrls field.
func (o *VacanciesVacancyEmployer) SetLogoUrls(v IncludesLogoUrls) {
	o.LogoUrls = &v
}

// GetName returns the Name field value
func (o *VacanciesVacancyEmployer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesVacancyEmployer) SetName(v string) {
	o.Name = v
}

// GetTrusted returns the Trusted field value
func (o *VacanciesVacancyEmployer) GetTrusted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trusted
}

// GetTrustedOk returns a tuple with the Trusted field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trusted, true
}

// SetTrusted sets field value
func (o *VacanciesVacancyEmployer) SetTrusted(v bool) {
	o.Trusted = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyEmployer) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyEmployer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *VacanciesVacancyEmployer) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *VacanciesVacancyEmployer) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *VacanciesVacancyEmployer) UnsetUrl() {
	o.Url.Unset()
}

// GetVacanciesUrl returns the VacanciesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyEmployer) GetVacanciesUrl() string {
	if o == nil || IsNil(o.VacanciesUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VacanciesUrl.Get()
}

// GetVacanciesUrlOk returns a tuple with the VacanciesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyEmployer) GetVacanciesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacanciesUrl.Get(), o.VacanciesUrl.IsSet()
}

// HasVacanciesUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasVacanciesUrl() bool {
	if o != nil && o.VacanciesUrl.IsSet() {
		return true
	}

	return false
}

// SetVacanciesUrl gets a reference to the given NullableString and assigns it to the VacanciesUrl field.
func (o *VacanciesVacancyEmployer) SetVacanciesUrl(v string) {
	o.VacanciesUrl.Set(&v)
}
// SetVacanciesUrlNil sets the value for VacanciesUrl to be an explicit nil
func (o *VacanciesVacancyEmployer) SetVacanciesUrlNil() {
	o.VacanciesUrl.Set(nil)
}

// UnsetVacanciesUrl ensures that no value is present for VacanciesUrl, not even an explicit nil
func (o *VacanciesVacancyEmployer) UnsetVacanciesUrl() {
	o.VacanciesUrl.Unset()
}

// GetBlacklisted returns the Blacklisted field value if set, zero value otherwise.
func (o *VacanciesVacancyEmployer) GetBlacklisted() bool {
	if o == nil || IsNil(o.Blacklisted) {
		var ret bool
		return ret
	}
	return *o.Blacklisted
}

// GetBlacklistedOk returns a tuple with the Blacklisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetBlacklistedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blacklisted) {
		return nil, false
	}
	return o.Blacklisted, true
}

// HasBlacklisted returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasBlacklisted() bool {
	if o != nil && !IsNil(o.Blacklisted) {
		return true
	}

	return false
}

// SetBlacklisted gets a reference to the given bool and assigns it to the Blacklisted field.
func (o *VacanciesVacancyEmployer) SetBlacklisted(v bool) {
	o.Blacklisted = &v
}

// GetApplicantServices returns the ApplicantServices field value if set, zero value otherwise.
func (o *VacanciesVacancyEmployer) GetApplicantServices() IncludesEmployerApplicantServices {
	if o == nil || IsNil(o.ApplicantServices) {
		var ret IncludesEmployerApplicantServices
		return ret
	}
	return *o.ApplicantServices
}

// GetApplicantServicesOk returns a tuple with the ApplicantServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyEmployer) GetApplicantServicesOk() (*IncludesEmployerApplicantServices, bool) {
	if o == nil || IsNil(o.ApplicantServices) {
		return nil, false
	}
	return o.ApplicantServices, true
}

// HasApplicantServices returns a boolean if a field has been set.
func (o *VacanciesVacancyEmployer) HasApplicantServices() bool {
	if o != nil && !IsNil(o.ApplicantServices) {
		return true
	}

	return false
}

// SetApplicantServices gets a reference to the given IncludesEmployerApplicantServices and assigns it to the ApplicantServices field.
func (o *VacanciesVacancyEmployer) SetApplicantServices(v IncludesEmployerApplicantServices) {
	o.ApplicantServices = &v
}

func (o VacanciesVacancyEmployer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyEmployer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccreditedItEmployer) {
		toSerialize["accredited_it_employer"] = o.AccreditedItEmployer
	}
	if o.AlternateUrl.IsSet() {
		toSerialize["alternate_url"] = o.AlternateUrl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.LogoUrls) {
		toSerialize["logo_urls"] = o.LogoUrls
	}
	toSerialize["name"] = o.Name
	toSerialize["trusted"] = o.Trusted
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.VacanciesUrl.IsSet() {
		toSerialize["vacancies_url"] = o.VacanciesUrl.Get()
	}
	if !IsNil(o.Blacklisted) {
		toSerialize["blacklisted"] = o.Blacklisted
	}
	if !IsNil(o.ApplicantServices) {
		toSerialize["applicant_services"] = o.ApplicantServices
	}
	return toSerialize, nil
}

func (o *VacanciesVacancyEmployer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"trusted",
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

	varVacanciesVacancyEmployer := _VacanciesVacancyEmployer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacancyEmployer)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyEmployer(varVacanciesVacancyEmployer)

	return err
}

type NullableVacanciesVacancyEmployer struct {
	value *VacanciesVacancyEmployer
	isSet bool
}

func (v NullableVacanciesVacancyEmployer) Get() *VacanciesVacancyEmployer {
	return v.value
}

func (v *NullableVacanciesVacancyEmployer) Set(val *VacanciesVacancyEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyEmployer(val *VacanciesVacancyEmployer) *NullableVacanciesVacancyEmployer {
	return &NullableVacanciesVacancyEmployer{value: val, isSet: true}
}

func (v NullableVacanciesVacancyEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


