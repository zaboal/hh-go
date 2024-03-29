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
	"bytes"
	"fmt"
)

// checks if the SalaryStatisticsResultingParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalaryStatisticsResultingParameters{}

// SalaryStatisticsResultingParameters Набор параметров, по которым происходил расчет
type SalaryStatisticsResultingParameters struct {
	// Коды регионов
	Areas []IncludesIdName `json:"areas"`
	// Уровни специалистов
	EmployeeLevels []IncludesIdName `json:"employee_levels,omitempty"`
	// Количество работодателей, позиции которых участвуют в выборке
	EmployersCount float32 `json:"employers_count"`
	// Исключенные коды регионов
	ExcludedAreas []IncludesIdName `json:"excluded_areas,omitempty"`
	IndirectCalculation NullableSalaryStatisticsIndirectCalculation `json:"indirect_calculation,omitempty"`
	// Отрасли
	Industries []IncludesIdName `json:"industries,omitempty"`
	// Количество позиций, по которым построена выборка
	PositionsCount float32 `json:"positions_count"`
	// Источники данных. Возможные значения:  * `SALARIES` — данные из банка зарплат; * `RESUMES` — данные из резюме; * `VACANCIES` — данные из вакансий 
	Sources []string `json:"sources"`
	// Профессиональные области и специализаций
	Specialities []IncludesIdName `json:"specialities,omitempty"`
}

type _SalaryStatisticsResultingParameters SalaryStatisticsResultingParameters

// NewSalaryStatisticsResultingParameters instantiates a new SalaryStatisticsResultingParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalaryStatisticsResultingParameters(areas []IncludesIdName, employersCount float32, positionsCount float32, sources []string) *SalaryStatisticsResultingParameters {
	this := SalaryStatisticsResultingParameters{}
	this.Areas = areas
	this.EmployersCount = employersCount
	this.PositionsCount = positionsCount
	this.Sources = sources
	return &this
}

// NewSalaryStatisticsResultingParametersWithDefaults instantiates a new SalaryStatisticsResultingParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalaryStatisticsResultingParametersWithDefaults() *SalaryStatisticsResultingParameters {
	this := SalaryStatisticsResultingParameters{}
	return &this
}

// GetAreas returns the Areas field value
func (o *SalaryStatisticsResultingParameters) GetAreas() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value
// and a boolean to check if the value has been set.
func (o *SalaryStatisticsResultingParameters) GetAreasOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Areas, true
}

// SetAreas sets field value
func (o *SalaryStatisticsResultingParameters) SetAreas(v []IncludesIdName) {
	o.Areas = v
}

// GetEmployeeLevels returns the EmployeeLevels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalaryStatisticsResultingParameters) GetEmployeeLevels() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.EmployeeLevels
}

// GetEmployeeLevelsOk returns a tuple with the EmployeeLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalaryStatisticsResultingParameters) GetEmployeeLevelsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.EmployeeLevels) {
		return nil, false
	}
	return o.EmployeeLevels, true
}

// HasEmployeeLevels returns a boolean if a field has been set.
func (o *SalaryStatisticsResultingParameters) HasEmployeeLevels() bool {
	if o != nil && !IsNil(o.EmployeeLevels) {
		return true
	}

	return false
}

// SetEmployeeLevels gets a reference to the given []IncludesIdName and assigns it to the EmployeeLevels field.
func (o *SalaryStatisticsResultingParameters) SetEmployeeLevels(v []IncludesIdName) {
	o.EmployeeLevels = v
}

// GetEmployersCount returns the EmployersCount field value
func (o *SalaryStatisticsResultingParameters) GetEmployersCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EmployersCount
}

// GetEmployersCountOk returns a tuple with the EmployersCount field value
// and a boolean to check if the value has been set.
func (o *SalaryStatisticsResultingParameters) GetEmployersCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmployersCount, true
}

// SetEmployersCount sets field value
func (o *SalaryStatisticsResultingParameters) SetEmployersCount(v float32) {
	o.EmployersCount = v
}

// GetExcludedAreas returns the ExcludedAreas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalaryStatisticsResultingParameters) GetExcludedAreas() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.ExcludedAreas
}

// GetExcludedAreasOk returns a tuple with the ExcludedAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalaryStatisticsResultingParameters) GetExcludedAreasOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ExcludedAreas) {
		return nil, false
	}
	return o.ExcludedAreas, true
}

// HasExcludedAreas returns a boolean if a field has been set.
func (o *SalaryStatisticsResultingParameters) HasExcludedAreas() bool {
	if o != nil && !IsNil(o.ExcludedAreas) {
		return true
	}

	return false
}

// SetExcludedAreas gets a reference to the given []IncludesIdName and assigns it to the ExcludedAreas field.
func (o *SalaryStatisticsResultingParameters) SetExcludedAreas(v []IncludesIdName) {
	o.ExcludedAreas = v
}

// GetIndirectCalculation returns the IndirectCalculation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalaryStatisticsResultingParameters) GetIndirectCalculation() SalaryStatisticsIndirectCalculation {
	if o == nil || IsNil(o.IndirectCalculation.Get()) {
		var ret SalaryStatisticsIndirectCalculation
		return ret
	}
	return *o.IndirectCalculation.Get()
}

