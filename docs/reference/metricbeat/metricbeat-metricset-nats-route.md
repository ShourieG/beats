---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-metricset-nats-route.html
---

# NATS route metricset [metricbeat-metricset-nats-route]

This is the route metricset of the module nats collecting metrcis per connection.

## Fields [_fields_190]

For a description of each field in the metricset, see the [exported fields](/reference/metricbeat/exported-fields-nats.md) section.

Here is an example document generated by this metricset:

```json
{
    "@timestamp":"2016-05-23T08:05:34.853Z",
    "beat":{
        "hostname":"beathost",
        "name":"beathost"
    },
    "metricset":{
        "host":"localhost",
        "module":"nats",
        "name":"routes",
        "rtt":44269
    },
    "nats":{
        "server": {
            "id": "bUAdpRFtMWddIBWw80Yd9D",
            "time": "2018-12-28T12:33:53.026865597Z"
        },
        "routes": {
            "in": {
                "bytes": 0,
                "messages": 0
            },
            "ip": "172.25.255.2",
            "out": {
                "bytes": 0,
                "messages": 0
            },
            "pending_size": 0,
            "port": 60736,
            "remote_id": "NCYX4P3ZCXZT3UF4BBZUG46S7EU224XXXLUTUBUDOIAUAIR5WJV73BV5",
            "subscriptions": 0,
        }
    },
    "type":"metricsets"
}
```


