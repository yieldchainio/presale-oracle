# presale-oracle

This is a simple oracle watching configured contracts and counting contributions in them to make sure the total contirbution does not go over the set hard cap. Once the sum of contributions reaches the hard cap, the oracle closes the presales.

## Build

```
go build
```

## Install & Run

```
cp presale-oracle /opt
cp systemd/presale-oracle.service /etc/systemd/system/
systemctl enable presale-oracle.service
systemctl start presale-oracle.service
```