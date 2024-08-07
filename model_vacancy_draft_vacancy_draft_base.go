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

// checks if the VacancyDraftVacancyDraftBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyDraftVacancyDraftBase{}

// VacancyDraftVacancyDraftBase struct for VacancyDraftVacancyDraftBase
type VacancyDraftVacancyDraftBase struct {
	AutoPublication NullableVacancyDraftAutoPublicationState `json:"auto_publication,omitempty"`
	// Процент заполнения черновика
	CompletedFieldsPercentage float32 `json:"completed_fields_percentage"`
	// Идентификатор черновика
	DraftId string `json:"draft_id"`
	// Массив объектов с информацией о том, каких публикаций не хватает на счету для публикации вакансии из данного черновика
	InsufficientPublications []VacancyDraftPublications `json:"insufficient_publications,omitempty"`
	// Массив объектов с информацией о том, какие квоты превышены
	InsufficientQuotas []VacancyDraftPublications `json:"insufficient_quotas,omitempty"`
	// Время изменения черновика (в формате [ISO 8601](https://ru.wikipedia.org/wiki/ISO_8601) с точностью до секунды `YYYY-MM-DDThh:mm:ss±hhmm`)
	LastChangeTime NullableString `json:"last_change_time,omitempty"`
	// Готовность черновика к публикации
	PublicationReady bool `json:"publication_ready"`
	// Массив объектов с информацией о необходимых публикациях на счету
	RequiredPublications []VacancyDraftPublications `json:"required_publications,omitempty"`
	// Время запланированной публикации вакансии (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: `YYYY-MM-DDThh:mm:ss±hhmm`
	ScheduledAt string `json:"scheduled_at" validate:"regexp=^\\\\d{4}-[01]\\\\d-[0-3]\\\\dT[0-2]\\\\d:[0-5]\\\\d:[0-5]\\\\d(?:\\\\.\\\\d+)?((\\\\+|-)[0-1][0-9]{3})?$"`
}

type _VacancyDraftVacancyDraftBase VacancyDraftVacancyDraftBase

// NewVacancyDraftVacancyDraftBase instantiates a new VacancyDraftVacancyDraftBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyDraftVacancyDraftBase(completedFieldsPercentage float32, draftId string, publicationReady bool, scheduledAt string) *VacancyDraftVacancyDraftBase {
	this := VacancyDraftVacancyDraftBase{}
	this.CompletedFieldsPercentage = completedFieldsPercentage
	this.DraftId = draftId
	this.PublicationReady = publicationReady
	this.ScheduledAt = scheduledAt
	return &this
}

// NewVacancyDraftVacancyDraftBaseWithDefaults instantiates a new VacancyDraftVacancyDraftBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyDraftVacancyDraftBaseWithDefaults() *VacancyDraftVacancyDraftBase {
	this := VacancyDraftVacancyDraftBase{}
	return &this
}

// GetAutoPublication returns the AutoPublication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftBase) GetAutoPublication() VacancyDraftAutoPublicationState {
	if o == nil || IsNil(o.AutoPublication.Get()) {
		var ret VacancyDraftAutoPublicationState
		return ret
	}
	return *o.AutoPublication.Get()
}

// GetAutoPublicationOk returns a tuple with the AutoPublication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftBase) GetAutoPublicationOk() (*VacancyDraftAutoPublicationState, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoPublication.Get(), o.AutoPublication.IsSet()
}

// HasAutoPublication returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftBase) HasAutoPublication() bool {
	if o != nil && o.AutoPublication.IsSet() {
		return true
	}

	return false
}

// SetAutoPublication gets a reference to the given NullableVacancyDraftAutoPublicationState and assigns it to the AutoPublication field.
func (o *VacancyDraftVacancyDraftBase) SetAutoPublication(v VacancyDraftAutoPublicationState) {
	o.AutoPublication.Set(&v)
}
// SetAutoPublicationNil sets the value for AutoPublication to be an explicit nil
func (o *VacancyDraftVacancyDraftBase) SetAutoPublicationNil() {
	o.AutoPublication.Set(nil)
}

