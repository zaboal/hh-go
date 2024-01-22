/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookPayloadNewResponseOrInvitationVacancy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookPayloadNewResponseOrInvitationVacancy{}

// WebhookPayloadNewResponseOrInvitationVacancy Новый отклик или приглашение на вакансии (в отличие от NEW_NEGOTIATION_VACANCY, данное событие будет вызываться как на отклик со стороны соискателя, так и на приглашение со стороны работодателя)
type WebhookPayloadNewResponseOrInvitationVacancy struct {
	// Идентификатор работодателя
	EmployerId string `json:"employer_id"`
	// Дата отклика или приглашения в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: `YYYY-MM-DDThh:mm:ss±hhmm`
	ResponseDate string `json:"response_date"`
	// Идентификатор резюме
	ResumeId NullableString `json:"resume_id"`
	// Идентификатор топика
	TopicId string `json:"topic_id"`
	// Идентификатор вакансии
	VacancyId NullableString `json:"vacancy_id"`
}

type _WebhookPayloadNewResponseOrInvitationVacancy WebhookPayloadNewResponseOrInvitationVacancy

// NewWebhookPayloadNewResponseOrInvitationVacancy instantiates a new WebhookPayloadNewResponseOrInvitationVacancy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPayloadNewResponseOrInvitationVacancy(employerId string, responseDate string, resumeId NullableString, topicId string, vacancyId NullableString) *WebhookPayloadNewResponseOrInvitationVacancy {
	this := WebhookPayloadNewResponseOrInvitationVacancy{}
	this.EmployerId = employerId
	this.ResponseDate = responseDate
	this.ResumeId = resumeId
	this.TopicId = topicId
	this.VacancyId = vacancyId
	return &this
}

// NewWebhookPayloadNewResponseOrInvitationVacancyWithDefaults instantiates a new WebhookPayloadNewResponseOrInvitationVacancy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPayloadNewResponseOrInvitationVacancyWithDefaults() *WebhookPayloadNewResponseOrInvitationVacancy {
	this := WebhookPayloadNewResponseOrInvitationVacancy{}
	return &this
}

// GetEmployerId returns the EmployerId field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetEmployerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployerId
}

// GetEmployerIdOk returns a tuple with the EmployerId field value
// and a boolean to check if the value has been set.
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetEmployerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmployerId, true
}

// SetEmployerId sets field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetEmployerId(v string) {
	o.EmployerId = v
}

// GetResponseDate returns the ResponseDate field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResponseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseDate
}

// GetResponseDateOk returns a tuple with the ResponseDate field value
// and a boolean to check if the value has been set.
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResponseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseDate, true
}

// SetResponseDate sets field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetResponseDate(v string) {
	o.ResponseDate = v
}

// GetResumeId returns the ResumeId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResumeId() string {
	if o == nil || o.ResumeId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResumeId.Get()
}

// GetResumeIdOk returns a tuple with the ResumeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeId.Get(), o.ResumeId.IsSet()
}

// SetResumeId sets field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetResumeId(v string) {
	o.ResumeId.Set(&v)
}

// GetTopicId returns the TopicId field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetTopicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value
// and a boolean to check if the value has been set.
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetTopicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicId, true
}

// SetTopicId sets field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetTopicId(v string) {
	o.TopicId = v
}

// GetVacancyId returns the VacancyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetVacancyId() string {
	if o == nil || o.VacancyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.VacancyId.Get()
}

// GetVacancyIdOk returns a tuple with the VacancyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetVacancyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyId.Get(), o.VacancyId.IsSet()
}

// SetVacancyId sets field value
func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetVacancyId(v string) {
	o.VacancyId.Set(&v)
}

func (o WebhookPayloadNewResponseOrInvitationVacancy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookPayloadNewResponseOrInvitationVacancy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["employer_id"] = o.EmployerId
	toSerialize["response_date"] = o.ResponseDate
	toSerialize["resume_id"] = o.ResumeId.Get()
	toSerialize["topic_id"] = o.TopicId
	toSerialize["vacancy_id"] = o.VacancyId.Get()
	return toSerialize, nil
}

func (o *WebhookPayloadNewResponseOrInvitationVacancy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"employer_id",
		"response_date",
		"resume_id",
		"topic_id",
		"vacancy_id",
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

	varWebhookPayloadNewResponseOrInvitationVacancy := _WebhookPayloadNewResponseOrInvitationVacancy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookPayloadNewResponseOrInvitationVacancy)

	if err != nil {
		return err
	}

	*o = WebhookPayloadNewResponseOrInvitationVacancy(varWebhookPayloadNewResponseOrInvitationVacancy)

	return err
}

type NullableWebhookPayloadNewResponseOrInvitationVacancy struct {
	value *WebhookPayloadNewResponseOrInvitationVacancy
	isSet bool
}

func (v NullableWebhookPayloadNewResponseOrInvitationVacancy) Get() *WebhookPayloadNewResponseOrInvitationVacancy {
	return v.value
}

func (v *NullableWebhookPayloadNewResponseOrInvitationVacancy) Set(val *WebhookPayloadNewResponseOrInvitationVacancy) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookPayloadNewResponseOrInvitationVacancy) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookPayloadNewResponseOrInvitationVacancy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookPayloadNewResponseOrInvitationVacancy(val *WebhookPayloadNewResponseOrInvitationVacancy) *NullableWebhookPayloadNewResponseOrInvitationVacancy {
	return &NullableWebhookPayloadNewResponseOrInvitationVacancy{value: val, isSet: true}
}

func (v NullableWebhookPayloadNewResponseOrInvitationVacancy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookPayloadNewResponseOrInvitationVacancy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


