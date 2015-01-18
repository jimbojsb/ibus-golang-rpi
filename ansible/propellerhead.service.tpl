[Unit]
Description=Propellerhead Service
After=network.target

[Service]
Type=oneshot
ExecStart=/root/ph/propellerhead /dev/ttyUSB0

[Install]
WantedBy=multi-user.target