// UnsetAutoPublication ensures that no value is present for AutoPublication, not even an explicit nil
func (o *VacancyDraftVacancyDraftBase) UnsetAutoPublication() {
	o.AutoPublication.Unset()
}

// GetCompletedFieldsPercentage returns the CompletedFieldsPercentage field value
func (o *VacancyDraftVacancyDraftBase) GetCompletedFieldsPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CompletedFieldsPercentage
}

// GetCompletedFieldsPercentageOk returns a tuple with the CompletedFieldsPercentage field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftBase) GetCompletedFieldsPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedFieldsPercentage, true
}

// SetCompletedFieldsPercentage sets field value
func (o *VacancyDraftVacancyDraftBase) SetCompletedFieldsPercentage(v float32) {
	o.CompletedFieldsPercentage = v
}

// GetDraftId returns the DraftId field value
func (o *VacancyDraftVacancyDraftBase) GetDraftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DraftId
}

// GetDraftIdOk returns a tuple with the DraftId field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftBase) GetDraftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftId, true
}

// SetDraftId sets field value
func (o *VacancyDraftVacancyDraftBase) SetDraftId(v string) {
	o.DraftId = v
}

// GetInsufficientPublications returns the InsufficientPublications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftBase) GetInsufficientPublications() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.InsufficientPublications
}

// GetInsufficientPublicationsOk returns a tuple with the InsufficientPublications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftBase) GetInsufficientPublicationsOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.InsufficientPublications) {
		return nil, false
	}
	return o.InsufficientPublications, true
}

// HasInsufficientPublications returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftBase) HasInsufficientPublications() bool {
	if o != nil && !IsNil(o.InsufficientPublications) {
		return true
	}

	return false
}

// SetInsufficientPublications gets a reference to the given []VacancyDraftPublications and assigns it to the InsufficientPublications field.
func (o *VacancyDraftVacancyDraftBase) SetInsufficientPublications(v []VacancyDraftPublications) {
	o.InsufficientPublications = v
}

// GetInsufficientQuotas returns the InsufficientQuotas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftBase) GetInsufficientQuotas() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.InsufficientQuotas
}

// GetInsufficientQuotasOk returns a tuple with the InsufficientQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftBase) GetInsufficientQuotasOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.InsufficientQuotas) {
		return nil, false
	}
	return o.InsufficientQuotas, true
}

// HasInsufficientQuotas returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftBase) HasInsufficientQuotas() bool {
	if o != nil && !IsNil(o.InsufficientQuotas) {
		return true
	}

	return false
}

// SetInsufficientQuotas gets a reference to the given []VacancyDraftPublications and assigns it to the InsufficientQuotas field.
func (o *VacancyDraftVacancyDraftBase) SetInsufficientQuotas(v []VacancyDraftPublications) {
	o.InsufficientQuotas = v
}

// GetLastChangeTime returns the LastChangeTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftBase) GetLastChangeTime() string {
	if o == nil || IsNil(o.LastChangeTime.Get()) {
		var ret string
		return ret
	}
	return *o.LastChangeTime.Get()
}

// GetLastChangeTimeOk returns a tuple with the LastChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftBase) GetLastChangeTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastChangeTime.Get(), o.LastChangeTime.IsSet()
}

// HasLastChangeTime returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftBase) HasLastChangeTime() bool {
	if o != nil && o.LastChangeTime.IsSet() {
		return true
	}

	return false
}

// SetLastChangeTime gets a reference to the given NullableString and assigns it to the LastChangeTime field.
func (o *VacancyDraftVacancyDraftBase) SetLastChangeTime(v string) {
	o.LastChangeTime.Set(&v)
}
// SetLastChangeTimeNil sets the value for LastChangeTime to be an explicit nil
func (o *VacancyDraftVacancyDraftBase) SetLastChangeTimeNil() {
	o.LastChangeTime.Set(nil)
}

