# NegotiationsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]NegotiationsListItem**](NegotiationsListItem.md) | Список откликов | 
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Collections** | [**[]NegotiationsNegotiationsCollection**](NegotiationsNegotiationsCollection.md) | Коллекции откликов/приглашений для данной вакансии | 
**EmployerStates** | [**[]EmployersEmployersState**](EmployersEmployersState.md) | Состояния [откликов/приглашений](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#term-employer-state) вакансии для работодателя  | 
**GeneratedCollections** | [**[]NegotiationsNegotiationsCollection**](NegotiationsNegotiationsCollection.md) | Сгенерированные коллекции откликов/приглашений для данной вакансии | 

## Methods

### NewNegotiationsListResponse

`func NewNegotiationsListResponse(items []NegotiationsListItem, found float32, page float32, pages float32, perPage float32, collections []NegotiationsNegotiationsCollection, employerStates []EmployersEmployersState, generatedCollections []NegotiationsNegotiationsCollection, ) *NegotiationsListResponse`

NewNegotiationsListResponse instantiates a new NegotiationsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsListResponseWithDefaults

`func NewNegotiationsListResponseWithDefaults() *NegotiationsListResponse`

NewNegotiationsListResponseWithDefaults instantiates a new NegotiationsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *NegotiationsListResponse) GetItems() []NegotiationsListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NegotiationsListResponse) GetItemsOk() (*[]NegotiationsListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NegotiationsListResponse) SetItems(v []NegotiationsListItem)`

SetItems sets Items field to given value.


### GetFound

`func (o *NegotiationsListResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *NegotiationsListResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *NegotiationsListResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *NegotiationsListResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NegotiationsListResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NegotiationsListResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *NegotiationsListResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NegotiationsListResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NegotiationsListResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *NegotiationsListResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *NegotiationsListResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *NegotiationsListResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetCollections

`func (o *NegotiationsListResponse) GetCollections() []NegotiationsNegotiationsCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *NegotiationsListResponse) GetCollectionsOk() (*[]NegotiationsNegotiationsCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *NegotiationsListResponse) SetCollections(v []NegotiationsNegotiationsCollection)`

SetCollections sets Collections field to given value.


### GetEmployerStates

`func (o *NegotiationsListResponse) GetEmployerStates() []EmployersEmployersState`

GetEmployerStates returns the EmployerStates field if non-nil, zero value otherwise.

### GetEmployerStatesOk

`func (o *NegotiationsListResponse) GetEmployerStatesOk() (*[]EmployersEmployersState, bool)`

GetEmployerStatesOk returns a tuple with the EmployerStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerStates

`func (o *NegotiationsListResponse) SetEmployerStates(v []EmployersEmployersState)`

SetEmployerStates sets EmployerStates field to given value.


### GetGeneratedCollections

`func (o *NegotiationsListResponse) GetGeneratedCollections() []NegotiationsNegotiationsCollection`

GetGeneratedCollections returns the GeneratedCollections field if non-nil, zero value otherwise.

### GetGeneratedCollectionsOk

`func (o *NegotiationsListResponse) GetGeneratedCollectionsOk() (*[]NegotiationsNegotiationsCollection, bool)`

GetGeneratedCollectionsOk returns a tuple with the GeneratedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCollections

`func (o *NegotiationsListResponse) SetGeneratedCollections(v []NegotiationsNegotiationsCollection)`

SetGeneratedCollections sets GeneratedCollections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


