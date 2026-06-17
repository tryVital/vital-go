## v2.0.0 - 2026-04-02
* The `GetWheelchairUse()` accessor method has been removed from both `ClientFacingActivity` and `ClientFacingProfile`. Callers that reference these methods must update their code; if direct access to the `WheelchairUse` field is still required, access the struct field directly (e.g. `obj.WheelchairUse`).
* Three new query expression types are now available across `AggregateExprArg`, `QueryGroupByItem`, and `QuerySelectItem`: `DerivedReadinessColumnExpr` (with a `DerivedReadinessColumnExprDerivedReadiness` enum), `AwakeningsValueMacroExpr`, and `InsulinInjectionTimeseriesExpr` (with an `InsulinInjectionTimeseriesExprField` enum). The `insulin_injection` timeseries variant has been promoted from `IntervalTimeseriesExpr` into its own dedicated type. **Breaking:** all three visitor interfaces (`AggregateExprArgVisitor`, `QueryGroupByItemVisitor`, `QuerySelectItemVisitor`) now require three additional `Visit*` methods; existing implementations must add `VisitDerivedReadinessColumnExpr`, `VisitAwakeningsValueMacroExpr`, and `VisitInsulinInjectionTimeseriesExpr`. The `Placeholder` type's `GetPlaceholder()` / `SetPlaceholder()` API has been replaced with a `Placeholder()` literal accessor, and the removed `IntervalTimeseriesExprTimeseriesInsulinInjection` constant must be migrated to `InsulinInjectionTimeseriesExpr`.
* **Breaking changes** to lab report parsing types:
* `ParsingJob.JobId` field, `GetJobId()`, and `SetJobId()` have been removed. A new `FailureReason *ParsingJobFailureReason` field (with getter/setter) and a new `ParsingJobFailureReason` enum type have been added in their place.
* `ResultMetadata` fields `PatientFirstName`, `PatientLastName`, `Dob`, and `LabName` have changed from `string` to `*string`; callers must update all reads and writes (including `Set*` method calls) to use pointer values.
* `BodyCreateLabReportParserJob.File` has changed from `io.Reader` to `[]io.Reader` to support multiple file uploads.
* A new optional `Gender *ResultMetadataGender` field has been added to `ResultMetadata`, along with a new `ResultMetadataGender` enum type.
* Several breaking changes have been made to bulk operation request types. The `TeamId` field and `SetTeamId()` method have been removed from `BulkExportConnectionsBody`, `BulkImportConnectionsBody`, `BulkPauseConnectionsBody`, `BulkTriggerHistoricalPullBody`, and `LinkListBulkOpsRequest`. The exported types `LinkBulkExportRequestTeamId`, `LinkBulkImportRequestTeamId`, `LinkBulkPauseRequestTeamId`, `LinkBulkTriggerHistoricalPullRequestTeamId`, and `LinkListBulkOpsRequestTeamId` have also been removed — callers referencing these types must update their code.
* New capabilities added: `ManualProvidersSamsungHealth` and `PasswordProvidersTandemSource` enum constants are now available, `ClientFacingSleep` includes a new `RecoveryReadinessScore` field, and `ManualConnectionData` gains `VitalIosSdkVersion`, `VitalAndroidSdkVersion`, and `GrantedPermissions` fields.
* Several breaking changes and new features have been introduced.
**Breaking changes:**
* `UserRefreshErrorResponse.Success` is no longer a public struct field. Access it via the new `Success()` method instead.
* `HistoricalPullCompleted.IsFinal` is no longer a public struct field. Access it via the new `IsFinal()` method instead.
**New capabilities:**
* New `BiomarkerResult`, `MissingBiomarkerResult`, and `ParentBiomarkerData` types for structured lab result data.
* New `LabResultsRaw`, `LabResultsMetadata`, `SampleData`, `PerformingLaboratory`, and `ClinicalInformation` types for raw lab result payloads.
* New `ClientFacingLabReportParsingJobCreatedEvent` and `ClientFacingLabReportParsingJobUpdatedEvent` webhook event types.
* `ClientFacingOrder` now includes `ClinicalNotes`, `LastEvent`, `Origin`, and `OrderTransaction` fields.
* `ClientFacingOrderEvent` now includes a `StatusDetail` field.
* `Address` and `PatientAddressCompatible` now include an `AccessNotes` field.
* New `ClientFacingOrderTransaction` and `ClientFacingOrderInTransaction` types for order transaction tracking.
* New enum values: `OrderLowLevelStatus`, `OrderOrigin`, `OrderTransactionStatus`, `OrderStatusDetail`, `ResultType`, `FailureType`.
* Several breaking changes are included in this release:
* `UserInfo.Address` and `UserInfoCreateRequest.Address` now use `*UserAddress` instead of `*Address`. Callers must update references from `*Address` to `*UserAddress`.
* `UserRefreshSuccessResponse.Success` (exported field) and `SetSuccess` method have been removed. Use the new `Success() bool` method instead.
* `BookPscAppointment` now accepts `*LabTestsBookPscAppointmentRequest` instead of `*AppointmentBookingRequest`; update call sites accordingly.
* `ParserCreateJob` (`labreport`) now expects a slice of files rather than a single file.
* New capabilities also added: three new client sub-services (`Compendium`, `LabAccount`, `OrderTransaction`), new `UserAddress` type, new `ClientFacingInsulinInjectionSample` delivery fields (`DeliveryMode`, `DeliveryForm`, `BolusPurpose`), and a new `ClientFacingDeviceSourceTypeInsulinPump` enum constant.
* The SDK now includes types for searching and converting lab compendium data. New structs (`SearchCompendiumBody`, `ConvertCompendiumBody`, `SearchCompendiumResponse`, `ConvertCompendiumResponse`) and supporting types (`CanonicalCandidate`, `PerLabCandidate`, `RelatedCandidate`, `ProviderIdConversionResponse`) are available, along with `CompendiumSearchLabs` and `SearchMode` enums to support lab test identifier lookups and cross-lab conversions.
* The SDK now includes two new API client packages: `compendium` (with `Search` and `Convert` methods) and `labaccount` (with a `GetTeamLabAccounts` method). New supporting types have been added including `ClientFacingLabAccount`, `GetTeamLabAccountsResponse`, `LabAccountDelegatedFlow`, `LabAccountStatus`, and `UsState`. The `insulin_pump` source type constant is now available on `ClientFacingElectrocardiogramSourceType` and `ClientFacingMenstrualCycleSourceType`.
* The SDK now includes a new `ordertransaction` package with a `Client` that exposes `GetTransaction`, `GetTransactionResult`, and `GetTransactionResultPdf` methods for retrieving order transaction details, lab results, and PDF reports. New types `GetOrderTransactionResponse` and `OrderSummary` are available in the top-level package, and `ClientFacingSleepCycleSourceType` gains a new `insulin_pump` value.