// UnsetLastChangeTime ensures that no value is present for LastChangeTime, not even an explicit nil
func (o *VacancyDraftVacancyDraftBase) UnsetLastChangeTime() {
	o.LastChangeTime.Unset()
}

// GetPublicationReady returns the PublicationReady field value
func (o *VacancyDraftVacancyDraftBase) GetPublicationReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicationReady
}

// GetPublicationReadyOk returns a tuple with the PublicationReady field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftBase) GetPublicationReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicationReady, true
}

// SetPublicationReady sets field value
func (o *VacancyDraftVacancyDraftBase) SetPublicationReady(v bool) {
	o.PublicationReady = v
}

// GetRequiredPublications returns the RequiredPublications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftBase) GetRequiredPublications() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.RequiredPublications
}

// GetRequiredPublicationsOk returns a tuple with the RequiredPublications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftBase) GetRequiredPublicationsOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.RequiredPublications) {
		return nil, false
	}
	return o.RequiredPublications, true
}

// HasRequiredPublications returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftBase) HasRequiredPublications() bool {
	if o != nil && !IsNil(o.RequiredPublications) {
		return true
	}

	return false
}

// SetRequiredPublications gets a reference to the given []VacancyDraftPublications and assigns it to the RequiredPublications field.
func (o *VacancyDraftVacancyDraftBase) SetRequiredPublications(v []VacancyDraftPublications) {
	o.RequiredPublications = v
}

// GetScheduledAt returns the ScheduledAt field value
func (o *VacancyDraftVacancyDraftBase) GetScheduledAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftBase) GetScheduledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledAt, true
}

// SetScheduledAt sets field value
func (o *VacancyDraftVacancyDraftBase) SetScheduledAt(v string) {
	o.ScheduledAt = v
}

func (o VacancyDraftVacancyDraftBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyDraftVacancyDraftBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoPublication.IsSet() {
		toSerialize["auto_publication"] = o.AutoPublication.Get()
	}
	toSerialize["completed_fields_percentage"] = o.CompletedFieldsPercentage
	toSerialize["draft_id"] = o.DraftId
	if o.InsufficientPublications != nil {
		toSerialize["insufficient_publications"] = o.InsufficientPublications
	}
	if o.InsufficientQuotas != nil {
		toSerialize["insufficient_quotas"] = o.InsufficientQuotas
	}
	if o.LastChangeTime.IsSet() {
		toSerialize["last_change_time"] = o.LastChangeTime.Get()
	}
	toSerialize["publication_ready"] = o.PublicationReady
	if o.RequiredPublications != nil {
		toSerialize["required_publications"] = o.RequiredPublications
	}
	toSerialize["scheduled_at"] = o.ScheduledAt
	return toSerialize, nil
}

func (o *VacancyDraftVacancyDraftBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completed_fields_percentage",
		"draft_id",
		"publication_ready",
		"scheduled_at",
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

	varVacancyDraftVacancyDraftBase := _VacancyDraftVacancyDraftBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacancyDraftVacancyDraftBase)

	if err != nil {
		return err
	}

	*o = VacancyDraftVacancyDraftBase(varVacancyDraftVacancyDraftBase)

	return err
}

type NullableVacancyDraftVacancyDraftBase struct {
	value *VacancyDraftVacancyDraftBase
	isSet bool
}

func (v NullableVacancyDraftVacancyDraftBase) Get() *VacancyDraftVacancyDraftBase {
	return v.value
}

func (v *NullableVacancyDraftVacancyDraftBase) Set(val *VacancyDraftVacancyDraftBase) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyDraftVacancyDraftBase) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyDraftVacancyDraftBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyDraftVacancyDraftBase(val *VacancyDraftVacancyDraftBase) *NullableVacancyDraftVacancyDraftBase {
	return &NullableVacancyDraftVacancyDraftBase{value: val, isSet: true}
}

func (v NullableVacancyDraftVacancyDraftBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyDraftVacancyDraftBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


