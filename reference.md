# Reference
## Link
<details><summary><code>client.Link.ListBulkOps() -> *vitalgo.BulkOpsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ListBulkOpsLinkRequest{}
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

**teamId:** `*vitalgo.ListBulkOpsLinkRequestTeamId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.BulkImport(request) -> *vitalgo.BulkImportConnectionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BulkImportConnectionsBody{
        Provider: vitalgo.OAuthProvidersOura,
        Connections: []*vitalgo.ConnectionRecipe{
            &vitalgo.ConnectionRecipe{
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

**teamId:** `*vitalgo.BulkImportLinkRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.OAuthProviders` 
    
</dd>
</dl>

<dl>
<dd>

**connections:** `[]*vitalgo.ConnectionRecipe` 
    
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
request := &vitalgo.BulkTriggerHistoricalPullBody{
        UserIds: []string{
            "user_ids",
        },
        Provider: vitalgo.OAuthProvidersOura,
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

**teamId:** `*vitalgo.BulkTriggerHistoricalPullLinkRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.OAuthProviders` 
    
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

<details><summary><code>client.Link.BulkExport(request) -> *vitalgo.BulkExportConnectionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BulkExportConnectionsBody{
        Provider: vitalgo.OAuthProvidersOura,
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

**teamId:** `*vitalgo.BulkExportLinkRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.OAuthProviders` 
    
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
request := &vitalgo.BulkPauseConnectionsBody{
        UserIds: []string{
            "user_ids",
        },
        Provider: vitalgo.OAuthProvidersOura,
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

**teamId:** `*vitalgo.BulkPauseLinkRequestTeamId` 
    
</dd>
</dl>

<dl>
<dd>

**userIds:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.OAuthProviders` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.Token(request) -> *vitalgo.LinkTokenExchangeResponse</code></summary>
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
request := &vitalgo.LinkTokenExchange{
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

**provider:** `*vitalgo.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**redirectUrl:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterOnProviders:** `[]*vitalgo.Providers` 

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
request := &vitalgo.LinkTokenValidationRequest{
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

<details><summary><code>client.Link.CodeCreate() -> *vitalgo.VitalTokenCreatedResponse</code></summary>
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
request := &vitalgo.CodeCreateLinkRequest{
        UserId: "user_id",
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
request := &vitalgo.BeginLinkTokenRequest{
        LinkToken: "link_token",
        Provider: vitalgo.ProvidersOura,
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

**provider:** `*vitalgo.Providers` 
    
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
request := &vitalgo.TokenStateLinkRequest{}
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
request := &vitalgo.EmailAuthLink{
        Email: "email",
        Provider: vitalgo.ProvidersOura,
        AuthType: vitalgo.AuthTypePassword,
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

**provider:** `*vitalgo.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**authType:** `*vitalgo.AuthType` 
    
</dd>
</dl>

<dl>
<dd>

**region:** `*vitalgo.Region` 
    
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
request := &vitalgo.PasswordAuthLink{
        Username: "username",
        Password: "password",
        Provider: vitalgo.ProvidersOura,
        AuthType: vitalgo.AuthTypePassword,
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

**provider:** `*vitalgo.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**authType:** `*vitalgo.AuthType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.GenerateOauthLink(OauthProvider) -> *vitalgo.Source</code></summary>
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
request := &vitalgo.GenerateOauthLinkLinkRequest{}
client.Link.GenerateOauthLink(
        context.TODO(),
        vitalgo.OAuthProvidersOura.Ptr(),
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

**oauthProvider:** `*vitalgo.OAuthProviders` 
    
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

<details><summary><code>client.Link.ConnectPasswordProvider(Provider, request) -> *vitalgo.ProviderLinkResponse</code></summary>
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
request := &vitalgo.IndividualProviderData{
        Username: "username",
        Password: "password",
    }
client.Link.ConnectPasswordProvider(
        context.TODO(),
        vitalgo.PasswordProvidersWhoop.Ptr(),
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

**provider:** `*vitalgo.PasswordProviders` 
    
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

**region:** `*vitalgo.Region` — Provider region to authenticate against. Only applicable to specific providers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.CompletePasswordProviderMfa(Provider, request) -> *vitalgo.ProviderLinkResponse</code></summary>
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
request := &vitalgo.CompletePasswordProviderMfaBody{
        MfaCode: "mfa_code",
    }
client.Link.CompletePasswordProviderMfa(
        context.TODO(),
        vitalgo.PasswordProvidersWhoop.Ptr(),
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

**provider:** `*vitalgo.PasswordProviders` 
    
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
request := &vitalgo.EmailProviderAuthLink{
        Email: "email",
    }
client.Link.ConnectEmailAuthProvider(
        context.TODO(),
        vitalgo.EmailProviders(
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

**provider:** `vitalgo.EmailProviders` 
    
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

**emailProviderAuthLinkProvider:** `*vitalgo.Providers` 
    
</dd>
</dl>

<dl>
<dd>

**region:** `*vitalgo.Region` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.GetAllProviders() -> []*vitalgo.SourceLink</code></summary>
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
request := &vitalgo.GetAllProvidersLinkRequest{}
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
request := &vitalgo.ManualConnectionData{
        UserId: "user_id",
    }
client.Link.ConnectManualProvider(
        context.TODO(),
        vitalgo.ManualProvidersBeurerBle.Ptr(),
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

**provider:** `*vitalgo.ManualProviders` 
    
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
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Link.ConnectDemoProvider(request) -> *vitalgo.DemoConnectionStatus</code></summary>
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
request := &vitalgo.DemoConnectionCreationPayload{
        UserId: "user_id",
        Provider: vitalgo.DemoProvidersAppleHealthKit,
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

**provider:** `*vitalgo.DemoProviders` — Demo provider. For more information, please check out our docs (https://docs.tryvital.io/wearables/providers/test_data)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Electrocardiogram
<details><summary><code>client.Electrocardiogram.Get(UserId) -> *vitalgo.ClientFacingElectrocardiogramResponse</code></summary>
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
request := &vitalgo.GetElectrocardiogramRequest{
        StartDate: vitalgo.MustParseDateTime(
            "2023-01-15",
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

**startDate:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 
    
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
<details><summary><code>client.SleepCycle.Get(UserId) -> *vitalgo.ClientSleepCycleResponse</code></summary>
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
request := &vitalgo.GetSleepCycleRequest{
        StartDate: vitalgo.MustParseDateTime(
            "2023-01-15",
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

**startDate:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 
    
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
<details><summary><code>client.Profile.Get(UserId) -> *vitalgo.ClientFacingProfile</code></summary>
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
request := &vitalgo.GetProfileRequest{}
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

<details><summary><code>client.Profile.GetRaw(UserId) -> *vitalgo.RawProfile</code></summary>
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
request := &vitalgo.GetRawProfileRequest{}
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
<details><summary><code>client.Devices.GetRaw(UserId) -> *vitalgo.RawDevices</code></summary>
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
request := &vitalgo.GetRawDevicesRequest{}
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
<details><summary><code>client.Activity.Get(UserId) -> *vitalgo.ClientActivityResponse</code></summary>
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
request := &vitalgo.GetActivityRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Activity.GetRaw(UserId) -> *vitalgo.RawActivity</code></summary>
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
request := &vitalgo.GetRawActivityRequest{
        StartDate: "start_date",
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
<details><summary><code>client.Workouts.Get(UserId) -> *vitalgo.ClientWorkoutResponse</code></summary>
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
request := &vitalgo.GetWorkoutsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Workouts.GetRaw(UserId) -> *vitalgo.RawWorkout</code></summary>
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
request := &vitalgo.GetRawWorkoutsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Workouts.GetByWorkoutId(WorkoutId) -> *vitalgo.ClientFacingStream</code></summary>
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
<details><summary><code>client.Sleep.Get(UserId) -> *vitalgo.ClientSleepResponse</code></summary>
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
request := &vitalgo.GetSleepRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Sleep.GetRaw(UserId) -> *vitalgo.RawSleep</code></summary>
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
request := &vitalgo.GetRawSleepRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Sleep.GetStreamBySleepId(SleepId) -> *vitalgo.ClientFacingSleepStream</code></summary>
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
<details><summary><code>client.Body.Get(UserId) -> *vitalgo.ClientBodyResponse</code></summary>
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
request := &vitalgo.GetBodyRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Body.GetRaw(UserId) -> *vitalgo.RawBody</code></summary>
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
request := &vitalgo.GetRawBodyRequest{
        StartDate: "start_date",
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
<details><summary><code>client.Meal.Get(UserId) -> *vitalgo.ClientFacingMealResponse</code></summary>
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
request := &vitalgo.GetMealRequest{
        StartDate: "start_date",
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
<details><summary><code>client.MenstrualCycle.Get(UserId) -> *vitalgo.MenstrualCycleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetMenstrualCycleRequest{
        StartDate: vitalgo.MustParseDateTime(
            "2023-01-15",
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

**startDate:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Vitals
<details><summary><code>client.Vitals.WorkoutSwimmingStrokeGrouped(UserId) -> *vitalgo.GroupedWorkoutSwimmingStrokeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WorkoutSwimmingStrokeGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.WorkoutDistanceGrouped(UserId) -> *vitalgo.GroupedWorkoutDistanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WorkoutDistanceGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HeartRateRecoveryOneMinuteGrouped(UserId) -> *vitalgo.GroupedHeartRateRecoveryOneMinuteResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HeartRateRecoveryOneMinuteGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.WaistCircumferenceGrouped(UserId) -> *vitalgo.GroupedWaistCircumferenceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WaistCircumferenceGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.LeanBodyMassGrouped(UserId) -> *vitalgo.GroupedLeanBodyMassResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.LeanBodyMassGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyMassIndexGrouped(UserId) -> *vitalgo.GroupedBodyMassIndexResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyMassIndexGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BasalBodyTemperatureGrouped(UserId) -> *vitalgo.GroupedBasalBodyTemperatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BasalBodyTemperatureGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HandwashingGrouped(UserId) -> *vitalgo.GroupedHandwashingResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HandwashingGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.DaylightExposureGrouped(UserId) -> *vitalgo.GroupedDaylightExposureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.DaylightExposureGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.UvExposureGrouped(UserId) -> *vitalgo.GroupedUvExposureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.UvExposureGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.FallGrouped(UserId) -> *vitalgo.GroupedFallResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.FallGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.InhalerUsageGrouped(UserId) -> *vitalgo.GroupedInhalerUsageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.InhalerUsageGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.PeakExpiratoryFlowRateGrouped(UserId) -> *vitalgo.GroupedPeakExpiratoryFlowRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.PeakExpiratoryFlowRateGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.ForcedVitalCapacityGrouped(UserId) -> *vitalgo.GroupedForcedVitalCapacityResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ForcedVitalCapacityGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.ForcedExpiratoryVolume1Grouped(UserId) -> *vitalgo.GroupedForcedExpiratoryVolume1Response</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ForcedExpiratoryVolume1GroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.WheelchairPushGrouped(UserId) -> *vitalgo.GroupedWheelchairPushResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WheelchairPushGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.SleepBreathingDisturbanceGrouped(UserId) -> *vitalgo.GroupedSleepBreathingDisturbanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.SleepBreathingDisturbanceGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.SleepApneaAlertGrouped(UserId) -> *vitalgo.GroupedSleepApneaAlertResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.SleepApneaAlertGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.StandDurationGrouped(UserId) -> *vitalgo.GroupedStandDurationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StandDurationGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.StandHourGrouped(UserId) -> *vitalgo.GroupedStandHourResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StandHourGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HeartRateAlertGrouped(UserId) -> *vitalgo.GroupedHeartRateAlertResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HeartRateAlertGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.AfibBurdenGrouped(UserId) -> *vitalgo.GroupedAFibBurdenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.AfibBurdenGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.WorkoutDurationGrouped(UserId) -> *vitalgo.GroupedWorkoutDurationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WorkoutDurationGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Vo2MaxGrouped(UserId) -> *vitalgo.GroupedVo2MaxResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.Vo2MaxGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.StressLevelGrouped(UserId) -> *vitalgo.GroupedStressLevelResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StressLevelGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.MindfulnessMinutesGrouped(UserId) -> *vitalgo.GroupedMindfulnessMinutesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.MindfulnessMinutesGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CaffeineGrouped(UserId) -> *vitalgo.GroupedCaffeineResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaffeineGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.WaterGrouped(UserId) -> *vitalgo.GroupedWaterResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WaterGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.StepsGrouped(UserId) -> *vitalgo.GroupedStepsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StepsGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.FloorsClimbedGrouped(UserId) -> *vitalgo.GroupedFloorsClimbedResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.FloorsClimbedGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.DistanceGrouped(UserId) -> *vitalgo.GroupedDistanceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.DistanceGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CaloriesBasalGrouped(UserId) -> *vitalgo.GroupedCaloriesBasalResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaloriesBasalGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CaloriesActiveGrouped(UserId) -> *vitalgo.GroupedCaloriesActiveResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaloriesActiveGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.RespiratoryRateGrouped(UserId) -> *vitalgo.GroupedRespiratoryRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.RespiratoryRateGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.NoteGrouped(UserId) -> *vitalgo.GroupedNoteResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.NoteGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.InsulinInjectionGrouped(UserId) -> *vitalgo.GroupedInsulinInjectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.InsulinInjectionGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.IgeGrouped(UserId) -> *vitalgo.GroupedIgeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.IgeGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.IggGrouped(UserId) -> *vitalgo.GroupedIggResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.IggGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HypnogramGrouped(UserId) -> *vitalgo.GroupedHypnogramResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HypnogramGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HrvGrouped(UserId) -> *vitalgo.GroupedHrvResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HrvGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.HeartrateGrouped(UserId) -> *vitalgo.GroupedHeartRateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HeartrateGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.GlucoseGrouped(UserId) -> *vitalgo.GroupedGlucoseResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GlucoseGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CholesterolGrouped(UserId) -> *vitalgo.GroupedCholesterolResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CarbohydratesGrouped(UserId) -> *vitalgo.GroupedCarbohydratesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CarbohydratesGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyTemperatureDeltaGrouped(UserId) -> *vitalgo.GroupedBodyTemperatureDeltaResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyTemperatureDeltaGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyTemperatureGrouped(UserId) -> *vitalgo.GroupedBodyTemperatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyTemperatureGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyWeightGrouped(UserId) -> *vitalgo.GroupedBodyWeightResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyWeightGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyFatGrouped(UserId) -> *vitalgo.GroupedBodyFatResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyFatGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BloodOxygenGrouped(UserId) -> *vitalgo.GroupedBloodOxygenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BloodOxygenGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.ElectrocardiogramVoltageGrouped(UserId) -> *vitalgo.GroupedElectrocardiogramVoltageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ElectrocardiogramVoltageGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BloodPressureGrouped(UserId) -> *vitalgo.GroupedBloodPressureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BloodPressureGroupedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Vo2Max(UserId) -> []*vitalgo.ClientFacingVo2MaxTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.Vo2MaxVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.StressLevel(UserId) -> []*vitalgo.ClientFacingStressLevelTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StressLevelVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.MindfulnessMinutes(UserId) -> []*vitalgo.ClientFacingMindfulnessMinutesTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.MindfulnessMinutesVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Caffeine(UserId) -> []*vitalgo.ClientFacingCaffeineTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaffeineVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Water(UserId) -> []*vitalgo.ClientFacingWaterTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.WaterVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Steps(UserId) -> []*vitalgo.ClientFacingStepsTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.StepsVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.FloorsClimbed(UserId) -> []*vitalgo.ClientFacingFloorsClimbedTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.FloorsClimbedVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Distance(UserId) -> []*vitalgo.ClientFacingDistanceTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.DistanceVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CaloriesBasal(UserId) -> []*vitalgo.ClientFacingCaloriesBasalTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaloriesBasalVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CaloriesActive(UserId) -> []*vitalgo.ClientFacingCaloriesActiveTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CaloriesActiveVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.RespiratoryRate(UserId) -> []*vitalgo.ClientFacingRespiratoryRateTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.RespiratoryRateVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Ige(UserId) -> []*vitalgo.ClientFacingIgeTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.IgeVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Igg(UserId) -> []*vitalgo.ClientFacingIggTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.IggVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Hypnogram(UserId) -> []*vitalgo.ClientFacingHypnogramTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HypnogramVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Hrv(UserId) -> []*vitalgo.ClientFacingHrvTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HrvVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Heartrate(UserId) -> []*vitalgo.ClientFacingHeartRateTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.HeartrateVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Glucose(UserId) -> []*vitalgo.ClientFacingGlucoseTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GlucoseVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CholesterolTriglycerides(UserId) -> []*vitalgo.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolTriglyceridesVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CholesterolTotal(UserId) -> []*vitalgo.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolTotalVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CholesterolLdl(UserId) -> []*vitalgo.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolLdlVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.CholesterolHdl(UserId) -> []*vitalgo.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolHdlVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.Cholesterol(UserId) -> []*vitalgo.ClientFacingCholesterolTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CholesterolVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyWeight(UserId) -> []*vitalgo.ClientFacingBodyWeightTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyWeightVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BodyFat(UserId) -> []*vitalgo.ClientFacingBodyFatTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BodyFatVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BloodOxygen(UserId) -> []*vitalgo.ClientFacingBloodOxygenTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BloodOxygenVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.ElectrocardiogramVoltage(UserId) -> []*vitalgo.ClientFacingElectrocardiogramVoltageTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ElectrocardiogramVoltageVitalsRequest{
        StartDate: "start_date",
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

<details><summary><code>client.Vitals.BloodPressure(UserId) -> []*vitalgo.ClientFacingBloodPressureTimeseries</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BloodPressureVitalsRequest{
        StartDate: "start_date",
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
<details><summary><code>client.User.GetAll() -> *vitalgo.PaginatedUsersResponse</code></summary>
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
request := &vitalgo.GetAllUserRequest{}
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

<details><summary><code>client.User.Create(request) -> *vitalgo.ClientFacingUserKey</code></summary>
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
request := &vitalgo.UserCreateBody{
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

**fallbackBirthDate:** `*time.Time` — Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
    
</dd>
</dl>

<dl>
<dd>

**ingestionStart:** `*time.Time` — Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>

<dl>
<dd>

**ingestionEnd:** `*time.Time` — Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetTeamMetrics() -> *vitalgo.MetricsResult</code></summary>
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

<details><summary><code>client.User.GetConnectedProviders(UserId) -> map[string][]*vitalgo.ClientFacingProviderWithStatus</code></summary>
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

<details><summary><code>client.User.GetLatestUserInfo(UserId) -> *vitalgo.UserInfo</code></summary>
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

<details><summary><code>client.User.CreateInsurance(UserId, request) -> *vitalgo.ClientFacingInsurance</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CreateInsuranceRequest{
        PayorCode: "payor_code",
        MemberId: "member_id",
        Relationship: vitalgo.ResponsibleRelationshipSelf,
        Insured: &vitalgo.VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails{
            FirstName: "first_name",
            LastName: "last_name",
            Gender: vitalgo.GenderFemale,
            Address: &vitalgo.Address{
                FirstLine: "first_line",
                Country: "country",
                Zip: "zip",
                City: "city",
                State: "state",
            },
            Dob: vitalgo.MustParseDateTime(
                "2023-01-15",
            ),
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

**relationship:** `*vitalgo.ResponsibleRelationship` 
    
</dd>
</dl>

<dl>
<dd>

**insured:** `*vitalgo.VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails` 
    
</dd>
</dl>

<dl>
<dd>

**guarantor:** `*vitalgo.GuarantorDetails` 
    
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

<details><summary><code>client.User.GetLatestInsurance(UserId) -> *vitalgo.ClientFacingInsurance</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetLatestInsuranceUserRequest{}
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

<details><summary><code>client.User.UpsertUserInfo(UserId, request) -> *vitalgo.UserInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.UserInfoCreateRequest{
        FirstName: "first_name",
        LastName: "last_name",
        Email: "email",
        PhoneNumber: "phone_number",
        Gender: "gender",
        Dob: vitalgo.MustParseDateTime(
            "2023-01-15",
        ),
        Address: &vitalgo.Address{
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

**dob:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*vitalgo.Address` 
    
</dd>
</dl>

<dl>
<dd>

**medicalProxy:** `*vitalgo.GuarantorDetails` 
    
</dd>
</dl>

<dl>
<dd>

**race:** `*vitalgo.Race` 
    
</dd>
</dl>

<dl>
<dd>

**ethnicity:** `*vitalgo.Ethnicity` 
    
</dd>
</dl>

<dl>
<dd>

**sexualOrientation:** `*vitalgo.SexualOrientation` 
    
</dd>
</dl>

<dl>
<dd>

**genderIdentity:** `*vitalgo.GenderIdentity` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetByClientUserId(ClientUserId) -> *vitalgo.ClientFacingUser</code></summary>
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

<details><summary><code>client.User.DeregisterProvider(UserId, Provider) -> *vitalgo.UserSuccessResponse</code></summary>
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
        vitalgo.ProvidersOura.Ptr(),
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

**provider:** `*vitalgo.Providers` — Provider slug. e.g., `oura`, `fitbit`, `garmin`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Get(UserId) -> *vitalgo.ClientFacingUser</code></summary>
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

<details><summary><code>client.User.Delete(UserId) -> *vitalgo.UserSuccessResponse</code></summary>
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
request := &vitalgo.UserPatchBody{}
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

**fallbackBirthDate:** `*time.Time` — Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
    
</dd>
</dl>

<dl>
<dd>

**ingestionStart:** `*time.Time` — Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
</dd>
</dl>

<dl>
<dd>

**ingestionEnd:** `*time.Time` — Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
    
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

<details><summary><code>client.User.UndoDelete() -> *vitalgo.UserSuccessResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.UndoDeleteUserRequest{}
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

<details><summary><code>client.User.Refresh(UserId) -> *vitalgo.UserRefreshSuccessResponse</code></summary>
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
request := &vitalgo.RefreshUserRequest{}
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

<details><summary><code>client.User.GetDevices(UserId) -> []*vitalgo.ClientFacingDevice</code></summary>
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

<details><summary><code>client.User.GetDevice(UserId, DeviceId) -> *vitalgo.ClientFacingDevice</code></summary>
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

<details><summary><code>client.User.GetUserSignInToken(UserId) -> *vitalgo.UserSignInTokenResponse</code></summary>
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

<details><summary><code>client.User.CreatePortalUrl(UserId, request) -> *vitalgo.CreateUserPortalUrlResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CreateUserPortalUrlBody{
        Context: vitalgo.CreateUserPortalUrlBodyContextLaunch,
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

**context:** `*vitalgo.CreateUserPortalUrlBodyContext` 

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
request := &vitalgo.GetLinkConfigTeamRequest{}
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

<details><summary><code>client.Team.Get(TeamId) -> *vitalgo.ClientFacingTeam</code></summary>
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

<details><summary><code>client.Team.GetUserById() -> []*vitalgo.ClientFacingUser</code></summary>
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
request := &vitalgo.GetUserByIdTeamRequest{}
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
request := &vitalgo.GetSourcePrioritiesTeamRequest{}
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

**dataType:** `*vitalgo.PriorityResource` 
    
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

<details><summary><code>client.Team.GetPhysicians(TeamId) -> []*vitalgo.ClientFacingPhysician</code></summary>
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
<details><summary><code>client.Providers.GetAll() -> []*vitalgo.ClientFacingProviderDetailed</code></summary>
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
request := &vitalgo.GetAllProvidersRequest{}
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
<details><summary><code>client.Introspect.GetUserResources() -> *vitalgo.UserResourcesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetUserResourcesIntrospectRequest{}
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

**provider:** `*vitalgo.Providers` 
    
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

<details><summary><code>client.Introspect.GetUserHistoricalPulls() -> *vitalgo.UserHistoricalPullsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetUserHistoricalPullsIntrospectRequest{}
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

**provider:** `*vitalgo.GetUserHistoricalPullsIntrospectRequestProvider` 
    
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
<details><summary><code>client.LabTests.Get() -> []*vitalgo.ClientFacingLabTest</code></summary>
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
request := &vitalgo.GetLabTestsRequest{}
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

**generationMethod:** `*vitalgo.LabTestGenerationMethodFilter` — Filter on whether auto-generated lab tests created by Vital, manually created lab tests, or all lab tests should be returned.
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*string` — Filter by the slug of the lab for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*vitalgo.LabTestCollectionMethod` — Filter by the collection method for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vitalgo.LabTestStatus` — Filter by the status of these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**markerIds:** `[]int` — Filter to only include lab tests containing these marker IDs.
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `[]string` — Filter to only include lab tests containing these provider IDs.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter by the name of the lab test (a case-insensitive substring search).
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*vitalgo.GetLabTestsRequestOrderKey` 
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*vitalgo.GetLabTestsRequestOrderDirection` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.Create(request) -> *vitalgo.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CreateLabTestRequest{
        Name: "name",
        Method: vitalgo.LabTestCollectionMethodTestkit,
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

**method:** `*vitalgo.LabTestCollectionMethod` 
    
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
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetById(LabTestId) -> *vitalgo.ClientFacingLabTest</code></summary>
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
request := &vitalgo.GetByIdLabTestsRequest{}
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

<details><summary><code>client.LabTests.UpdateLabTest(LabTestId, request) -> *vitalgo.ClientFacingLabTest</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.UpdateLabTestRequest{}
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

<details><summary><code>client.LabTests.GetMarkers() -> *vitalgo.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

GET all the markers for the given lab.
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
request := &vitalgo.GetMarkersLabTestsRequest{}
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

**labId:** `[]int` — The identifier Vital assigned to a lab partner.
    
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

<details><summary><code>client.LabTests.GetMarkersForOrderSet(request) -> *vitalgo.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetMarkersForOrderSetLabTestsRequest{
        Body: &vitalgo.OrderSetRequest{},
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

**request:** `*vitalgo.OrderSetRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.GetMarkersForLabTest(LabTestId) -> *vitalgo.GetMarkersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetMarkersForLabTestLabTestsRequest{}
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

<details><summary><code>client.LabTests.GetMarkersByLabAndProviderId(LabId, ProviderId) -> *vitalgo.ClientFacingMarker</code></summary>
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
request := &vitalgo.GetMarkersByLabAndProviderIdLabTestsRequest{}
client.LabTests.GetMarkersByLabAndProviderId(
        context.TODO(),
        1,
        "provider_id",
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

**labId:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**providerId:** `string` 
    
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

<details><summary><code>client.LabTests.GetLabs() -> []*vitalgo.ClientFacingLab</code></summary>
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

<details><summary><code>client.LabTests.GetPaginated() -> *vitalgo.LabTestResourcesResponse</code></summary>
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
request := &vitalgo.GetPaginatedLabTestsRequest{}
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

**generationMethod:** `*vitalgo.LabTestGenerationMethodFilter` — Filter on whether auto-generated lab tests created by Vital, manually created lab tests, or all lab tests should be returned.
    
</dd>
</dl>

<dl>
<dd>

**labSlug:** `*string` — Filter by the slug of the lab for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*vitalgo.LabTestCollectionMethod` — Filter by the collection method for these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vitalgo.LabTestStatus` — Filter by the status of these lab tests.
    
</dd>
</dl>

<dl>
<dd>

**markerIds:** `[]int` — Filter to only include lab tests containing these marker IDs.
    
</dd>
</dl>

<dl>
<dd>

**providerIds:** `[]string` — Filter to only include lab tests containing these provider IDs.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter by the name of the lab test (a case-insensitive substring search).
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*vitalgo.GetPaginatedLabTestsRequestOrderKey` 
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*vitalgo.GetPaginatedLabTestsRequestOrderDirection` 
    
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

<details><summary><code>client.LabTests.GetOrders() -> *vitalgo.GetOrdersResponse</code></summary>
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
request := &vitalgo.GetOrdersLabTestsRequest{}
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

**status:** `[]*vitalgo.OrderLowLevelStatus` — Filter by low level status.
    
</dd>
</dl>

<dl>
<dd>

**orderKey:** `*vitalgo.GetOrdersLabTestsRequestOrderKey` — Order key to sort by.
    
</dd>
</dl>

<dl>
<dd>

**orderDirection:** `*vitalgo.GetOrdersLabTestsRequestOrderDirection` — Order direction to sort by.
    
</dd>
</dl>

<dl>
<dd>

**orderType:** `[]*vitalgo.LabTestCollectionMethod` — Filter by method used to perform the lab test.
    
</dd>
</dl>

<dl>
<dd>

**isCritical:** `*bool` — Filter by critical order status.
    
</dd>
</dl>

<dl>
<dd>

**interpretation:** `*vitalgo.Interpretation` — Filter by result interpretation of the lab test.
    
</dd>
</dl>

<dl>
<dd>

**orderActivationTypes:** `[]*vitalgo.OrderActivationType` — Filter by activation type.
    
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

**orderIds:** `[]string` — Filter by order ids.
    
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

<details><summary><code>client.LabTests.GetPhlebotomyAppointmentAvailability(request) -> *vitalgo.AppointmentAvailabilitySlots</code></summary>
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
request := &vitalgo.GetPhlebotomyAppointmentAvailabilityLabTestsRequest{
        Body: &vitalgo.UsAddress{
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

**startDate:** `*time.Time` — Start date for appointment availability
    
</dd>
</dl>

<dl>
<dd>

**request:** `*vitalgo.UsAddress` — At-home phlebotomy appointment address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.BookPhlebotomyAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
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
request := &vitalgo.BookPhlebotomyAppointmentLabTestsRequest{
        Body: &vitalgo.AppointmentBookingRequest{
            BookingKey: "booking_key",
        },
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

**request:** `*vitalgo.AppointmentBookingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.RequestPhlebotomyAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
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
request := &vitalgo.RequestAppointmentRequest{
        Address: &vitalgo.UsAddress{
            FirstLine: "first_line",
            City: "city",
            State: "state",
            ZipCode: "zip_code",
        },
        Provider: vitalgo.AppointmentProviderGetlabs,
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

**address:** `*vitalgo.UsAddress` — At-home phlebotomy appointment address.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.AppointmentProvider` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ReschedulePhlebotomyAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
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
request := &vitalgo.ReschedulePhlebotomyAppointmentLabTestsRequest{
        Body: &vitalgo.AppointmentRescheduleRequest{
            BookingKey: "booking_key",
        },
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

**request:** `*vitalgo.AppointmentRescheduleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CancelPhlebotomyAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
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
request := &vitalgo.ApiApiV1EndpointsVitalApiLabTestingOrdersHelpersAppointmentCancelRequest{
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

<details><summary><code>client.LabTests.GetPhlebotomyAppointmentCancellationReason() -> []*vitalgo.ClientFacingAppointmentCancellationReason</code></summary>
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

<details><summary><code>client.LabTests.GetPhlebotomyAppointment(OrderId) -> *vitalgo.ClientFacingAppointment</code></summary>
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

<details><summary><code>client.LabTests.GetAreaInfo() -> *vitalgo.AreaInfo</code></summary>
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
request := &vitalgo.GetAreaInfoLabTestsRequest{
        ZipCode: "zip_code",
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

**radius:** `*vitalgo.AllowedRadius` — Radius in which to search in miles
    
</dd>
</dl>

<dl>
<dd>

**lab:** `*vitalgo.ClientFacingLabs` — Lab to check for PSCs
    
</dd>
</dl>

<dl>
<dd>

**labs:** `[]*vitalgo.ClientFacingLabs` — List of labs to check for PSCs
    
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

<details><summary><code>client.LabTests.GetPscInfo() -> *vitalgo.PscInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetPscInfoLabTestsRequest{
        ZipCode: "zip_code",
        LabId: 1,
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

**radius:** `*vitalgo.AllowedRadius` — Radius in which to search in miles. Note that we limit to 30 PSCs.
    
</dd>
</dl>

<dl>
<dd>

**capabilities:** `[]*vitalgo.LabLocationCapability` — Filter for only locations with certain capabilities
    
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

<details><summary><code>client.LabTests.GetOrderPscInfo(OrderId) -> *vitalgo.PscInfo</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetOrderPscInfoLabTestsRequest{}
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

**radius:** `*vitalgo.AllowedRadius` — Radius in which to search in miles
    
</dd>
</dl>

<dl>
<dd>

**capabilities:** `[]*vitalgo.LabLocationCapability` — Filter for only locations with certain capabilities
    
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

<details><summary><code>client.LabTests.GetResultMetadata(OrderId) -> *vitalgo.LabResultsMetadata</code></summary>
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

<details><summary><code>client.LabTests.GetResultRaw(OrderId) -> *vitalgo.LabResultsRaw</code></summary>
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
request := &vitalgo.GetLabelsPdfLabTestsRequest{
        CollectionDate: vitalgo.MustParseDateTime(
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

<details><summary><code>client.LabTests.GetPscAppointmentAvailability() -> *vitalgo.AppointmentAvailabilitySlots</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetPscAppointmentAvailabilityLabTestsRequest{
        Lab: vitalgo.AppointmentPscLabs(
            "quest",
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

**lab:** `vitalgo.AppointmentPscLabs` — Lab to check for availability
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Start date for appointment availability
    
</dd>
</dl>

<dl>
<dd>

**siteCodes:** `[]string` — List of site codes to fetch availability for
    
</dd>
</dl>

<dl>
<dd>

**zipCode:** `*string` — Zip code of the area to check
    
</dd>
</dl>

<dl>
<dd>

**radius:** `*vitalgo.AllowedRadius` — Radius in which to search. (meters)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.BookPscAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.BookPscAppointmentLabTestsRequest{
        Body: &vitalgo.AppointmentBookingRequest{
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

**request:** `*vitalgo.AppointmentBookingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ReschedulePscAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ReschedulePscAppointmentLabTestsRequest{
        Body: &vitalgo.AppointmentRescheduleRequest{
            BookingKey: "booking_key",
        },
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

**request:** `*vitalgo.AppointmentRescheduleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.CancelPscAppointment(OrderId, request) -> *vitalgo.ClientFacingAppointment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.VitalCoreClientsLabTestGetlabsSchemaAppointmentCancelRequest{
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

<details><summary><code>client.LabTests.GetPscAppointmentCancellationReason() -> []*vitalgo.ClientFacingAppointmentCancellationReason</code></summary>
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

<details><summary><code>client.LabTests.GetPscAppointment(OrderId) -> *vitalgo.ClientFacingAppointment</code></summary>
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

<details><summary><code>client.LabTests.GetOrder(OrderId) -> *vitalgo.ClientFacingOrder</code></summary>
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

<details><summary><code>client.LabTests.CreateOrder(request) -> *vitalgo.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CreateOrderRequestCompatible{
        UserId: "user_id",
        PatientDetails: &vitalgo.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: vitalgo.MustParseDateTime(
                "2023-01-15",
            ),
            Gender: vitalgo.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &vitalgo.PatientAddressWithValidation{
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

**orderSet:** `*vitalgo.OrderSetRequest` 
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*vitalgo.LabTestCollectionMethod` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*vitalgo.PhysicianCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**healthInsurance:** `*vitalgo.HealthInsuranceCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*bool` — Defines whether order is priority or not. For some labs, this refers to a STAT order.
    
</dd>
</dl>

<dl>
<dd>

**billingType:** `*vitalgo.Billing` 
    
</dd>
</dl>

<dl>
<dd>

**icdCodes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**consents:** `[]*vitalgo.Consent` 
    
</dd>
</dl>

<dl>
<dd>

**activateBy:** `*time.Time` — Schedule an Order to be processed in a future date.
    
</dd>
</dl>

<dl>
<dd>

**aoeAnswers:** `[]*vitalgo.AoEAnswer` 
    
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

<dl>
<dd>

**creatorMemberId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patientDetails:** `*vitalgo.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*vitalgo.PatientAddressWithValidation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.ImportOrder(request) -> *vitalgo.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ImportOrderBody{
        UserId: "user_id",
        BillingType: vitalgo.BillingClientBill,
        OrderSet: &vitalgo.OrderSetRequest{},
        CollectionMethod: vitalgo.LabTestCollectionMethodTestkit,
        PatientDetails: &vitalgo.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: vitalgo.MustParseDateTime(
                "2023-01-15",
            ),
            Gender: vitalgo.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &vitalgo.PatientAddress{
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

**billingType:** `*vitalgo.Billing` 
    
</dd>
</dl>

<dl>
<dd>

**orderSet:** `*vitalgo.OrderSetRequest` 
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*vitalgo.LabTestCollectionMethod` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*vitalgo.PhysicianCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**patientDetails:** `*vitalgo.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*vitalgo.PatientAddress` 
    
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

<details><summary><code>client.LabTests.CancelOrder(OrderId) -> *vitalgo.PostOrderResponse</code></summary>
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
request := &vitalgo.SimulateOrderProcessLabTestsRequest{
        Body: &vitalgo.SimulationFlags{},
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

**finalStatus:** `*vitalgo.OrderStatus` 
    
</dd>
</dl>

<dl>
<dd>

**delay:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*vitalgo.SimulationFlags` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LabTests.UpdateOnSiteCollectionOrderDrawCompleted(OrderId) -> *vitalgo.PostOrderResponse</code></summary>
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

<details><summary><code>client.LabTests.ValidateIcdCodes(request) -> *vitalgo.ValidateIcdCodesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.ValidateIcdCodesBody{
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

## Testkit
<details><summary><code>client.Testkit.Register(request) -> *vitalgo.PostOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.RegisterTestkitRequest{
        SampleId: "sample_id",
        PatientDetails: &vitalgo.PatientDetailsWithValidation{
            FirstName: "first_name",
            LastName: "last_name",
            Dob: vitalgo.MustParseDateTime(
                "2023-01-15",
            ),
            Gender: vitalgo.GenderFemale,
            PhoneNumber: "phone_number",
            Email: "email",
        },
        PatientAddress: &vitalgo.PatientAddressWithValidation{
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

**patientDetails:** `*vitalgo.PatientDetailsWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**patientAddress:** `*vitalgo.PatientAddressWithValidation` 
    
</dd>
</dl>

<dl>
<dd>

**physician:** `*vitalgo.PhysicianCreateRequestBase` 
    
</dd>
</dl>

<dl>
<dd>

**healthInsurance:** `*vitalgo.HealthInsuranceCreateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**consents:** `[]*vitalgo.Consent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Testkit.CreateOrder(request) -> *vitalgo.PostOrderResponse</code></summary>
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
request := &vitalgo.CreateRegistrableTestkitOrderRequest{
        UserId: "user_id",
        LabTestId: "lab_test_id",
        ShippingDetails: &vitalgo.ShippingAddressWithValidation{
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

**shippingDetails:** `*vitalgo.ShippingAddressWithValidation` 
    
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
<details><summary><code>client.Order.ResendEvents(request) -> *vitalgo.ResendWebhookResponse</code></summary>
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
request := &vitalgo.ResendWebhookBody{}
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
<details><summary><code>client.Insurance.SearchGetPayorInfo() -> []*vitalgo.ClientFacingPayorSearchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.SearchGetPayorInfoInsuranceRequest{}
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

**provider:** `*vitalgo.PayorCodeExternalProvider` 
    
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

<details><summary><code>client.Insurance.SearchPayorInfo(request) -> []*vitalgo.ClientFacingPayorSearchResponseDeprecated</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.PayorSearchRequest{}
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

**provider:** `*vitalgo.PayorCodeExternalProvider` 
    
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

<details><summary><code>client.Insurance.SearchDiagnosis() -> []*vitalgo.ClientFacingDiagnosisInformation</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.SearchDiagnosisInsuranceRequest{
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
<details><summary><code>client.Payor.CreatePayor(request) -> *vitalgo.ClientFacingPayor</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.CreatePayorBody{
        Name: "name",
        Address: &vitalgo.Address{
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

**address:** `*vitalgo.Address` 
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*vitalgo.PayorCodeExternalProvider` 
    
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
<details><summary><code>client.LabReport.ParserCreateJob(request) -> *vitalgo.ParsingJob</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a parse job, uploads the file to provider, persists the job row,
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
request := &vitalgo.BodyCreateLabReportParserJob{
        UserId: "user_id",
    }
client.LabReport.ParserCreateJob(
        context.TODO(),
        strings.NewReader(
            "",
        ),
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

<details><summary><code>client.LabReport.ParserGetJob(JobId) -> *vitalgo.ParsingJob</code></summary>
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
<details><summary><code>client.Aggregate.QueryOne(UserId, request) -> *vitalgo.AggregationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.QueryBatch{
        Timeframe: &vitalgo.QueryBatchTimeframe{
            RelativeTimeframe: &vitalgo.RelativeTimeframe{
                Anchor: vitalgo.MustParseDateTime(
                    "2023-01-15",
                ),
                Past: &vitalgo.Period{
                    Unit: vitalgo.PeriodUnitMinute,
                },
            },
        },
        Queries: []*vitalgo.Query{
            &vitalgo.Query{
                Select: []*vitalgo.QuerySelectItem{
                    &vitalgo.QuerySelectItem{
                        AggregateExpr: &vitalgo.AggregateExpr{
                            Arg: &vitalgo.AggregateExprArg{
                                SleepColumnExpr: &vitalgo.SleepColumnExpr{
                                    Sleep: vitalgo.SleepColumnExprSleepId,
                                },
                            },
                            Func: vitalgo.AggregateExprFuncMean,
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

**timeframe:** `*vitalgo.QueryBatchTimeframe` 
    
</dd>
</dl>

<dl>
<dd>

**queries:** `[]*vitalgo.Query` 
    
</dd>
</dl>

<dl>
<dd>

**config:** `*vitalgo.QueryConfig` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Aggregate.GetResultTableForContinuousQuery(UserId, QueryIdOrSlug) -> *vitalgo.AggregationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetResultTableForContinuousQueryAggregateRequest{}
client.Aggregate.GetResultTableForContinuousQuery(
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

**accept:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Aggregate.GetTaskHistoryForContinuousQuery(UserId, QueryIdOrSlug) -> *vitalgo.ContinuousQueryTaskHistoryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vitalgo.GetTaskHistoryForContinuousQueryAggregateRequest{}
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
