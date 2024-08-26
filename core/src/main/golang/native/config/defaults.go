package config

var (
	defaultNameServers = []string{
		"https://223.5.5.5/dns-query#h3=true",
		"https://1.12.12.12/dns-query#h3=true",
		"tcp://30.30.30.30",
	}
	defaultFakeIPFilter = []string{
		// Stun Services
		"+.stun.*.*",
		"+.stun.*.*.*",
		"+.stun.*.*.*.*",
		"+.stun.*.*.*.*.*",

		// Google Voices
		"lens.l.google.com",

		// Nintendo Switch STUN
		"*.n.n.srv.nintendo.net",

		// PlayStation STUN
		"+.stun.playstation.net",

		// XBox
		"xbox.*.*.microsoft.com",
		"*.*.xboxlive.com",

		// Microsoft Captive Portal
		"*.msftncsi.com",
		"*.msftconnecttest.com",

		// Bilibili CDN
		"*.mcdn.bilivideo.cn",

		// Windows Default LAN WorkGroup
		"WORKGROUP",
	}
	defaultFakeIPRange = "28.0.0.0/8"
)
