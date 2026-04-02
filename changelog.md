## v2.0.0 - 2026-04-02
* The `GetWheelchairUse()` accessor method has been removed from both `ClientFacingActivity` and `ClientFacingProfile`. Callers that reference these methods must be updated — remove calls to `GetWheelchairUse()` or access the `WheelchairUse` field directly on the struct instead.
* The SDK now supports three new query expression types: `DerivedReadinessColumnExpr` for readiness-related columns (chronotype, sleep score, recovery score, recovery zone, stress score, strain score, strain zone), `InsulinInjectionTimeseriesExpr` for insulin injection timeseries data (replacing the former `IntervalTimeseriesExprTimeseriesInsulinInjection` constant), and `AwakeningsValueMacroExpr` for awakenings macro expressions. All three are available as variants in `AggregateExprArg`, `QueryGroupByItem`, and `QuerySelectItem`.
**Breaking changes:** The `AggregateExprArgVisitor`, `QueryGroupByItemVisitor`, and `QuerySelectItemVisitor` interfaces each gain three new required `Visit*` methods (`VisitDerivedReadinessColumnExpr`, `VisitAwakeningsValueMacroExpr`, `VisitInsulinInjectionTimeseriesExpr`) — any existing implementations must add these methods. The `IntervalTimeseriesExprTimeseriesInsulinInjection` constant has been removed. The `Placeholder` struct's `Placeholder` field is now a private literal (accessed via the `Placeholder()` method); the exported `GetPlaceholder()` and `SetPlaceholder()` methods have been removed. `QueryBatch.Timeframe` and `QueryBatch.Queries` JSON tags no longer use `omitempty`.
* **Breaking changes** to `ParsingJob` and `ResultMetadata`:
* `ParsingJob.JobId` field and its `GetJobId`/`SetJobId` methods have been removed. The new `FailureReason` field (type `ParsingJobFailureReason`) is available in its place, along with a new `ParsingJobFailureReason` enum (`invalid_input`, `low_quality`, `not_english`).
* `ResultMetadata` fields `PatientFirstName`, `PatientLastName`, `Dob`, and `LabName` are now `*string` (previously `string`); callers must dereference these values and update calls to their corresponding `Set*` methods to pass pointer arguments.
* `BodyCreateLabReportParserJob.File` changed from `io.Reader` to `[]io.Reader`; wrap existing single-reader usages in a slice.
* A new `Gender` field (`*ResultMetadataGender`) has been added to `ResultMetadata`, along with a new `ResultMetadataGender` enum (`male`, `female`, `other`).
* Several breaking changes have been made to bulk-operation request types. The `TeamId` field and `SetTeamId` method have been removed from `BulkExportConnectionsBody`, `BulkImportConnectionsBody`, `BulkPauseConnectionsBody`, `BulkTriggerHistoricalPullBody`, and `LinkListBulkOpsRequest`; callers referencing these fields must remove them. The exported types `LinkBulkExportRequestTeamId`, `LinkBulkImportRequestTeamId`, `LinkBulkPauseRequestTeamId`, `LinkBulkTriggerHistoricalPullRequestTeamId`, and `LinkListBulkOpsRequestTeamId` have also been removed. Additionally, `ClientFacingSleep` now includes a new `RecoveryReadinessScore` field, `ManualConnectionData` gains `VitalIosSdkVersion`, `VitalAndroidSdkVersion`, and `GrantedPermissions` fields, and new enum values `ManualProvidersSamsungHealth` and `PasswordProvidersTandemSource` are available.
* The SDK now includes a wide set of new types and fields for lab ordering, order tracking, and lab result parsing.
**New types and events:**
* `BiomarkerResult`, `MissingBiomarkerResult`, and `ParentBiomarkerData` for structured lab result data.
* `ClientFacingLabReportParsingJobCreatedEvent` and `ClientFacingLabReportParsingJobUpdatedEvent` webhook events for lab report parsing jobs.
* `ClientFacingOrderTransaction` and `ClientFacingOrderInTransaction` for order transaction tracking.
* `LabResultsRaw`, `LabResultsMetadata`, `SampleData`, `PerformingLaboratory`, and `ClinicalInformation` for raw lab result payloads.
* `UtcTimestampWithTimezoneOffset` for timezone-aware timestamps in lab result sample data.
**New fields:**
* `ClientFacingOrder` now includes `last_event`, `clinical_notes`, `origin`, and `order_transaction`.
* `ClientFacingOrderEvent` now includes `status_detail`.
* `Address`, `PatientAddressCompatible`, and `PatientAddressWithValidation` now include `access_notes`.
**New enum values:**
* `OrderLowLevelStatus`, `OrderOrigin`, `OrderStatusDetail`, and `OrderTransactionStatus` enums added.
* New corrected order status constants (e.g., `completed.testkit.corrected`).
* Additional `Labs`, `Providers`, and `PayorCodeExternalProvider` values.
* Several breaking changes are included in this release:
* `UserInfo.GetAddress()` and `UserInfo.SetAddress()` (and the same methods on `UserInfoCreateRequest`) now use `*UserAddress` instead of `*Address`. Update any code that passes or receives an `*Address` to use `*UserAddress`.
* `UserRefreshSuccessResponse` no longer exposes the `Success` field, `GetSuccess()`, or `SetSuccess()`. The `success` JSON literal is now validated internally; use the new `Success() bool` method to read it.
* `BookPscAppointment` on both `labtests.Client` and `labtests.RawClient` now accepts `*vitalgo.LabTestsBookPscAppointmentRequest` instead of `*vitalgo.AppointmentBookingRequest`.
* `labreport.ParserCreateJob` now expects `request.File` to be a slice rather than a single file.
* Bulk link operations (`BulkImport`, `BulkTriggerHistoricalPull`, `BulkExport`, `BulkPause`) no longer append query parameters from the request struct.
* New additions include `Compendium`, `LabAccount`, and `OrderTransaction` sub-clients on the root `Client`, new insulin delivery enum types (`DeliveryMode`, `DeliveryForm`, `BolusPurpose`) on `ClientFacingInsulinInjectionSample`, and a new `ClientFacingDeviceSourceTypeInsulinPump` enum value.
* The SDK now includes types for lab test compendium search and conversion. New types include `SearchCompendiumBody`, `ConvertCompendiumBody`, `SearchCompendiumResponse`, `ConvertCompendiumResponse`, `CanonicalCandidate`, `PerLabCandidate`, `RelatedCandidate`, and `ProviderIdConversionResponse`, along with the `CompendiumSearchLabs` and `SearchMode` enums.
* The SDK now includes two new API client packages: `compendium` (with `Search` and `Convert` methods for lab test data) and `labaccount` (with `GetTeamLabAccounts` for managing team lab accounts). New types `ClientFacingLabAccount`, `GetTeamLabAccountsResponse`, `LabAccountDelegatedFlow`, `LabAccountStatus`, and `UsState` are also available. The `insulin_pump` source type has been added to the electrocardiogram and menstrual cycle source type enums.
* The SDK now includes a new `ordertransaction` package providing `GetTransaction`, `GetTransactionResult`, and `GetTransactionResultPdf` methods for accessing order transaction data and lab results via the `/v3/order_transaction` endpoints. New exported types `GetOrderTransactionResponse` and `OrderSummary` are available in the top-level package. A new `insulin_pump` source type constant (`ClientFacingSleepCycleSourceTypeInsulinPump`) has also been added to the sleep cycle source type enum.

