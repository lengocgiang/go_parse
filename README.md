# go_parse

Input
```http
	URL := "localhost:8080/bk/api/actions.json?distinct_id=763_1495187301909_3495&timestamp=1495187523&event=product_add_cart&params%5BproductRefId%5D=8284563078&params%5Bapps%5D%5B%5D=precommend&params%5Bapps%5D%5B%5D=bsales&params%5Bsource%5D=item&params%5Boptions%5D%5Bsegment%5D=cart_recommendation&params%5Boptions%5D%5Btype%5D=up_sell&params%5BtimeExpire%5D=1495187599642&params%5Brecommend_system_product_source%5D=item&params%5Bproduct_id%5D=8284563078&params%5Bvariant_id%5D=27661944134&params%5Bsku%5D=00483332%20(black)&params%5Bsources%5D%5B%5D=product_recommendation&params%5Bcart_token%5D=dc2c336a009edf2762128e65806dfb1d&params%5Bquantity%5D=1&params%5Bnew_popup_upsell_mobile%5D=false&params%5BclientDevice%5D=desktop&params%5BclientIsMobile%5D=false&params%5BclientIsSmallScreen%5D=false&params%5Bnew_popup_crossell_mobile%5D=false&api_key=14c5b7dacea9157029265b174491d340"

```
Output
```json
{
    "api_key": "14c5b7dacea9157029265b174491d340",
    "distinct_id": "763_1495187301909_3495",
    "event": "product_add_cart",
    "params": {
        "apps": [
            "precommend",
            "bsales"
        ],
        "cart_token": "dc2c336a009edf2762128e65806dfb1d",
        "clientDevice": "desktop",
        "clientIsMobile": "false",
        "clientIsSmallScreen": "false",
        "new_popup_crossell_mobile": "false",
        "new_popup_upsell_mobile": "false",
        "optionssegment": "cart_recommendation",
        "optionstype": "up_sell",
        "productRefId": "8284563078",
        "product_id": "8284563078",
        "quantity": "1",
        "recommend_system_product_source": "item",
        "sku": "00483332 (black)",
        "source": "item",
        "sources": "product_recommendation",
        "timeExpire": "1495187599642",
        "variant_id": "27661944134"
    },
    "timestamp": "1495187523"
}
```