// GetIndirectCalculationOk returns a tuple with the IndirectCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalaryStatisticsResultingParameters) GetIndirectCalculationOk() (*SalaryStatisticsIndirectCalculation, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndirectCalculation.Get(), o.IndirectCalculation.IsSet()
}

// HasIndirectCalculation returns a boolean if a field has been set.
func (o *SalaryStatisticsResultingParameters) HasIndirectCalculation() bool {
	if o != nil && o.IndirectCalculation.IsSet() {
		return true
	}

	return false
}

// SetIndirectCalculation gets a reference to the given NullableSalaryStatisticsIndirectCalculation and assigns it to the IndirectCalculation field.
func (o *SalaryStatisticsResultingParameters) SetIndirectCalculation(v SalaryStatisticsIndirectCalculation) {
	o.IndirectCalculation.Set(&v)
}
// SetIndirectCalculationNil sets the value for IndirectCalculation to be an explicit nil
func (o *SalaryStatisticsResultingParameters) SetIndirectCalculationNil() {
	o.IndirectCalculation.Set(nil)
}

// UnsetIndirectCalculation ensures that no value is present for IndirectCalculation, not even an explicit nil
func (o *SalaryStatisticsResultingParameters) UnsetIndirectCalculation() {
	o.IndirectCalculation.Unset()
}

// GetIndustries returns the Industries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalaryStatisticsResultingParameters) GetIndustries() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalaryStatisticsResultingParameters) GetIndustriesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *SalaryStatisticsResultingParameters) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []IncludesIdName and assigns it to the Industries field.
func (o *SalaryStatisticsResultingParameters) SetIndustries(v []IncludesIdName) {
	o.Industries = v
}

// GetPositionsCount returns the PositionsCount field value
func (o *SalaryStatisticsResultingParameters) GetPositionsCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PositionsCount
}

// GetPositionsCountOk returns a tuple with the PositionsCount field value
// and a boolean to check if the value has been set.
func (o *SalaryStatisticsResultingParameters) GetPositionsCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionsCount, true
}

// SetPositionsCount sets field value
func (o *SalaryStatisticsResultingParameters) SetPositionsCount(v float32) {
	o.PositionsCount = v
}

// GetSources returns the Sources field value
func (o *SalaryStatisticsResultingParameters) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *SalaryStatisticsResultingParameters) GetSourcesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *SalaryStatisticsResultingParameters) SetSources(v []string) {
	o.Sources = v
}

// GetSpecialities returns the Specialities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalaryStatisticsResultingParameters) GetSpecialities() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Specialities
}

// GetSpecialitiesOk returns a tuple with the Specialities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalaryStatisticsResultingParameters) GetSpecialitiesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Specialities) {
		return nil, false
	}
	return o.Specialities, true
}

// HasSpecialities returns a boolean if a field has been set.
func (o *SalaryStatisticsResultingParameters) HasSpecialities() bool {
	if o != nil && !IsNil(o.Specialities) {
		return true
	}

	return false
}

// SetSpecialities gets a reference to the given []IncludesIdName and assigns it to the Specialities field.
func (o *SalaryStatisticsResultingParameters) SetSpecialities(v []IncludesIdName) {
	o.Specialities = v
}

func (o SalaryStatisticsResultingParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalaryStatisticsResultingParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["areas"] = o.Areas
	if o.EmployeeLevels != nil {
		toSerialize["employee_levels"] = o.EmployeeLevels
	}
	toSerialize["employers_count"] = o.EmployersCount
	if o.ExcludedAreas != nil {
		toSerialize["excluded_areas"] = o.ExcludedAreas
	}
	if o.IndirectCalculation.IsSet() {
		toSerialize["indirect_calculation"] = o.IndirectCalculation.Get()
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	toSerialize["positions_count"] = o.PositionsCount
	toSerialize["sources"] = o.Sources
	if o.Specialities != nil {
		toSerialize["specialities"] = o.Specialities
	}
	return toSerialize, nil
}

func (o *SalaryStatisticsResultingParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"areas",
		"employers_count",
		"positions_count",
		"sources",
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

	varSalaryStatisticsResultingParameters := _SalaryStatisticsResultingParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSalaryStatisticsResultingParameters)

	if err != nil {
		return err
	}

	*o = SalaryStatisticsResultingParameters(varSalaryStatisticsResultingParameters)

	return err
}

type NullableSalaryStatisticsResultingParameters struct {
	value *SalaryStatisticsResultingParameters
	isSet bool
}

func (v NullableSalaryStatisticsResultingParameters) Get() *SalaryStatisticsResultingParameters {
	return v.value
}

func (v *NullableSalaryStatisticsResultingParameters) Set(val *SalaryStatisticsResultingParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSalaryStatisticsResultingParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSalaryStatisticsResultingParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalaryStatisticsResultingParameters(val *SalaryStatisticsResultingParameters) *NullableSalaryStatisticsResultingParameters {
	return &NullableSalaryStatisticsResultingParameters{value: val, isSet: true}
}

func (v NullableSalaryStatisticsResultingParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalaryStatisticsResultingParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


