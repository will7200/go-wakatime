# go-wakatime api

Generated from swagger spec for the most part except for some exceptions
1. The swagger spec has not default date parse for the likes of 01/31/2019, so the internal package houses a near replica
of the date type the strfmt that open-api has.
