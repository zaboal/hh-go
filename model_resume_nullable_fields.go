/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Общая информация  * Всё API работает по протоколу HTTPS. * Авторизация осуществляется по протоколу OAuth2. * Все данные доступны только в формате JSON. * Базовый URL — `https://api.hh.ru/` * Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta) * <a name=\"date-format\"></a> Даты форматируются в соответствии с [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.   <a name=\"request-requirements\"></a> ## Требования к запросам  В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни один заголовок, то ответом будет `400 Bad Request`. Указание в заголовке названия приложения и контактной почты разработчика позволит нам оперативно с вами связаться в случае необходимости. Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка, обработан будет только `HH-User-Agent`.  ``` User-Agent: MyApp/1.0 (my-app-feedback@example.com) ```  Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).   <a name=\"request-body\"></a> ## Формат тела запроса при отправке JSON  Данные, передающиеся в теле запроса, должны удовлетворять требованиям:  * Валидный JSON (допускается передача как минифицированного варианта, так и pretty print варианта с дополнительными пробелами и сбросами строк). * Рекомендуется использование кодировки UTF-8 без дополнительного экранирования (`{\"name\": \"Иванов Иван\"}`). * Также возможно использовать ascii кодировку с экранированием (`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`). * К типам данных в определённым полях накладываются дополнительные условия, описанные в каждом конкретном методе. В JSON типами данных являются `string`, `number`, `boolean`, `null`, `object`, `array`.  ### Ответ Ответ свыше определенной длины будет сжиматься методом gzip.  ### Ошибки и коды ответов  API широко использует информирование при помощи кодов ответов. Приложение должно корректно их обрабатывать.  В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.  При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана дополнительная информация, позволяющая разработчику понять причину соответствующего ответа.  [Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).   ## Недокументированные поля и параметры запросов  В ответах и параметрах API можно найти ключи, не описанные в документации. Обычно это означает, что они оставлены для совместимости со старыми версиями. Их использование не рекомендуется. Если ваше приложение использует такие ключи, перейдите на использование актуальных ключей, описанных в документации.   ## Пагинация  К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где доступна пагинация, единообразный корневой объект:  ```json {   \"found\": 1,   \"per_page\": 1,   \"pages\": 1,   \"page\": 0,   \"items\": [{}] } ``` ## Выбор сайта  API HeadHunter позволяет получать данные со всех сайтов группы компании HeadHunter.  В частности:  * hh.ru * rabota.by * hh1.az * hh.uz * hh.kz * headhunter.ge * headhunter.kg  Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.  При необходимости учесть специфику сайта, можно добавить в запрос параметр `?host=`. По умолчанию используется `hh.ru`.  Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.  ## CORS (Cross-Origin Resource Sharing)  API поддерживает технологию CORS для запроса данных из браузера с произвольного домена. Этот метод более предпочтителен, чем использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен [специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр `?callback=callback_name`.  * [CORS specification on w3.org](http://www.w3.org/TR/cors/) * [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/) * [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/) * [CORS on caniuse.com](http://caniuse.com/#feat=cors) * [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)   ## Внешние ссылки на статьи и стандарты  * [HTTP/1.1](http://tools.ietf.org/html/rfc2616) * [JSON](http://json.org/) * [URI Template](http://tools.ietf.org/html/rfc6570) * [OAuth 2.0](http://tools.ietf.org/html/rfc6749) * [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) * [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#tag/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#tag/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпрометированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh

import (
	"encoding/json"
)

// checks if the ResumeNullableFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeNullableFields{}

