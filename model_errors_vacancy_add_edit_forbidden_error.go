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

// checks if the ErrorsVacancyAddEditForbiddenError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsVacancyAddEditForbiddenError{}

// ErrorsVacancyAddEditForbiddenError struct for ErrorsVacancyAddEditForbiddenError
type ErrorsVacancyAddEditForbiddenError struct {
	// Общее количество дубликатов вакансии. Возвращается только для `\"value\": \"duplicate\"` 
	Found NullableFloat32 `json:"found,omitempty"`
	// Ограниченное количество записей с информацией о дубликатах. Не гарантирует выдачу всех дубликатов. Возвращается только для `\"value\": \"duplicate\"` 
	Items []IncludesNumericId `json:"items,omitempty"`
	// Текстовый идентификатор типа ошибки
	Type string `json:"type"`
	// Ошибки при публикации и редактировании вакансии:   * `not_enough_purchased_services` — купленных услуг недостаточно для публикации или обновления данного типа вакансии   * `quota_exceeded` — квота менеджера на публикацию данного типа вакансии закончилась   * `duplicate` — аналогичная вакансия уже опубликована. В ответе передается информация по дубликатам вакансии. Данную ошибку можно форсировано отключить параметром `?ignore_duplicates=true`   * `creation_forbidden` — публикация вакансий недоступна текущему менеджеру   * `unavailable_for_archived` — редактирование недоступно для архивной вакансии   * `conflict_changes` — [конфликтные изменения](https://github.com/hhru/api/blob/master/docs/employer_vacancies.md#%D1%81%D0%BC%D0%B5%D0%BD%D0%B0-%D0%B1%D0%B8%D0%BB%D0%BB%D0%B8%D0%BD%D0%B3-%D1%82%D0%B8%D0%BF%D0%B0-%D0%BC%D0%B5%D0%BD%D0%B5%D0%B4%D0%B6%D0%B5%D1%80%D0%B0-%D0%B2%D0%B0%D0%BA%D0%B0%D0%BD%D1%81%D0%B8%D0%B8) данных вакансии 
	Value string `json:"value"`
}

type _ErrorsVacancyAddEditForbiddenError ErrorsVacancyAddEditForbiddenError

// NewErrorsVacancyAddEditForbiddenError instantiates a new ErrorsVacancyAddEditForbiddenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsVacancyAddEditForbiddenError(type_ string, value string) *ErrorsVacancyAddEditForbiddenError {
	this := ErrorsVacancyAddEditForbiddenError{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewErrorsVacancyAddEditForbiddenErrorWithDefaults instantiates a new ErrorsVacancyAddEditForbiddenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsVacancyAddEditForbiddenErrorWithDefaults() *ErrorsVacancyAddEditForbiddenError {
	this := ErrorsVacancyAddEditForbiddenError{}
	return &this
}

// GetFound returns the Found field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorsVacancyAddEditForbiddenError) GetFound() float32 {
	if o == nil || IsNil(o.Found.Get()) {
		var ret float32
		return ret
	}
	return *o.Found.Get()
}

// GetFoundOk returns a tuple with the Found field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorsVacancyAddEditForbiddenError) GetFoundOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Found.Get(), o.Found.IsSet()
}

// HasFound returns a boolean if a field has been set.
func (o *ErrorsVacancyAddEditForbiddenError) HasFound() bool {
	if o != nil && o.Found.IsSet() {
		return true
	}

	return false
}

// SetFound gets a reference to the given NullableFloat32 and assigns it to the Found field.
func (o *ErrorsVacancyAddEditForbiddenError) SetFound(v float32) {
	o.Found.Set(&v)
}
// SetFoundNil sets the value for Found to be an explicit nil
func (o *ErrorsVacancyAddEditForbiddenError) SetFoundNil() {
	o.Found.Set(nil)
}

// UnsetFound ensures that no value is present for Found, not even an explicit nil
func (o *ErrorsVacancyAddEditForbiddenError) UnsetFound() {
	o.Found.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorsVacancyAddEditForbiddenError) GetItems() []IncludesNumericId {
	if o == nil {
		var ret []IncludesNumericId
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorsVacancyAddEditForbiddenError) GetItemsOk() ([]IncludesNumericId, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ErrorsVacancyAddEditForbiddenError) HasItems() bool {
	if o != nil && IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IncludesNumericId and assigns it to the Items field.
func (o *ErrorsVacancyAddEditForbiddenError) SetItems(v []IncludesNumericId) {
	o.Items = v
}

// GetType returns the Type field value
func (o *ErrorsVacancyAddEditForbiddenError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ErrorsVacancyAddEditForbiddenError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ErrorsVacancyAddEditForbiddenError) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ErrorsVacancyAddEditForbiddenError) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ErrorsVacancyAddEditForbiddenError) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ErrorsVacancyAddEditForbiddenError) SetValue(v string) {
	o.Value = v
}

func (o ErrorsVacancyAddEditForbiddenError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsVacancyAddEditForbiddenError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Found.IsSet() {
		toSerialize["found"] = o.Found.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ErrorsVacancyAddEditForbiddenError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varErrorsVacancyAddEditForbiddenError := _ErrorsVacancyAddEditForbiddenError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorsVacancyAddEditForbiddenError)

	if err != nil {
		return err
	}

	*o = ErrorsVacancyAddEditForbiddenError(varErrorsVacancyAddEditForbiddenError)

	return err
}

type NullableErrorsVacancyAddEditForbiddenError struct {
	value *ErrorsVacancyAddEditForbiddenError
	isSet bool
}

func (v NullableErrorsVacancyAddEditForbiddenError) Get() *ErrorsVacancyAddEditForbiddenError {
	return v.value
}

func (v *NullableErrorsVacancyAddEditForbiddenError) Set(val *ErrorsVacancyAddEditForbiddenError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsVacancyAddEditForbiddenError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsVacancyAddEditForbiddenError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsVacancyAddEditForbiddenError(val *ErrorsVacancyAddEditForbiddenError) *NullableErrorsVacancyAddEditForbiddenError {
	return &NullableErrorsVacancyAddEditForbiddenError{value: val, isSet: true}
}

func (v NullableErrorsVacancyAddEditForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsVacancyAddEditForbiddenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


