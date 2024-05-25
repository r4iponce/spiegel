if getent passwd spiegel > /dev/null; then
  userdel -r spiegel
fi

if getent group spiegel > /dev/null; then
  groupdel spiegel
fi