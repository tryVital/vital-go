# Reference
## Link
<details><summary><code>client.Link.ListBulkOps() -> *v505.BulkOpsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkListBulkOpsRequest{
        NextCursor: v505.String(
            "next_cursor",
        ),
        PageSize: v505.Int(
            1,
        ),
        TeamId: v505.LinkListBulkOpsRequestTeamIdInferFromContext.Ptr(),
    }
client.Link.ListBulkOps(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**nextCursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**teamId:** `*v505.LinkListBulkOpsRequestTeamId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.BulkImport(request) -> *v505.BulkImportConnectionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BulkImportConnectionsBody{
        TeamId: v505.LinkBulkImportRequestTeamIdInferFromContext.Ptr(),
        Provider: v505.OAuthProvidersOura,
        Connections: []*v505.ConnectionRecipe{
            &v505.ConnectionRecipe{
                UserId: "user_id",
                AccessToken: "access_token",
                RefreshToken: "refresh_token",
                ProviderId: "provider_id",
                ExpiresAt: 1,
            },
        },
    }
client.Link.BulkImport(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.LinkBulkImportRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.OAuthProviders` 
    
</dd>
</dl>

<dl>
<dd>

**connections:** `[]*v505.ConnectionRecipe` 
    
</dd>
</dl>

<dl>
<dd>

**waitForCompletion:** `*bool` 


Whether or not the endpoint should wait for the Bulk Op to complete before responding.

When `wait_for_completion` is enabled, the endpoint may respond 200 OK if the Bulk Op takes less than 20 seconds to complete.

Otherwise, the endpoint always responds with 202 Created once the submitted data have been enqueued successfully. You can use
the [List Bulk Ops](https://docs.tryvital.io/api-reference/link/list-bulk-ops) endpoint to inspect the progress of the Bulk Op.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.BulkTriggerHistoricalPull(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BulkTriggerHistoricalPullBody{
        TeamId: v505.LinkBulkTriggerHistoricalPullRequestTeamIdInferFromContext.Ptr(),
        UserIds: []string{
            "user_ids",
        },
        Provider: v505.OAuthProvidersOura,
    }
client.Link.BulkTriggerHistoricalPull(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.LinkBulkTriggerHistoricalPullRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.OAuthProviders` 
    
</dd>
</dl>

<dl>
<dd>

**waitForCompletion:** `*bool` 


Whether or not the endpoint should wait for the Bulk Op to complete before responding.

When `wait_for_completion` is enabled, the endpoint may respond 200 OK if the Bulk Op takes less than 20 seconds to complete.

Otherwise, the endpoint always responds with 202 Created once the submitted data have been enqueued successfully. You can use
the [List Bulk Ops](https://docs.tryvital.io/api-reference/link/list-bulk-ops) endpoint to inspect the progress of the Bulk Op.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.BulkExport(request) -> *v505.BulkExportConnectionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BulkExportConnectionsBody{
        TeamId: v505.LinkBulkExportRequestTeamIdInferFromContext.Ptr(),
        Provider: v505.OAuthProvidersOura,
    }
client.Link.BulkExport(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.LinkBulkExportRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.OAuthProviders` 
    
</dd>
</dl>

<dl>
<dd>

**nextToken:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.BulkPause(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BulkPauseConnectionsBody{
        TeamId: v505.LinkBulkPauseRequestTeamIdInferFromContext.Ptr(),
        UserIds: []string{
            "user_ids",
        },
        Provider: v505.OAuthProvidersOura,
    }
client.Link.BulkPause(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.LinkBulkPauseRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.OAuthProviders` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.Token(request) -> *v505.LinkTokenExchangeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Endpoint to generate a user link token, to be used throughout the vital
link process. The vital link token is a one time use token, that
expires after 10 minutes. If you would like vital-link widget to launch
with a specific provider, pass in a provider in the body. If you would
like to redirect to a custom url after successful or error connection,
pass in your own custom redirect_url parameter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkTokenExchange{
        UserId: "user_id",
    }
client.Link.Token(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — User id returned by vital create user request. This id should be stored in your database against the user and used for all interactions with the vital api.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**redirectUrl:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterOnProviders:** `[]*v505.Providers` 

An allowlist of providers dictating what Vital Link Widget should show to the end user.
If unspecified, all linkable providers are shown.

This has no effect on programmatic Vital Link API usage.
    
</dd>
</dl>

<dl>
<dd>

**onError:** `*string` 

By default, Vital Link Widget input forms for password and email providers have in-built error handling.

Specifying `on_error=redirect` disables this Vital Link Widget UI behaviour — it would
instead redirect to your `redirect_url`, with Link Error parameters injected.

This has no effect on OAuth providers — they always redirect to your `redirect_url`. This also has
no effect on programmatic Vital Link API usage.
    
</dd>
</dl>

<dl>
<dd>

**onClose:** `*string` 

By default, Vital Link Widget closes the window or tab when the user taps the Close button.

Specifying `on_close=redirect` would change the Close button behaviour to redirect to your `redirect_url`
with the `user_cancelled` error specified.

This has no effect on programmatic Vital Link API usage.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.IsTokenValid(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkTokenValidationRequest{
        Token: "token",
    }
client.Link.IsTokenValid(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.CodeCreate() -> *v505.VitalTokenCreatedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generate a token to invite a user of Vital mobile app to your team
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkCodeCreateRequest{
        UserId: "user_id",
        ExpiresAt: v505.Time(
            v505.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Link.CodeCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — When the link code should expire. Defaults to server time plus 1 hour.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.StartConnect(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

REQUEST_SOURCE: VITAL-LINK
Start link token process
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BeginLinkTokenRequest{
        LinkToken: "link_token",
        Provider: v505.ProvidersOura,
    }
client.Link.StartConnect(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**linkToken:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.TokenState() -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

REQUEST_SOURCE: VITAL-LINK
Check link token state - can be hit continuously used as heartbeat
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkTokenStateRequest{}
client.Link.TokenState(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.EmailAuth(request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated. Use `POST /v2/link/provider/email/{provider}` instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.EmailAuthLink{
        Email: "email",
        Provider: v505.ProvidersOura,
        AuthType: v505.AuthTypePassword,
    }
client.Link.EmailAuth(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**authType:** `*v505.AuthType` 
    
</dd>
</dl>

<dl>
<dd>

**region:** `*v505.Region` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.PasswordAuth(request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated. Use `POST /v2/link/provider/password/{provider}` instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.PasswordAuthLink{
        Username: "username",
        Password: "password",
        Provider: v505.ProvidersOura,
        AuthType: v505.AuthTypePassword,
    }
client.Link.PasswordAuth(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**username:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**authType:** `*v505.AuthType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.GenerateOauthLink(OauthProvider) -> *v505.Source</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint generates an OAuth link for oauth provider
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkGenerateOauthLinkRequest{}
client.Link.GenerateOauthLink(
        context.TODO(),
        v505.OAuthProvidersOura.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**oauthProvider:** `*v505.OAuthProviders` 
    
</dd>
</dl>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.ConnectPasswordProvider(Provider, request) -> *v505.ProviderLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This connects auth providers that are password based.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.IndividualProviderData{
        Username: "username",
        Password: "password",
    }
client.Link.ConnectPasswordProvider(
        context.TODO(),
        v505.PasswordProvidersWhoop.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*v505.PasswordProviders` 
    
</dd>
</dl>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**username:** `string` — Username for provider
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` — Password for provider
    
</dd>
</dl>

<dl>
<dd>

**region:** `*v505.Region` — Provider region to authenticate against. Only applicable to specific providers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.CompletePasswordProviderMfa(Provider, request) -> *v505.ProviderLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This connects auth providers that are password based.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CompletePasswordProviderMfaBody{
        MfaCode: "mfa_code",
    }
client.Link.CompletePasswordProviderMfa(
        context.TODO(),
        v505.PasswordProvidersWhoop.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*v505.PasswordProviders` 
    
</dd>
</dl>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**mfaCode:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.ConnectEmailAuthProvider(Provider, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This connects auth providers that are email based.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.EmailProviderAuthLink{
        Email: "email",
    }
client.Link.ConnectEmailAuthProvider(
        context.TODO(),
        v505.EmailProviders(
            "freestyle_libre",
        ),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `v505.EmailProviders` 
    
</dd>
</dl>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**emailProviderAuthLinkProvider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**region:** `*v505.Region` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.GetAllProviders() -> []*v505.SourceLink</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET List of all available providers given the generated link token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LinkGetAllProvidersRequest{}
client.Link.GetAllProviders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.ConnectManualProvider(Provider, request) -> map[string]bool</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ManualConnectionData{
        UserId: "user_id",
    }
client.Link.ConnectManualProvider(
        context.TODO(),
        v505.ManualProvidersBeurerBle.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*v505.ManualProviders` 
    
</dd>
</dl>

<dl>
<dd>

**vitalIosSdkVersion:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**vitalAndroidSdkVersion:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**providerId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**grantedPermissions:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.ConnectDemoProvider(request) -> *v505.DemoConnectionStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

POST Connect the given Vital user to a demo provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.DemoConnectionCreationPayload{
        UserId: "user_id",
        Provider: v505.DemoProvidersAppleHealthKit,
    }
client.Link.ConnectDemoProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — Vital user ID
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.DemoProviders` — Demo provider. For more information, please check out our docs (https://docs.tryvital.io/wearables/providers/test_data)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Electrocardiogram
<details><summary><code>client.Electrocardiogram.Get(UserId) -> *v505.ClientFacingElectrocardiogramResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get electrocardiogram summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ElectrocardiogramGetRequest{
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
        Provider: v505.String(
            "provider",
        ),
    }
client.Electrocardiogram.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SleepCycle
<details><summary><code>client.SleepCycle.Get(UserId) -> *v505.ClientSleepCycleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sleep cycle for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SleepCycleGetRequest{
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
        Provider: v505.String(
            "provider",
        ),
    }
client.SleepCycle.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Profile
<details><summary><code>client.Profile.Get(UserId) -> *v505.ClientFacingProfile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get profile for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ProfileGetRequest{
        Provider: v505.String(
            "provider",
        ),
    }
client.Profile.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profile.GetRaw(UserId) -> *v505.RawProfile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw profile for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ProfileGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
    }
client.Profile.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Devices
<details><summary><code>client.Devices.GetRaw(UserId) -> *v505.RawDevices</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get Devices for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.DevicesGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
    }
client.Devices.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Activity
<details><summary><code>client.Activity.Get(UserId) -> *v505.ClientActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get activity summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ActivityGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Activity.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Activity.GetRaw(UserId) -> *v505.RawActivity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw activity summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ActivityGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Activity.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Workouts
<details><summary><code>client.Workouts.Get(UserId) -> *v505.ClientWorkoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get workout summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.WorkoutsGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Workouts.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workouts.GetRaw(UserId) -> *v505.RawWorkout</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw workout summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.WorkoutsGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Workouts.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workouts.GetByWorkoutId(WorkoutId) -> *v505.ClientFacingStream</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Workouts.GetByWorkoutId(
        context.TODO(),
        "workout_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workoutId:** `string` — The Vital ID for the workout
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sleep
<details><summary><code>client.Sleep.Get(UserId) -> *v505.ClientSleepResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sleep summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SleepGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Sleep.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sleep.GetRaw(UserId) -> *v505.RawSleep</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw sleep summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SleepGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Sleep.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sleep.GetStreamBySleepId(SleepId) -> *v505.ClientFacingSleepStream</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get Sleep stream for a user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sleep.GetStreamBySleepId(
        context.TODO(),
        "sleep_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sleepId:** `string` — The Vital Sleep ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Body
<details><summary><code>client.Body.Get(UserId) -> *v505.ClientBodyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get Body summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BodyGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Body.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Body.GetRaw(UserId) -> *v505.RawBody</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw Body summary for user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BodyGetRawRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Body.GetRaw(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Meal
<details><summary><code>client.Meal.Get(UserId) -> *v505.ClientFacingMealResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get user's meals
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.MealGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Meal.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MenstrualCycle
<details><summary><code>client.MenstrualCycle.Get(UserId) -> *v505.MenstrualCycleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.MenstrualCycleGetRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.MenstrualCycle.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Vitals
<details><summary><code>client.Vitals.WorkoutSwimmingStrokeGrouped(UserId) -> *v505.GroupedWorkoutSwimmingStrokeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWorkoutSwimmingStrokeGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WorkoutSwimmingStrokeGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.WorkoutDistanceGrouped(UserId) -> *v505.GroupedWorkoutDistanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWorkoutDistanceGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WorkoutDistanceGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HeartRateRecoveryOneMinuteGrouped(UserId) -> *v505.GroupedHeartRateRecoveryOneMinuteResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHeartRateRecoveryOneMinuteGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HeartRateRecoveryOneMinuteGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.WaistCircumferenceGrouped(UserId) -> *v505.GroupedWaistCircumferenceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWaistCircumferenceGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WaistCircumferenceGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.LeanBodyMassGrouped(UserId) -> *v505.GroupedLeanBodyMassResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsLeanBodyMassGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.LeanBodyMassGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyMassIndexGrouped(UserId) -> *v505.GroupedBodyMassIndexResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyMassIndexGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyMassIndexGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BasalBodyTemperatureGrouped(UserId) -> *v505.GroupedBasalBodyTemperatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBasalBodyTemperatureGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BasalBodyTemperatureGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HandwashingGrouped(UserId) -> *v505.GroupedHandwashingResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHandwashingGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HandwashingGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.DaylightExposureGrouped(UserId) -> *v505.GroupedDaylightExposureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsDaylightExposureGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.DaylightExposureGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.UvExposureGrouped(UserId) -> *v505.GroupedUvExposureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsUvExposureGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.UvExposureGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.FallGrouped(UserId) -> *v505.GroupedFallResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsFallGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.FallGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.InhalerUsageGrouped(UserId) -> *v505.GroupedInhalerUsageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsInhalerUsageGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.InhalerUsageGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.PeakExpiratoryFlowRateGrouped(UserId) -> *v505.GroupedPeakExpiratoryFlowRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsPeakExpiratoryFlowRateGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.PeakExpiratoryFlowRateGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.ForcedVitalCapacityGrouped(UserId) -> *v505.GroupedForcedVitalCapacityResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsForcedVitalCapacityGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.ForcedVitalCapacityGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.ForcedExpiratoryVolume1Grouped(UserId) -> *v505.GroupedForcedExpiratoryVolume1Response</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsForcedExpiratoryVolume1GroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.ForcedExpiratoryVolume1Grouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.WheelchairPushGrouped(UserId) -> *v505.GroupedWheelchairPushResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWheelchairPushGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WheelchairPushGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.SleepBreathingDisturbanceGrouped(UserId) -> *v505.GroupedSleepBreathingDisturbanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsSleepBreathingDisturbanceGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.SleepBreathingDisturbanceGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.SleepApneaAlertGrouped(UserId) -> *v505.GroupedSleepApneaAlertResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsSleepApneaAlertGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.SleepApneaAlertGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.StandDurationGrouped(UserId) -> *v505.GroupedStandDurationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStandDurationGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.StandDurationGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.StandHourGrouped(UserId) -> *v505.GroupedStandHourResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStandHourGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.StandHourGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HeartRateAlertGrouped(UserId) -> *v505.GroupedHeartRateAlertResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHeartRateAlertGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HeartRateAlertGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.AfibBurdenGrouped(UserId) -> *v505.GroupedAFibBurdenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsAfibBurdenGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.AfibBurdenGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.WorkoutDurationGrouped(UserId) -> *v505.GroupedWorkoutDurationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWorkoutDurationGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WorkoutDurationGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Vo2MaxGrouped(UserId) -> *v505.GroupedVo2MaxResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsVo2MaxGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Vo2MaxGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.StressLevelGrouped(UserId) -> *v505.GroupedStressLevelResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStressLevelGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.StressLevelGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.MindfulnessMinutesGrouped(UserId) -> *v505.GroupedMindfulnessMinutesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsMindfulnessMinutesGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.MindfulnessMinutesGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CaffeineGrouped(UserId) -> *v505.GroupedCaffeineResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaffeineGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CaffeineGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.WaterGrouped(UserId) -> *v505.GroupedWaterResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWaterGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.WaterGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.StepsGrouped(UserId) -> *v505.GroupedStepsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStepsGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.StepsGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.FloorsClimbedGrouped(UserId) -> *v505.GroupedFloorsClimbedResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsFloorsClimbedGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.FloorsClimbedGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.DistanceGrouped(UserId) -> *v505.GroupedDistanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsDistanceGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.DistanceGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CaloriesBasalGrouped(UserId) -> *v505.GroupedCaloriesBasalResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaloriesBasalGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CaloriesBasalGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CaloriesActiveGrouped(UserId) -> *v505.GroupedCaloriesActiveResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaloriesActiveGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CaloriesActiveGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.RespiratoryRateGrouped(UserId) -> *v505.GroupedRespiratoryRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsRespiratoryRateGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.RespiratoryRateGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.NoteGrouped(UserId) -> *v505.GroupedNoteResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsNoteGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.NoteGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.InsulinInjectionGrouped(UserId) -> *v505.GroupedInsulinInjectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsInsulinInjectionGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.InsulinInjectionGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.IgeGrouped(UserId) -> *v505.GroupedIgeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsIgeGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.IgeGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.IggGrouped(UserId) -> *v505.GroupedIggResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsIggGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.IggGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HypnogramGrouped(UserId) -> *v505.GroupedHypnogramResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHypnogramGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HypnogramGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HrvGrouped(UserId) -> *v505.GroupedHrvResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHrvGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HrvGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.HeartrateGrouped(UserId) -> *v505.GroupedHeartRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHeartrateGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.HeartrateGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.GlucoseGrouped(UserId) -> *v505.GroupedGlucoseResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsGlucoseGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.GlucoseGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CholesterolGrouped(UserId) -> *v505.GroupedCholesterolResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CholesterolGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CarbohydratesGrouped(UserId) -> *v505.GroupedCarbohydratesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCarbohydratesGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CarbohydratesGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyTemperatureDeltaGrouped(UserId) -> *v505.GroupedBodyTemperatureDeltaResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyTemperatureDeltaGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyTemperatureDeltaGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyTemperatureGrouped(UserId) -> *v505.GroupedBodyTemperatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyTemperatureGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyTemperatureGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyWeightGrouped(UserId) -> *v505.GroupedBodyWeightResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyWeightGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyWeightGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyFatGrouped(UserId) -> *v505.GroupedBodyFatResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyFatGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyFatGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BloodOxygenGrouped(UserId) -> *v505.GroupedBloodOxygenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBloodOxygenGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BloodOxygenGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.ElectrocardiogramVoltageGrouped(UserId) -> *v505.GroupedElectrocardiogramVoltageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsElectrocardiogramVoltageGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.ElectrocardiogramVoltageGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BloodPressureGrouped(UserId) -> *v505.GroupedBloodPressureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBloodPressureGroupedRequest{
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BloodPressureGrouped(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Vo2Max(UserId) -> []*v505.ClientFacingVo2MaxTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsVo2MaxRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Vo2Max(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.StressLevel(UserId) -> []*v505.ClientFacingStressLevelTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStressLevelRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.StressLevel(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.MindfulnessMinutes(UserId) -> []*v505.ClientFacingMindfulnessMinutesTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsMindfulnessMinutesRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.MindfulnessMinutes(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Caffeine(UserId) -> []*v505.ClientFacingCaffeineTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaffeineRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Caffeine(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Water(UserId) -> []*v505.ClientFacingWaterTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsWaterRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Water(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Steps(UserId) -> []*v505.ClientFacingStepsTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsStepsRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Steps(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.FloorsClimbed(UserId) -> []*v505.ClientFacingFloorsClimbedTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsFloorsClimbedRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.FloorsClimbed(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Distance(UserId) -> []*v505.ClientFacingDistanceTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsDistanceRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Distance(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CaloriesBasal(UserId) -> []*v505.ClientFacingCaloriesBasalTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaloriesBasalRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CaloriesBasal(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CaloriesActive(UserId) -> []*v505.ClientFacingCaloriesActiveTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCaloriesActiveRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CaloriesActive(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.RespiratoryRate(UserId) -> []*v505.ClientFacingRespiratoryRateTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsRespiratoryRateRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.RespiratoryRate(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Ige(UserId) -> []*v505.ClientFacingIgeTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsIgeRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Ige(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Igg(UserId) -> []*v505.ClientFacingIggTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsIggRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Igg(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Hypnogram(UserId) -> []*v505.ClientFacingHypnogramTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHypnogramRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Hypnogram(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Hrv(UserId) -> []*v505.ClientFacingHrvTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHrvRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Hrv(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Heartrate(UserId) -> []*v505.ClientFacingHeartRateTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsHeartrateRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Heartrate(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Glucose(UserId) -> []*v505.ClientFacingGlucoseTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsGlucoseRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Glucose(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CholesterolTriglycerides(UserId) -> []*v505.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolTriglyceridesRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CholesterolTriglycerides(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CholesterolTotal(UserId) -> []*v505.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolTotalRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CholesterolTotal(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CholesterolLdl(UserId) -> []*v505.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolLdlRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CholesterolLdl(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.CholesterolHdl(UserId) -> []*v505.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolHdlRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.CholesterolHdl(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.Cholesterol(UserId) -> []*v505.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsCholesterolRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.Cholesterol(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyWeight(UserId) -> []*v505.ClientFacingBodyWeightTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyWeightRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyWeight(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BodyFat(UserId) -> []*v505.ClientFacingBodyFatTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBodyFatRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BodyFat(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BloodOxygen(UserId) -> []*v505.ClientFacingBloodOxygenTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBloodOxygenRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BloodOxygen(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.ElectrocardiogramVoltage(UserId) -> []*v505.ClientFacingElectrocardiogramVoltageTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsElectrocardiogramVoltageRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.ElectrocardiogramVoltage(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vitals.BloodPressure(UserId) -> []*v505.ClientFacingBloodPressureTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalsBloodPressureRequest{
        Provider: v505.String(
            "provider",
        ),
        StartDate: "start_date",
        EndDate: v505.String(
            "end_date",
        ),
    }
client.Vitals.BloodPressure(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — Provider oura/strava etc
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## User
<details><summary><code>client.User.GetAll() -> *v505.PaginatedUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET All users for team.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserGetAllRequest{
        Offset: v505.Int(
            1,
        ),
        Limit: v505.Int(
            1,
        ),
    }
client.User.GetAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Create(request) -> *v505.ClientFacingUserKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

POST Create a Vital user given a client_user_id and returns the user_id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserCreateBody{
        ClientUserId: "client_user_id",
    }
client.User.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientUserId:** `string` — A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
    
</dd>
</dl>

<dl>
<dd>

**fallbackTimeZone:** `*string` 


    Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
    Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
    
    
</dd>
</dl>

<dl>
<dd>

**fallbackBirthDate:** `*string` — Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
    
</dd>
</dl>

<dl>
<dd>

**ingestionStart:** `*string` — Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>

<dl>
<dd>

**ingestionEnd:** `*string` — Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetTeamMetrics() -> *v505.MetricsResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET metrics for team.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetTeamMetrics(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetConnectedProviders(UserId) -> map[string][]*v505.ClientFacingProviderWithStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET Users connected providers
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetConnectedProviders(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetLatestUserInfo(UserId) -> *v505.UserInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetLatestUserInfo(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.CreateInsurance(UserId, request) -> *v505.ClientFacingInsurance</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreateInsuranceRequest{
        PayorCode: "payor_code",
        MemberId: "member_id",
        Relationship: v505.ResponsibleRelationshipSelf,
        Insured: &v505.VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails{
            FirstName: "first_name",
            LastName: "last_name",
            Gender: v505.GenderFemale,
            Address: &v505.Address{
                FirstLine: "first_line",
                Country: "country",
                Zip: "zip",
                City: "city",
                State: "state",
            },
            Dob: "dob",
            Email: "email",
            PhoneNumber: "phone_number",
        },
    }
client.User.CreateInsurance(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**payorCode:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**groupId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**relationship:** `*v505.ResponsibleRelationship` 
    
</dd>
</dl>

<dl>
<dd>

**insured:** `*v505.VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails` 
    
</dd>
</dl>

<dl>
<dd>

**guarantor:** `*v505.GuarantorDetails` 
    
</dd>
</dl>

<dl>
<dd>

**isPrimary:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetLatestInsurance(UserId) -> *v505.ClientFacingInsurance</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserGetLatestInsuranceRequest{
        IsPrimary: v505.Bool(
            true,
        ),
    }
client.User.GetLatestInsurance(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isPrimary:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.UpsertUserInfo(UserId, request) -> *v505.UserInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserInfoCreateRequest{
        FirstName: "first_name",
        LastName: "last_name",
        Email: "email",
        PhoneNumber: "phone_number",
        Gender: "gender",
        Dob: "dob",
        Address: &v505.UserAddress{
            FirstLine: "first_line",
            Country: "country",
            Zip: "zip",
            City: "city",
            State: "state",
        },
    }
client.User.UpsertUserInfo(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**gender:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**dob:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*v505.UserAddress` 
    
</dd>
</dl>

<dl>
<dd>

**medicalProxy:** `*v505.GuarantorDetails` 
    
</dd>
</dl>

<dl>
<dd>

**race:** `*v505.Race` 
    
</dd>
</dl>

<dl>
<dd>

**ethnicity:** `*v505.Ethnicity` 
    
</dd>
</dl>

<dl>
<dd>

**sexualOrientation:** `*v505.SexualOrientation` 
    
</dd>
</dl>

<dl>
<dd>

**genderIdentity:** `*v505.GenderIdentity` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetByClientUserId(ClientUserId) -> *v505.ClientFacingUser</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET user_id from client_user_id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetByClientUserId(
        context.TODO(),
        "client_user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientUserId:** `string` — A unique ID representing the end user. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.DeregisterProvider(UserId, Provider) -> *v505.UserSuccessResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.DeregisterProvider(
        context.TODO(),
        "user_id",
        v505.ProvidersOura.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` — Provider slug. e.g., `oura`, `fitbit`, `garmin`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Get(UserId) -> *v505.ClientFacingUser</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.Get(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Delete(UserId) -> *v505.UserSuccessResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.Delete(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Patch(UserId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserPatchBody{}
client.User.Patch(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fallbackTimeZone:** `*string` 


    Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
    Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
    
    
</dd>
</dl>

<dl>
<dd>

**fallbackBirthDate:** `*string` — Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
    
</dd>
</dl>

<dl>
<dd>

**ingestionStart:** `*string` — Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>

<dl>
<dd>

**ingestionEnd:** `*string` — Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>

<dl>
<dd>

**clientUserId:** `*string` — A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.UndoDelete() -> *v505.UserSuccessResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserUndoDeleteRequest{
        UserId: v505.String(
            "user_id",
        ),
        ClientUserId: v505.String(
            "client_user_id",
        ),
    }
client.User.UndoDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `*string` — User ID to undo deletion. Mutually exclusive with `client_user_id`.
    
</dd>
</dl>

<dl>
<dd>

**clientUserId:** `*string` — Client User ID to undo deletion. Mutually exclusive with `user_id`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Refresh(UserId) -> *v505.UserRefreshSuccessResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Trigger a manual refresh for a specific user
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UserRefreshRequest{
        Timeout: v505.Float64(
            1.1,
        ),
    }
client.User.Refresh(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetDevices(UserId) -> []*v505.ClientFacingDevice</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetDevices(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetDevice(UserId, DeviceId) -> *v505.ClientFacingDevice</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetDevice(
        context.TODO(),
        "user_id",
        "device_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**deviceId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetUserSignInToken(UserId) -> *v505.UserSignInTokenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetUserSignInToken(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.CreatePortalUrl(UserId, request) -> *v505.CreateUserPortalUrlResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreateUserPortalUrlBody{
        Context: v505.CreateUserPortalUrlBodyContextLaunch,
    }
client.User.CreatePortalUrl(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**context:** `*v505.CreateUserPortalUrlBodyContext` 

`launch`: Generates a short-lived (minutes) portal URL that is intended for launching a user from your
authenticated web context directly into the Junction User Portal. This URL is not suitable for asynchronous
communications due to its verbosity as well as short-lived nature.

`communications`: Generates a long-lived (weeks) but shortened portal URL that is suitable for Emails, SMS
messages and other communication channels. Users may be asked to verify their identity with Email and SMS
authentication, e.g., when they open a short link on a new device. ℹ️ This enum is non-exhaustive.
    
</dd>
</dl>

<dl>
<dd>

**orderId:** `*string` — If specified, the generated URL will deeplink to the specified Order.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Team
<details><summary><code>client.Team.GetLinkConfig() -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Post teams.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.TeamGetLinkConfigRequest{}
client.Team.GetLinkConfig(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vitalLinkToken:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.Get(TeamId) -> *v505.ClientFacingTeam</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get team.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Team.Get(
        context.TODO(),
        "team_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetUserById() -> []*v505.ClientFacingUser</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search team users by user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.TeamGetUserByIdRequest{
        QueryId: v505.String(
            "query_id",
        ),
    }
client.Team.GetUserById(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**queryId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetSvixUrl() -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Team.GetSvixUrl(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetSourcePriorities() -> []map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET source priorities.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.TeamGetSourcePrioritiesRequest{
        DataType: v505.PriorityResourceWorkouts.Ptr(),
    }
client.Team.GetSourcePriorities(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**dataType:** `*v505.PriorityResource` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.UpdateSourcePriorities() -> []map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patch source priorities.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Team.UpdateSourcePriorities(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetPhysicians(TeamId) -> []*v505.ClientFacingPhysician</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Team.GetPhysicians(
        context.TODO(),
        "team_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Providers
<details><summary><code>client.Providers.GetAll() -> []*v505.ClientFacingProviderDetailed</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get Provider list
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ProvidersGetAllRequest{
        SourceType: v505.String(
            "source_type",
        ),
    }
client.Providers.GetAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sourceType:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Introspect
<details><summary><code>client.Introspect.GetUserResources() -> *v505.UserResourcesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.IntrospectGetUserResourcesRequest{
        UserId: v505.String(
            "user_id",
        ),
        Provider: v505.ProvidersOura.Ptr(),
        UserLimit: v505.Int(
            1,
        ),
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
    }
client.Introspect.GetUserResources(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `*string` — Filter by user ID.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**userLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Introspect.GetUserHistoricalPulls() -> *v505.UserHistoricalPullsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.IntrospectGetUserHistoricalPullsRequest{
        UserId: v505.String(
            "user_id",
        ),
        Provider: v505.ProvidersOura.Ptr(),
        UserLimit: v505.Int(
            1,
        ),
        Cursor: v505.String(
            "cursor",
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
    }
client.Introspect.GetUserHistoricalPulls(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `*string` — Filter by user ID.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**userLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` — The cursor for fetching the next page, or `null` to fetch the first page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LabTests
<details><summary><code>client.LabTests.Get() -> []*v505.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET all the lab tests the team has access to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetRequest{
        GenerationMethod: v505.LabTestGenerationMethodFilterAuto.Ptr(),
        LabSlug: v505.String(
            "lab_slug",
        ),
        CollectionMethod: v505.LabTestCollectionMethodTestkit.Ptr(),
        Status: v505.LabTestStatusActive.Ptr(),
        Name: v505.String(
            "name",
        ),
        OrderKey: v505.LabTestsGetRequestOrderKeyPrice.Ptr(),
        OrderDirection: v505.LabTestsGetRequestOrderDirectionAsc.Ptr(),
    }
client.LabTests.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**generationMethod:** `*v505.LabTestGenerationMethodFilter` — Filter on whether auto-generated lab tests created by Vital, manually created lab tests, or all lab tests should be returned.
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*string` — Filter by the slug of the lab for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*v505.LabTestCollectionMethod` — Filter by the collection method for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v505.LabTestStatus` — Filter by the status of these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**markerIds:** `*int` — Filter to only include lab tests containing these marker IDs.
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `*string` — Filter to only include lab tests containing these provider IDs.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter by the name of the lab test (a case-insensitive substring search).
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*v505.LabTestsGetRequestOrderKey` 
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*v505.LabTestsGetRequestOrderDirection` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.Create(request) -> *v505.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreateLabTestRequest{
        Name: "name",
        Method: v505.LabTestCollectionMethodTestkit,
        Description: "description",
    }
client.LabTests.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**markerIds:** `[]int` 
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**method:** `*v505.LabTestCollectionMethod` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fasting:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*v505.Labs` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetById(LabTestId) -> *v505.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET all the lab tests the team has access to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetByIdRequest{
        LabAccountId: v505.String(
            "lab_account_id",
        ),
    }
client.LabTests.GetById(
        context.TODO(),
        "lab_test_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labTestId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — The lab account ID. This lab account is used to determine the availability of markers and lab tests.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.UpdateLabTest(LabTestId, request) -> *v505.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.UpdateLabTestRequest{}
client.LabTests.UpdateLabTest(
        context.TODO(),
        "lab_test_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labTestId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetMarkers() -> *v505.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List active and orderable markers for a given Lab. Note that reflex markers are not included.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetMarkersRequest{
        LabSlug: v505.String(
            "lab_slug",
        ),
        Name: v505.String(
            "name",
        ),
        ALaCarteEnabled: v505.Bool(
            true,
        ),
        LabAccountId: v505.String(
            "lab_account_id",
        ),
        Page: v505.Int(
            1,
        ),
        Size: v505.Int(
            1,
        ),
    }
client.LabTests.GetMarkers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labId:** `*int` — The identifier Vital assigned to a lab partner.
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*string` — The slug of the lab for these markers. If both lab_id and lab_slug are provided, lab_slug will be used.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name or test code of an individual biomarker or a panel.
    
</dd>
</dl>

<dl>
<dd>

**aLaCarteEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — The lab account ID. This lab account is used to determine the availability of markers and lab tests.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetMarkersForOrderSet(request) -> *v505.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetMarkersForOrderSetRequest{
        Page: v505.Int(
            1,
        ),
        Size: v505.Int(
            1,
        ),
        Body: &v505.OrderSetRequest{},
    }
client.LabTests.GetMarkersForOrderSet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.OrderSetRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetMarkersForLabTest(LabTestId) -> *v505.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all markers for a given Lab Test, as well as any associated reflex markers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetMarkersForLabTestRequest{
        LabAccountId: v505.String(
            "lab_account_id",
        ),
        Page: v505.Int(
            1,
        ),
        Size: v505.Int(
            1,
        ),
    }
client.LabTests.GetMarkersForLabTest(
        context.TODO(),
        "lab_test_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labTestId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — The lab account ID. This lab account is used to determine the availability of markers and lab tests.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetMarkersByLabAndProviderId(ProviderId, LabId) -> *v505.ClientFacingMarker</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET a specific marker for the given lab and provider_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetMarkersByLabAndProviderIdRequest{
        LabAccountId: v505.String(
            "lab_account_id",
        ),
    }
client.LabTests.GetMarkersByLabAndProviderId(
        context.TODO(),
        "provider_id",
        1,
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**providerId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labId:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — The lab account ID. This lab account is used to determine the availability of markers and lab tests.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetLabs() -> []*v505.ClientFacingLab</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET all the labs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetLabs(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPaginated() -> *v505.LabTestResourcesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET lab tests the team has access to as a paginated list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetPaginatedRequest{
        LabTestLimit: v505.Int(
            1,
        ),
        NextCursor: v505.String(
            "next_cursor",
        ),
        GenerationMethod: v505.LabTestGenerationMethodFilterAuto.Ptr(),
        LabSlug: v505.String(
            "lab_slug",
        ),
        CollectionMethod: v505.LabTestCollectionMethodTestkit.Ptr(),
        Status: v505.LabTestStatusActive.Ptr(),
        Name: v505.String(
            "name",
        ),
        OrderKey: v505.LabTestsGetPaginatedRequestOrderKeyPrice.Ptr(),
        OrderDirection: v505.LabTestsGetPaginatedRequestOrderDirectionAsc.Ptr(),
    }
client.LabTests.GetPaginated(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labTestLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**generationMethod:** `*v505.LabTestGenerationMethodFilter` — Filter on whether auto-generated lab tests created by Vital, manually created lab tests, or all lab tests should be returned.
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*string` — Filter by the slug of the lab for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*v505.LabTestCollectionMethod` — Filter by the collection method for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v505.LabTestStatus` — Filter by the status of these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**markerIds:** `*int` — Filter to only include lab tests containing these marker IDs.
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `*string` — Filter to only include lab tests containing these provider IDs.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter by the name of the lab test (a case-insensitive substring search).
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*v505.LabTestsGetPaginatedRequestOrderKey` 
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*v505.LabTestsGetPaginatedRequestOrderDirection` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetLabTestCollectionInstructionPdf(LabTestId) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetLabTestCollectionInstructionPdf(
        context.TODO(),
        "lab_test_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labTestId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrders() -> *v505.GetOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET many orders with filters.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetOrdersRequest{
        SearchInput: v505.String(
            "search_input",
        ),
        StartDate: v505.Time(
            v505.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndDate: v505.Time(
            v505.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        UpdatedStartDate: v505.Time(
            v505.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        UpdatedEndDate: v505.Time(
            v505.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OrderKey: v505.LabTestsGetOrdersRequestOrderKeyCreatedAt.Ptr(),
        OrderDirection: v505.LabTestsGetOrdersRequestOrderDirectionAsc.Ptr(),
        IsCritical: v505.Bool(
            true,
        ),
        Interpretation: v505.InterpretationNormal.Ptr(),
        UserId: v505.String(
            "user_id",
        ),
        PatientName: v505.String(
            "patient_name",
        ),
        ShippingRecipientName: v505.String(
            "shipping_recipient_name",
        ),
        OrderTransactionId: v505.String(
            "order_transaction_id",
        ),
        Page: v505.Int(
            1,
        ),
        Size: v505.Int(
            1,
        ),
    }
client.LabTests.GetOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**searchInput:** `*string` — Search by order id, user id, patient name, shipping dob, or shipping recipient name.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
    
</dd>
</dl>

<dl>
<dd>

**updatedStartDate:** `*time.Time` — Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**updatedEndDate:** `*time.Time` — Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v505.OrderLowLevelStatus` — Filter by low level status.
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*v505.LabTestsGetOrdersRequestOrderKey` — Order key to sort by.
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*v505.LabTestsGetOrdersRequestOrderDirection` — Order direction to sort by.
    
</dd>
</dl>

<dl>
<dd>

**orderType:** `*v505.LabTestCollectionMethod` — Filter by method used to perform the lab test.
    
</dd>
</dl>

<dl>
<dd>

**isCritical:** `*bool` — Filter by critical order status.
    
</dd>
</dl>

<dl>
<dd>

**interpretation:** `*v505.Interpretation` — Filter by result interpretation of the lab test.
    
</dd>
</dl>

<dl>
<dd>

**orderActivationTypes:** `*v505.OrderActivationType` — Filter by activation type.
    
</dd>
</dl>

<dl>
<dd>

**userId:** `*string` — Filter by user ID.
    
</dd>
</dl>

<dl>
<dd>

**patientName:** `*string` — Filter by patient name.
    
</dd>
</dl>

<dl>
<dd>

**shippingRecipientName:** `*string` — Filter by shipping recipient name.
    
</dd>
</dl>

<dl>
<dd>

**orderIds:** `*string` — Filter by order ids.
    
</dd>
</dl>

<dl>
<dd>

**orderTransactionId:** `*string` — Filter by order transaction ID
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPhlebotomyAppointmentAvailability(request) -> *v505.AppointmentAvailabilitySlots</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the available time slots to book an appointment with a phlebotomist
for the given address and order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetPhlebotomyAppointmentAvailabilityRequest{
        StartDate: v505.String(
            "start_date",
        ),
        Body: &v505.UsAddress{
            FirstLine: "first_line",
            City: "city",
            State: "state",
            ZipCode: "zip_code",
        },
    }
client.LabTests.GetPhlebotomyAppointmentAvailability(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**startDate:** `*string` — Start date for appointment availability
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.UsAddress` — At-home phlebotomy appointment address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.BookPhlebotomyAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Book an at-home phlebotomy appointment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AppointmentBookingRequest{
        BookingKey: "booking_key",
    }
client.LabTests.BookPhlebotomyAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.AppointmentBookingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.RequestPhlebotomyAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Request an at-home phlebotomy appointment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.RequestAppointmentRequest{
        Address: &v505.UsAddress{
            FirstLine: "first_line",
            City: "city",
            State: "state",
            ZipCode: "zip_code",
        },
        Provider: v505.AppointmentProviderGetlabs,
    }
client.LabTests.RequestPhlebotomyAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*v505.UsAddress` — At-home phlebotomy appointment address.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.AppointmentProvider` 
    
</dd>
</dl>

<dl>
<dd>

**appointmentNotes:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ReschedulePhlebotomyAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reschedule a previously booked at-home phlebotomy appointment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AppointmentRescheduleRequest{
        BookingKey: "booking_key",
    }
client.LabTests.ReschedulePhlebotomyAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.AppointmentRescheduleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CancelPhlebotomyAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel a previously booked at-home phlebotomy appointment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ApiApiV1EndpointsVitalApiLabTestingOrdersHelpersAppointmentCancelRequest{
        CancellationReasonId: "cancellation_reason_id",
    }
client.LabTests.CancelPhlebotomyAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**cancellationReasonId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPhlebotomyAppointmentCancellationReason() -> []*v505.ClientFacingAppointmentCancellationReason</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the list of reasons for cancelling an at-home phlebotomy appointment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetPhlebotomyAppointmentCancellationReason(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPhlebotomyAppointment(OrderId) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the appointment associated with an order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetPhlebotomyAppointment(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetAreaInfo() -> *v505.AreaInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET information about an area with respect to lab-testing.

Information returned:
* Whether a given zip code is served by our Phlebotomy network.
* List of Lab locations in the area.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetAreaInfoRequest{
        ZipCode: "zip_code",
        Radius: v505.AllowedRadiusTen.Ptr(),
        Lab: v505.ClientFacingLabsQuest.Ptr(),
        LabAccountId: v505.String(
            "lab_account_id",
        ),
    }
client.LabTests.GetAreaInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**zipCode:** `string` — Zip code of the area to check
    
</dd>
</dl>

<dl>
<dd>

**radius:** `*v505.AllowedRadius` — Radius in which to search in miles
    
</dd>
</dl>

<dl>
<dd>

**lab:** `*v505.ClientFacingLabs` — Lab to check for PSCs
    
</dd>
</dl>

<dl>
<dd>

**labs:** `*v505.ClientFacingLabs` — List of labs to check for PSCs
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — Lab Account ID to use for availability checks
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPscInfo() -> *v505.PscInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetPscInfoRequest{
        ZipCode: "zip_code",
        LabId: 1,
        Radius: v505.AllowedRadiusTen.Ptr(),
        LabAccountId: v505.String(
            "lab_account_id",
        ),
    }
client.LabTests.GetPscInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**zipCode:** `string` — Zip code of the area to check
    
</dd>
</dl>

<dl>
<dd>

**labId:** `int` — Lab ID to check for PSCs
    
</dd>
</dl>

<dl>
<dd>

**radius:** `*v505.AllowedRadius` — Radius in which to search in miles. Note that we limit to 30 PSCs.
    
</dd>
</dl>

<dl>
<dd>

**capabilities:** `*v505.LabLocationCapability` — Filter for only locations with certain capabilities
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` — Lab Account ID to use for availability checks
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrderPscInfo(OrderId) -> *v505.PscInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetOrderPscInfoRequest{
        Radius: v505.AllowedRadiusTen.Ptr(),
    }
client.LabTests.GetOrderPscInfo(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**radius:** `*v505.AllowedRadius` — Radius in which to search in miles
    
</dd>
</dl>

<dl>
<dd>

**capabilities:** `*v505.LabLocationCapability` — Filter for only locations with certain capabilities
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetResultPdf(OrderId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint returns the lab results for the order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetResultPdf(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetResultMetadata(OrderId) -> *v505.LabResultsMetadata</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return metadata related to order results, such as lab metadata,
provider and sample dates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetResultMetadata(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetResultRaw(OrderId) -> *v505.LabResultsRaw</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return both metadata and raw json test data
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetResultRaw(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetLabelsPdf(OrderId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint returns the printed labels for the order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetLabelsPdfRequest{
        CollectionDate: v505.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
    }
client.LabTests.GetLabelsPdf(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**numberOfLabels:** `*int` — Number of labels to generate
    
</dd>
</dl>

<dl>
<dd>

**collectionDate:** `time.Time` — Collection date
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPscAppointmentAvailability() -> *v505.AppointmentAvailabilitySlots</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsGetPscAppointmentAvailabilityRequest{
        Lab: v505.AppointmentPscLabsQuest,
        StartDate: v505.String(
            "start_date",
        ),
        ZipCode: v505.String(
            "zip_code",
        ),
        Radius: v505.AllowedRadiusTen.Ptr(),
        AllowStale: v505.Bool(
            true,
        ),
    }
client.LabTests.GetPscAppointmentAvailability(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lab:** `*v505.AppointmentPscLabs` — Lab to check for availability
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — Start date for appointment availability
    
</dd>
</dl>

<dl>
<dd>

**siteCodes:** `*string` — List of site codes to fetch availability for
    
</dd>
</dl>

<dl>
<dd>

**zipCode:** `*string` — Zip code of the area to check
    
</dd>
</dl>

<dl>
<dd>

**radius:** `*v505.AllowedRadius` — Radius in which to search. (meters)
    
</dd>
</dl>

<dl>
<dd>

**allowStale:** `*bool` — If true, allows cached availability data to be returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.BookPscAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsBookPscAppointmentRequest{
        Body: &v505.AppointmentBookingRequest{
            BookingKey: "booking_key",
        },
    }
client.LabTests.BookPscAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — [!] This feature (Idempotency Key) is under closed beta. Idempotency Key support for booking PSC appointment.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyError:** `*string` — If `no-cache`, applies idempotency only to successful outcomes.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.AppointmentBookingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ReschedulePscAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AppointmentRescheduleRequest{
        BookingKey: "booking_key",
    }
client.LabTests.ReschedulePscAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.AppointmentRescheduleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CancelPscAppointment(OrderId, request) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.VitalCoreClientsLabTestGetlabsSchemaAppointmentCancelRequest{
        CancellationReasonId: "cancellationReasonId",
    }
client.LabTests.CancelPscAppointment(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>

<dl>
<dd>

**cancellationReasonId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPscAppointmentCancellationReason() -> []*v505.ClientFacingAppointmentCancellationReason</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetPscAppointmentCancellationReason(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetPscAppointment(OrderId) -> *v505.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the appointment associated with an order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetPscAppointment(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrderCollectionInstructionPdf(OrderId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET collection instructions for an order
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetOrderCollectionInstructionPdf(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrderRequistionPdf(OrderId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET requisition pdf for an order
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetOrderRequistionPdf(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrderAbnPdf(OrderId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET ABN pdf for an order
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetOrderAbnPdf(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetOrder(OrderId) -> *v505.ClientFacingOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET individual order by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.GetOrder(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CreateOrder(request) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreateOrderRequestCompatible{
        UserId: "user_id",
        PatientDetails: &v505.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: "dob",
            Gender: v505.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &v505.PatientAddressWithValidation{
            FirstLine: "first_line",
            City: "city",
            State: "state",
            Zip: "zip",
            Country: "country",
        },
    }
client.LabTests.CreateOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyError:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labTestId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderSet:** `*v505.OrderSetRequest` 
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*v505.LabTestCollectionMethod` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*v505.PhysicianCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**healthInsurance:** `*v505.HealthInsuranceCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*bool` — Defines whether order is priority or not. For some labs, this refers to a STAT order.
    
</dd>
</dl>

<dl>
<dd>

**billingType:** `*v505.Billing` 
    
</dd>
</dl>

<dl>
<dd>

**icdCodes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**consents:** `[]*v505.Consent` 
    
</dd>
</dl>

<dl>
<dd>

**activateBy:** `*string` — Schedule an Order to be processed in a future date.
    
</dd>
</dl>

<dl>
<dd>

**aoeAnswers:** `[]*v505.AoEAnswer` 
    
</dd>
</dl>

<dl>
<dd>

**passthrough:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clinicalNotes:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**creatorMemberId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patientDetails:** `*v505.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*v505.PatientAddressWithValidation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ImportOrder(request) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ImportOrderBody{
        UserId: "user_id",
        BillingType: v505.BillingClientBill,
        OrderSet: &v505.OrderSetRequest{},
        CollectionMethod: v505.LabTestCollectionMethodTestkit,
        PatientDetails: &v505.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: "dob",
            Gender: v505.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &v505.PatientAddress{
            ReceiverName: "receiver_name",
            FirstLine: "first_line",
            City: "city",
            State: "state",
            Zip: "zip",
            Country: "country",
        },
        SampleId: "sample_id",
    }
client.LabTests.ImportOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**billingType:** `*v505.Billing` 
    
</dd>
</dl>

<dl>
<dd>

**orderSet:** `*v505.OrderSetRequest` 
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*v505.LabTestCollectionMethod` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*v505.PhysicianCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**patientDetails:** `*v505.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*v505.PatientAddress` 
    
</dd>
</dl>

<dl>
<dd>

**sampleId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CancelOrder(OrderId) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

POST cancel order
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.CancelOrder(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.SimulateOrderProcess(OrderId, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get available test kits.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabTestsSimulateOrderProcessRequest{
        FinalStatus: v505.OrderStatusReceivedWalkInTestOrdered.Ptr(),
        Delay: v505.Int(
            1,
        ),
        Body: &v505.SimulationFlags{},
    }
client.LabTests.SimulateOrderProcess(
        context.TODO(),
        "order_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**finalStatus:** `*v505.OrderStatus` 
    
</dd>
</dl>

<dl>
<dd>

**delay:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v505.SimulationFlags` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.UpdateOnSiteCollectionOrderDrawCompleted(OrderId) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

PATCH update on site collection order when draw is completed
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabTests.UpdateOnSiteCollectionOrderDrawCompleted(
        context.TODO(),
        "order_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderId:** `string` — Your Order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ValidateIcdCodes(request) -> *v505.ValidateIcdCodesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ValidateIcdCodesBody{
        Codes: []string{
            "codes",
        },
    }
client.LabTests.ValidateIcdCodes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**codes:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Compendium
<details><summary><code>client.Compendium.Search(request) -> *v505.SearchCompendiumResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchCompendiumBody{
        TeamId: v505.CompendiumSearchRequestTeamIdInferFromContext.Ptr(),
        Mode: v505.SearchModeCanonical,
    }
client.Compendium.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.CompendiumSearchRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*v505.SearchMode` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cptCodes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**loincSetHash:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**labs:** `[]*v505.CompendiumSearchLabs` 
    
</dd>
</dl>

<dl>
<dd>

**includeRelated:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Compendium.Convert(request) -> *v505.ConvertCompendiumResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ConvertCompendiumBody{
        TeamId: v505.CompendiumConvertRequestTeamIdInferFromContext.Ptr(),
        TargetLab: v505.CompendiumSearchLabsLabcorp,
    }
client.Compendium.Convert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamId:** `*v505.CompendiumConvertRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**labTestId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**targetLab:** `*v505.CompendiumSearchLabs` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LabAccount
<details><summary><code>client.LabAccount.GetTeamLabAccounts() -> *v505.GetTeamLabAccountsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LabAccountGetTeamLabAccountsRequest{
        LabAccountId: v505.String(
            "lab_account_id",
        ),
        Status: v505.LabAccountStatusActive.Ptr(),
    }
client.LabAccount.GetTeamLabAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**labAccountId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v505.LabAccountStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## OrderTransaction
<details><summary><code>client.OrderTransaction.GetTransaction(TransactionId) -> *v505.GetOrderTransactionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrderTransaction.GetTransaction(
        context.TODO(),
        "transaction_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrderTransaction.GetTransactionResult(TransactionId) -> *v505.LabResultsRaw</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrderTransaction.GetTransactionResult(
        context.TODO(),
        "transaction_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrderTransaction.GetTransactionResultPdf(TransactionId) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrderTransaction.GetTransactionResultPdf(
        context.TODO(),
        "transaction_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Testkit
<details><summary><code>client.Testkit.Register(request) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.RegisterTestkitRequest{
        SampleId: "sample_id",
        PatientDetails: &v505.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: "dob",
            Gender: v505.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &v505.PatientAddressWithValidation{
            FirstLine: "first_line",
            City: "city",
            State: "state",
            Zip: "zip",
            Country: "country",
        },
    }
client.Testkit.Register(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `*string` — The user ID of the patient.
    
</dd>
</dl>

<dl>
<dd>

**sampleId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientDetails:** `*v505.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*v505.PatientAddressWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*v505.PhysicianCreateRequestBase` 
    
</dd>
</dl>

<dl>
<dd>

**healthInsurance:** `*v505.HealthInsuranceCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**consents:** `[]*v505.Consent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Testkit.CreateOrder(request) -> *v505.PostOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an order for an unregistered testkit
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreateRegistrableTestkitOrderRequest{
        UserId: "user_id",
        LabTestId: "lab_test_id",
        ShippingDetails: &v505.ShippingAddressWithValidation{
            ReceiverName: "receiver_name",
            FirstLine: "first_line",
            City: "city",
            State: "state",
            Zip: "zip",
            Country: "country",
            PhoneNumber: "phone_number",
        },
    }
client.Testkit.CreateOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**labTestId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**shippingDetails:** `*v505.ShippingAddressWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**passthrough:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**labAccountId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Order
<details><summary><code>client.Order.ResendEvents(request) -> *v505.ResendWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replay a webhook for a given set of orders
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.ResendWebhookBody{}
client.Order.ResendEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**startAt:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endAt:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Insurance
<details><summary><code>client.Insurance.SearchGetPayorInfo() -> []*v505.ClientFacingPayorSearchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.InsuranceSearchGetPayorInfoRequest{
        InsuranceName: v505.String(
            "insurance_name",
        ),
        Provider: v505.PayorCodeExternalProviderChangeHealthcare.Ptr(),
        ProviderPayorId: v505.String(
            "provider_payor_id",
        ),
    }
client.Insurance.SearchGetPayorInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**insuranceName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.PayorCodeExternalProvider` 
    
</dd>
</dl>

<dl>
<dd>

**providerPayorId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Insurance.SearchPayorInfo(request) -> []*v505.ClientFacingPayorSearchResponseDeprecated</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.PayorSearchRequest{}
client.Insurance.SearchPayorInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**insuranceName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.PayorCodeExternalProvider` 
    
</dd>
</dl>

<dl>
<dd>

**providerId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Insurance.SearchDiagnosis() -> []*v505.ClientFacingDiagnosisInformation</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.InsuranceSearchDiagnosisRequest{
        DiagnosisQuery: "diagnosis_query",
    }
client.Insurance.SearchDiagnosis(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**diagnosisQuery:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payor
<details><summary><code>client.Payor.CreatePayor(request) -> *v505.ClientFacingPayor</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.CreatePayorBody{
        Name: "name",
        Address: &v505.Address{
            FirstLine: "first_line",
            Country: "country",
            Zip: "zip",
            City: "city",
            State: "state",
        },
    }
client.Payor.CreatePayor(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*v505.Address` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*v505.PayorCodeExternalProvider` 
    
</dd>
</dl>

<dl>
<dd>

**providerPayorId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LabReport
<details><summary><code>client.LabReport.ParserCreateJob(request) -> *v505.ParsingJob</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a parse job, uploads the file(s) to provider, persists the job row,
and starts the ParseLabReport. Returns a generated job_id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BodyCreateLabReportParserJob{
        UserId: "user_id",
    }
client.LabReport.ParserCreateJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabReport.ParserGetJob(JobId) -> *v505.ParsingJob</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the parse job status and stored result if completed.

Returns:
    ParseLabResultJobResponse with job status and parsed data (if complete)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LabReport.ParserGetJob(
        context.TODO(),
        "job_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Aggregate
<details><summary><code>client.Aggregate.QueryOne(UserId, request) -> *v505.AggregationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.QueryBatch{
        Timeframe: &v505.QueryBatchTimeframe{
            RelativeTimeframe: &v505.RelativeTimeframe{
                Anchor: "anchor",
                Past: &v505.Period{
                    Unit: v505.PeriodUnitMinute,
                },
            },
        },
        Queries: []*v505.Query{
            &v505.Query{
                Select: []*v505.QuerySelectItem{
                    &v505.QuerySelectItem{
                        AggregateExpr: &v505.AggregateExpr{
                            Arg: &v505.AggregateExprArg{
                                SleepColumnExpr: &v505.SleepColumnExpr{
                                    Sleep: v505.SleepColumnExprSleepId,
                                },
                            },
                            Func: v505.AggregateExprFuncMean,
                        },
                    },
                },
            },
        },
    }
client.Aggregate.QueryOne(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**accept:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**timeframe:** `*v505.QueryBatchTimeframe` 
    
</dd>
</dl>

<dl>
<dd>

**queries:** `[]*v505.Query` 
    
</dd>
</dl>

<dl>
<dd>

**config:** `*v505.QueryConfig` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Aggregate.GetResultTableForContinuousQuery(UserId, QueryIdOrSlug) -> *v505.AggregationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Aggregate.GetResultTableForContinuousQuery(
        context.TODO(),
        "user_id",
        "query_id_or_slug",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**queryIdOrSlug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**accept:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Aggregate.GetTaskHistoryForContinuousQuery(UserId, QueryIdOrSlug) -> *v505.ContinuousQueryTaskHistoryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AggregateGetTaskHistoryForContinuousQueryRequest{
        NextCursor: v505.String(
            "next_cursor",
        ),
        Limit: v505.Int(
            1,
        ),
    }
client.Aggregate.GetTaskHistoryForContinuousQuery(
        context.TODO(),
        "user_id",
        "query_id_or_slug",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**queryIdOrSlug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nextCursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