// ResumeNullableFields Необязательные поля при создании резюме
type ResumeNullableFields struct {
	Access NullableResumeObjectsAccess `json:"access,omitempty"`
	// День рождения (в формате `ГГГГ-ММ-ДД`)
	BirthDate NullableString `json:"birth_date,omitempty"`
	BusinessTripReadiness NullableIncludesId `json:"business_trip_readiness,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate,omitempty"`
	// Список категорий водительских прав соискателя
	DriverLicenseTypes []ResumeObjectsDriverLicenseTypes `json:"driver_license_types,omitempty"`
	// Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Employments []IncludesIdName `json:"employments,omitempty"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	// Наличие личного автомобиля у соискателя
	HasVehicle NullableBool `json:"has_vehicle,omitempty"`
	// Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields,omitempty"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	Metro NullableIncludesId `json:"metro,omitempty"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Photo NullableResumeObjectsPhoto `json:"photo,omitempty"`
	// Список изображений в портфолио пользователя
	Portfolio []ResumeObjectsPortfolio `json:"portfolio,omitempty"`
	// Массив объектов профролей. Элемент справочника [professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)
	ProfessionalRoles []IncludesId `json:"professional_roles,omitempty"`
	// Список рекомендаций
	Recommendation []ResumeObjectsRecommendation `json:"recommendation,omitempty"`
	Relocation NullableResumeObjectsRelocationPublic `json:"relocation,omitempty"`
	ResumeLocale NullableIncludesIdName `json:"resume_locale,omitempty"`
	Salary NullableResumeObjectsSalaryAddEdit `json:"salary,omitempty"`
	// Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Schedules []IncludesIdName `json:"schedules,omitempty"`
	// Профили в соц. сетях и других сервисах
	Site []ResumeObjectsSite `json:"site,omitempty"`
	// Ключевые навыки (список уникальных строк)
	SkillSet []string `json:"skill_set,omitempty"`
	// Дополнительная информация, описание навыков в свободной форме
	Skills NullableString `json:"skills,omitempty"`
	// Желаемая должность
	Title NullableString `json:"title,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	TravelTime NullableIncludesId `json:"travel_time,omitempty"`
	// Список регионов, в который соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) 
	WorkTicket []IncludesId `json:"work_ticket,omitempty"`
}

// NewResumeNullableFields instantiates a new ResumeNullableFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeNullableFields() *ResumeNullableFields {
	this := ResumeNullableFields{}
	return &this
}

// NewResumeNullableFieldsWithDefaults instantiates a new ResumeNullableFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeNullableFieldsWithDefaults() *ResumeNullableFields {
	this := ResumeNullableFields{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetAccess() ResumeObjectsAccess {
	if o == nil || IsNil(o.Access.Get()) {
		var ret ResumeObjectsAccess
		return ret
	}
	return *o.Access.Get()
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetAccessOk() (*ResumeObjectsAccess, bool) {
	if o == nil {
		return nil, false
	}
	return o.Access.Get(), o.Access.IsSet()
}

// HasAccess returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasAccess() bool {
	if o != nil && o.Access.IsSet() {
		return true
	}

	return false
}

// SetAccess gets a reference to the given NullableResumeObjectsAccess and assigns it to the Access field.
func (o *ResumeNullableFields) SetAccess(v ResumeObjectsAccess) {
	o.Access.Set(&v)
}
// SetAccessNil sets the value for Access to be an explicit nil
func (o *ResumeNullableFields) SetAccessNil() {
	o.Access.Set(nil)
}

// UnsetAccess ensures that no value is present for Access, not even an explicit nil
func (o *ResumeNullableFields) UnsetAccess() {
	o.Access.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *ResumeNullableFields) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *ResumeNullableFields) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *ResumeNullableFields) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetBusinessTripReadiness() IncludesId {
	if o == nil || IsNil(o.BusinessTripReadiness.Get()) {
		var ret IncludesId
		return ret
	}
	return *o.BusinessTripReadiness.Get()
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetBusinessTripReadinessOk() (*IncludesId, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessTripReadiness.Get(), o.BusinessTripReadiness.IsSet()
}

// HasBusinessTripReadiness returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasBusinessTripReadiness() bool {
	if o != nil && o.BusinessTripReadiness.IsSet() {
		return true
	}

	return false
}

// SetBusinessTripReadiness gets a reference to the given NullableIncludesId and assigns it to the BusinessTripReadiness field.
func (o *ResumeNullableFields) SetBusinessTripReadiness(v IncludesId) {
	o.BusinessTripReadiness.Set(&v)
}
// SetBusinessTripReadinessNil sets the value for BusinessTripReadiness to be an explicit nil
func (o *ResumeNullableFields) SetBusinessTripReadinessNil() {
	o.BusinessTripReadiness.Set(nil)
}

// UnsetBusinessTripReadiness ensures that no value is present for BusinessTripReadiness, not even an explicit nil
func (o *ResumeNullableFields) UnsetBusinessTripReadiness() {
	o.BusinessTripReadiness.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}
	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given []ResumeObjectsCertificate and assigns it to the Certificate field.
func (o *ResumeNullableFields) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes {
	if o == nil {
		var ret []ResumeObjectsDriverLicenseTypes
		return ret
	}
	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetDriverLicenseTypesOk() ([]ResumeObjectsDriverLicenseTypes, bool) {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasDriverLicenseTypes() bool {
	if o != nil && !IsNil(o.DriverLicenseTypes) {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given []ResumeObjectsDriverLicenseTypes and assigns it to the DriverLicenseTypes field.
func (o *ResumeNullableFields) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes) {
	o.DriverLicenseTypes = v
}

// GetEmployments returns the Employments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetEmployments() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetEmploymentsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Employments) {
		return nil, false
	}
	return o.Employments, true
}

// HasEmployments returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasEmployments() bool {
	if o != nil && !IsNil(o.Employments) {
		return true
	}

	return false
}

// SetEmployments gets a reference to the given []IncludesIdName and assigns it to the Employments field.
func (o *ResumeNullableFields) SetEmployments(v []IncludesIdName) {
	o.Employments = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ResumeNullableFields) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ResumeNullableFields) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ResumeNullableFields) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetHasVehicle returns the HasVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetHasVehicle() bool {
	if o == nil || IsNil(o.HasVehicle.Get()) {
		var ret bool
		return ret
	}
	return *o.HasVehicle.Get()
}

// GetHasVehicleOk returns a tuple with the HasVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetHasVehicleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasVehicle.Get(), o.HasVehicle.IsSet()
}

// HasHasVehicle returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasHasVehicle() bool {
	if o != nil && o.HasVehicle.IsSet() {
		return true
	}

	return false
}

// SetHasVehicle gets a reference to the given NullableBool and assigns it to the HasVehicle field.
func (o *ResumeNullableFields) SetHasVehicle(v bool) {
	o.HasVehicle.Set(&v)
}
// SetHasVehicleNil sets the value for HasVehicle to be an explicit nil
func (o *ResumeNullableFields) SetHasVehicleNil() {
	o.HasVehicle.Set(nil)
}

// UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
func (o *ResumeNullableFields) UnsetHasVehicle() {
	o.HasVehicle.Unset()
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasHiddenFields() bool {
	if o != nil && !IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []IncludesIdName and assigns it to the HiddenFields field.
func (o *ResumeNullableFields) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ResumeNullableFields) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ResumeNullableFields) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ResumeNullableFields) UnsetLastName() {
	o.LastName.Unset()
}

// GetMetro returns the Metro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetMetro() IncludesId {
	if o == nil || IsNil(o.Metro.Get()) {
		var ret IncludesId
		return ret
	}
	return *o.Metro.Get()
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetMetroOk() (*IncludesId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metro.Get(), o.Metro.IsSet()
}

// HasMetro returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasMetro() bool {
	if o != nil && o.Metro.IsSet() {
		return true
	}

	return false
}

// SetMetro gets a reference to the given NullableIncludesId and assigns it to the Metro field.
func (o *ResumeNullableFields) SetMetro(v IncludesId) {
	o.Metro.Set(&v)
}
// SetMetroNil sets the value for Metro to be an explicit nil
func (o *ResumeNullableFields) SetMetroNil() {
	o.Metro.Set(nil)
}

// UnsetMetro ensures that no value is present for Metro, not even an explicit nil
func (o *ResumeNullableFields) UnsetMetro() {
	o.Metro.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *ResumeNullableFields) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *ResumeNullableFields) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *ResumeNullableFields) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPhoto returns the Photo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetPhoto() ResumeObjectsPhoto {
	if o == nil || IsNil(o.Photo.Get()) {
		var ret ResumeObjectsPhoto
		return ret
	}
	return *o.Photo.Get()
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetPhotoOk() (*ResumeObjectsPhoto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo.Get(), o.Photo.IsSet()
}

// HasPhoto returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasPhoto() bool {
	if o != nil && o.Photo.IsSet() {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given NullableResumeObjectsPhoto and assigns it to the Photo field.
func (o *ResumeNullableFields) SetPhoto(v ResumeObjectsPhoto) {
	o.Photo.Set(&v)
}
// SetPhotoNil sets the value for Photo to be an explicit nil
func (o *ResumeNullableFields) SetPhotoNil() {
	o.Photo.Set(nil)
}

// UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
func (o *ResumeNullableFields) UnsetPhoto() {
	o.Photo.Unset()
}

// GetPortfolio returns the Portfolio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetPortfolio() []ResumeObjectsPortfolio {
	if o == nil {
		var ret []ResumeObjectsPortfolio
		return ret
	}
	return o.Portfolio
}

// GetPortfolioOk returns a tuple with the Portfolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetPortfolioOk() ([]ResumeObjectsPortfolio, bool) {
	if o == nil || IsNil(o.Portfolio) {
		return nil, false
	}
	return o.Portfolio, true
}

// HasPortfolio returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasPortfolio() bool {
	if o != nil && !IsNil(o.Portfolio) {
		return true
	}

	return false
}

// SetPortfolio gets a reference to the given []ResumeObjectsPortfolio and assigns it to the Portfolio field.
func (o *ResumeNullableFields) SetPortfolio(v []ResumeObjectsPortfolio) {
	o.Portfolio = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise.
func (o *ResumeNullableFields) GetProfessionalRoles() []IncludesId {
	if o == nil || IsNil(o.ProfessionalRoles) {
		var ret []IncludesId
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeNullableFields) GetProfessionalRolesOk() ([]IncludesId, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasProfessionalRoles() bool {
	if o != nil && !IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []IncludesId and assigns it to the ProfessionalRoles field.
func (o *ResumeNullableFields) SetProfessionalRoles(v []IncludesId) {
	o.ProfessionalRoles = v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetRecommendation() []ResumeObjectsRecommendation {
	if o == nil {
		var ret []ResumeObjectsRecommendation
		return ret
	}
	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetRecommendationOk() ([]ResumeObjectsRecommendation, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasRecommendation() bool {
	if o != nil && !IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given []ResumeObjectsRecommendation and assigns it to the Recommendation field.
func (o *ResumeNullableFields) SetRecommendation(v []ResumeObjectsRecommendation) {
	o.Recommendation = v
}

// GetRelocation returns the Relocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetRelocation() ResumeObjectsRelocationPublic {
	if o == nil || IsNil(o.Relocation.Get()) {
		var ret ResumeObjectsRelocationPublic
		return ret
	}
	return *o.Relocation.Get()
}

// GetRelocationOk returns a tuple with the Relocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relocation.Get(), o.Relocation.IsSet()
}

// HasRelocation returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasRelocation() bool {
	if o != nil && o.Relocation.IsSet() {
		return true
	}

	return false
}

// SetRelocation gets a reference to the given NullableResumeObjectsRelocationPublic and assigns it to the Relocation field.
func (o *ResumeNullableFields) SetRelocation(v ResumeObjectsRelocationPublic) {
	o.Relocation.Set(&v)
}
// SetRelocationNil sets the value for Relocation to be an explicit nil
func (o *ResumeNullableFields) SetRelocationNil() {
	o.Relocation.Set(nil)
}

// UnsetRelocation ensures that no value is present for Relocation, not even an explicit nil
func (o *ResumeNullableFields) UnsetRelocation() {
	o.Relocation.Unset()
}

// GetResumeLocale returns the ResumeLocale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetResumeLocale() IncludesIdName {
	if o == nil || IsNil(o.ResumeLocale.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.ResumeLocale.Get()
}

// GetResumeLocaleOk returns a tuple with the ResumeLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetResumeLocaleOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeLocale.Get(), o.ResumeLocale.IsSet()
}

// HasResumeLocale returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasResumeLocale() bool {
	if o != nil && o.ResumeLocale.IsSet() {
		return true
	}

	return false
}

// SetResumeLocale gets a reference to the given NullableIncludesIdName and assigns it to the ResumeLocale field.
func (o *ResumeNullableFields) SetResumeLocale(v IncludesIdName) {
	o.ResumeLocale.Set(&v)
}
// SetResumeLocaleNil sets the value for ResumeLocale to be an explicit nil
func (o *ResumeNullableFields) SetResumeLocaleNil() {
	o.ResumeLocale.Set(nil)
}

// UnsetResumeLocale ensures that no value is present for ResumeLocale, not even an explicit nil
func (o *ResumeNullableFields) UnsetResumeLocale() {
	o.ResumeLocale.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetSalary() ResumeObjectsSalaryAddEdit {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret ResumeObjectsSalaryAddEdit
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetSalaryOk() (*ResumeObjectsSalaryAddEdit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableResumeObjectsSalaryAddEdit and assigns it to the Salary field.
func (o *ResumeNullableFields) SetSalary(v ResumeObjectsSalaryAddEdit) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *ResumeNullableFields) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *ResumeNullableFields) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedules returns the Schedules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetSchedules() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetSchedulesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []IncludesIdName and assigns it to the Schedules field.
func (o *ResumeNullableFields) SetSchedules(v []IncludesIdName) {
	o.Schedules = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetSite() []ResumeObjectsSite {
	if o == nil {
		var ret []ResumeObjectsSite
		return ret
	}
	return o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetSiteOk() ([]ResumeObjectsSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given []ResumeObjectsSite and assigns it to the Site field.
func (o *ResumeNullableFields) SetSite(v []ResumeObjectsSite) {
	o.Site = v
}

// GetSkillSet returns the SkillSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetSkillSet() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SkillSet
}

// GetSkillSetOk returns a tuple with the SkillSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetSkillSetOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillSet) {
		return nil, false
	}
	return o.SkillSet, true
}

// HasSkillSet returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasSkillSet() bool {
	if o != nil && !IsNil(o.SkillSet) {
		return true
	}

	return false
}

// SetSkillSet gets a reference to the given []string and assigns it to the SkillSet field.
func (o *ResumeNullableFields) SetSkillSet(v []string) {
	o.SkillSet = v
}

// GetSkills returns the Skills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetSkills() string {
	if o == nil || IsNil(o.Skills.Get()) {
		var ret string
		return ret
	}
	return *o.Skills.Get()
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetSkillsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills.Get(), o.Skills.IsSet()
}

// HasSkills returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasSkills() bool {
	if o != nil && o.Skills.IsSet() {
		return true
	}

	return false
}

// SetSkills gets a reference to the given NullableString and assigns it to the Skills field.
func (o *ResumeNullableFields) SetSkills(v string) {
	o.Skills.Set(&v)
}
// SetSkillsNil sets the value for Skills to be an explicit nil
func (o *ResumeNullableFields) SetSkillsNil() {
	o.Skills.Set(nil)
}

// UnsetSkills ensures that no value is present for Skills, not even an explicit nil
func (o *ResumeNullableFields) UnsetSkills() {
	o.Skills.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ResumeNullableFields) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ResumeNullableFields) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ResumeNullableFields) UnsetTitle() {
	o.Title.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *ResumeNullableFields) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *ResumeNullableFields) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *ResumeNullableFields) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetTravelTime returns the TravelTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetTravelTime() IncludesId {
	if o == nil || IsNil(o.TravelTime.Get()) {
		var ret IncludesId
		return ret
	}
	return *o.TravelTime.Get()
}

// GetTravelTimeOk returns a tuple with the TravelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetTravelTimeOk() (*IncludesId, bool) {
	if o == nil {
		return nil, false
	}
	return o.TravelTime.Get(), o.TravelTime.IsSet()
}

// HasTravelTime returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasTravelTime() bool {
	if o != nil && o.TravelTime.IsSet() {
		return true
	}

	return false
}

// SetTravelTime gets a reference to the given NullableIncludesId and assigns it to the TravelTime field.
func (o *ResumeNullableFields) SetTravelTime(v IncludesId) {
	o.TravelTime.Set(&v)
}
// SetTravelTimeNil sets the value for TravelTime to be an explicit nil
func (o *ResumeNullableFields) SetTravelTimeNil() {
	o.TravelTime.Set(nil)
}

// UnsetTravelTime ensures that no value is present for TravelTime, not even an explicit nil
func (o *ResumeNullableFields) UnsetTravelTime() {
	o.TravelTime.Unset()
}

// GetWorkTicket returns the WorkTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeNullableFields) GetWorkTicket() []IncludesId {
	if o == nil {
		var ret []IncludesId
		return ret
	}
	return o.WorkTicket
}

// GetWorkTicketOk returns a tuple with the WorkTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeNullableFields) GetWorkTicketOk() ([]IncludesId, bool) {
	if o == nil || IsNil(o.WorkTicket) {
		return nil, false
	}
	return o.WorkTicket, true
}

// HasWorkTicket returns a boolean if a field has been set.
func (o *ResumeNullableFields) HasWorkTicket() bool {
	if o != nil && !IsNil(o.WorkTicket) {
		return true
	}

	return false
}

// SetWorkTicket gets a reference to the given []IncludesId and assigns it to the WorkTicket field.
func (o *ResumeNullableFields) SetWorkTicket(v []IncludesId) {
	o.WorkTicket = v
}

func (o ResumeNullableFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeNullableFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Access.IsSet() {
		toSerialize["access"] = o.Access.Get()
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if o.BusinessTripReadiness.IsSet() {
		toSerialize["business_trip_readiness"] = o.BusinessTripReadiness.Get()
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.DriverLicenseTypes != nil {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes
	}
	if o.Employments != nil {
		toSerialize["employments"] = o.Employments
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.HasVehicle.IsSet() {
		toSerialize["has_vehicle"] = o.HasVehicle.Get()
	}
	if o.HiddenFields != nil {
		toSerialize["hidden_fields"] = o.HiddenFields
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Metro.IsSet() {
		toSerialize["metro"] = o.Metro.Get()
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if o.Photo.IsSet() {
		toSerialize["photo"] = o.Photo.Get()
	}
	if o.Portfolio != nil {
		toSerialize["portfolio"] = o.Portfolio
	}
	if !IsNil(o.ProfessionalRoles) {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	if o.Recommendation != nil {
		toSerialize["recommendation"] = o.Recommendation
	}
	if o.Relocation.IsSet() {
		toSerialize["relocation"] = o.Relocation.Get()
	}
	if o.ResumeLocale.IsSet() {
		toSerialize["resume_locale"] = o.ResumeLocale.Get()
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.SkillSet != nil {
		toSerialize["skill_set"] = o.SkillSet
	}
	if o.Skills.IsSet() {
		toSerialize["skills"] = o.Skills.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	if o.TravelTime.IsSet() {
		toSerialize["travel_time"] = o.TravelTime.Get()
	}
	if o.WorkTicket != nil {
		toSerialize["work_ticket"] = o.WorkTicket
	}
	return toSerialize, nil
}

type NullableResumeNullableFields struct {
	value *ResumeNullableFields
	isSet bool
}

func (v NullableResumeNullableFields) Get() *ResumeNullableFields {
	return v.value
}

func (v *NullableResumeNullableFields) Set(val *ResumeNullableFields) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeNullableFields) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeNullableFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeNullableFields(val *ResumeNullableFields) *NullableResumeNullableFields {
	return &NullableResumeNullableFields{value: val, isSet: true}
}

func (v NullableResumeNullableFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeNullableFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


