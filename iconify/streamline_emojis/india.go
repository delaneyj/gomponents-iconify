package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func India(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#ff8a14" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2 15.21h7.37A45.73 45.73 0 0 1 24 17.59A45.62 45.62 0 0 0 38.68 20h7.37v-7a.94.94 0 0 0-1-.79h-6.37A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79Z"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2 34.3a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45c.58 0 1-.35 1-.79v-6.92h-7.32A45.62 45.62 0 0 1 24 29.75a45.72 45.72 0 0 0-14.63-2.37H2Z"/><path fill="none" stroke="#009fd9" stroke-linecap="round" stroke-linejoin="round" d="M20.157 25.436a4.77 4.26 65.41 1 0 7.747-3.545a4.77 4.26 65.41 1 0-7.747 3.545Z"/><path fill="#009fd9" d="M22.139 24.53a2.33 2.08 65.41 1 0 3.783-1.732a2.33 2.08 65.41 1 0-3.783 1.731Z"/><path fill="none" stroke="#009fd9" stroke-linecap="round" stroke-linejoin="round" d="M24.03 19v9.34m-4.36-5.06l8.71.78M27 20.53l-5.91 6.31m-.17-6.73L27 27.12"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}