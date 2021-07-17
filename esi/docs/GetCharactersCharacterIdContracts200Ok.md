# GetCharactersCharacterIdContracts200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptorId** | **int32** | Who will accept the contract | [default to null]
**AssigneeId** | **int32** | ID to whom the contract is assigned, can be corporation or character ID | [default to null]
**Availability** | **string** | To whom the contract is available | [default to null]
**Buyout** | **float64** | Buyout price (for Auctions only) | [optional] [default to null]
**Collateral** | **float64** | Collateral price (for Couriers only) | [optional] [default to null]
**ContractId** | **int32** | contract_id integer | [default to null]
**DateAccepted** | [**time.Time**](time.Time.md) | Date of confirmation of contract | [optional] [default to null]
**DateCompleted** | [**time.Time**](time.Time.md) | Date of completed of contract | [optional] [default to null]
**DateExpired** | [**time.Time**](time.Time.md) | Expiration date of the contract | [default to null]
**DateIssued** | [**time.Time**](time.Time.md) | Сreation date of the contract | [default to null]
**DaysToComplete** | **int32** | Number of days to perform the contract | [optional] [default to null]
**EndLocationId** | **int64** | End location ID (for Couriers contract) | [optional] [default to null]
**ForCorporation** | **bool** | true if the contract was issued on behalf of the issuer&#x27;s corporation | [default to null]
**IssuerCorporationId** | **int32** | Character&#x27;s corporation ID for the issuer | [default to null]
**IssuerId** | **int32** | Character ID for the issuer | [default to null]
**Price** | **float64** | Price of contract (for ItemsExchange and Auctions) | [optional] [default to null]
**Reward** | **float64** | Remuneration for contract (for Couriers only) | [optional] [default to null]
**StartLocationId** | **int64** | Start location ID (for Couriers contract) | [optional] [default to null]
**Status** | **string** | Status of the the contract | [default to null]
**Title** | **string** | Title of the contract | [optional] [default to null]
**Type_** | **string** | Type of the contract | [default to null]
**Volume** | **float64** | Volume of items in the contract | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

