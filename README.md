# BSKU001 Backend

SKU module backend built with Clean Architecture, Gin, GORM, PostgreSQL, zap, JWT, and Swagger UI.

## Run

```bash
cp .env.example .env
go run ./cmd/server
```

Swagger UI: `http://localhost:8080/docs/swagger.html`

## Auth

All `/api/v1/*` endpoints require `Authorization: Bearer <jwt>`.
JWT must include `merchant_id` claim or request header `X-Merchant-ID` as fallback, and scopes such as:

- `sku:read`
- `sku:create`
- `sku:update`
- `sku:delete`
- `sku:*`

## Tables

- `skus`
- `sku_dimensions`
- `sku_images`
- `sku_option_groups`
- `sku_option_values`
- `brands`
- `categories`
- `sub_categories`
- `materials`
- `colors`

## Swagger Failed to Fetch

Open Swagger through the running backend URL, for example `http://localhost:8080/docs/swagger.html`. Do not open `docs/swagger.html` directly from the filesystem, because browser CORS requires an `http` or `https` URL scheme. The OpenAPI server URL uses `/api/v1` relative to the current host.


## SKU Options Extension

Merchants can enable/disable the SKU options extension through `merchant_feature_configs.sku_options`, returned by `GET /api/v1/merchant-config` as `sku_options`. Update it with `PATCH /api/v1/merchant-config`, for example `{ "sku_options": false }`.

Use option groups for product customizations:

- `selection_type: RADIO` for single-choice options such as sweetness level.
- `selection_type: CHECKBOX` for multi-choice options such as toppings.

Main endpoints:

- `GET /api/v1/skus/{id_or_code}/options`
- `POST /api/v1/skus/{id_or_code}/options`
- `GET /api/v1/sku-option-groups`
- `PUT /api/v1/sku-option-groups/{id}`
- `POST /api/v1/sku-option-groups/{id}/options`
- `PUT /api/v1/sku-option-values/{id}`
# b_sku001
