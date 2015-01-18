[Unit]
Description=Propellerhead Service
After=network.target

[Service]
Type=oneshot
WorkingDirectory=/root/ph
ExecStart=/root/ph/propellerhead /dev/ttyUSB0

[Install]
WantedBy=multi-user.target