if ! getent group spiegel > /dev/null; then
  groupadd -r spiegel
fi

if ! getent passwd spiegel > /dev/null; then
    useradd -r -d /var/lib/spiegel -s /sbin/nologin -g spiegel -c "Spiegel server" spiegel
fi
if ! test -d /var/lib/spiegel; then
    mkdir -p /var/lib/spiegel
    chmod 0750 /var/lib/spiegel
    chown -R spiegel:spiegel /var/lib/spiegel
fi
