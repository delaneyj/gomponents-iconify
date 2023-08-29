package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.66 40.3c4.065-3.752 6.41-9.264 6.41-15.811c0-1.525-.137-2.99-.39-4.398H24.43v8.326H36c-.507 2.678-2.032 4.945-4.319 6.47M5.236 33.656C8.774 40.672 16.026 45.5 24.43 45.5c5.805 0 10.672-1.915 14.23-5.2l-6.979-5.413c-1.915 1.29-4.358 2.072-7.251 2.072c-5.59 0-10.34-3.773-12.04-8.854M5.236 14.364C3.77 17.257 2.93 20.521 2.93 24s.84 6.743 2.306 9.636l5.57-4.34l1.584-1.211c-.43-1.29-.684-2.658-.684-4.085s.254-2.795.684-4.085"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.43 11.06c3.166 0 5.98 1.095 8.229 3.206l6.156-6.156C35.082 4.63 30.235 2.5 24.43 2.5c-8.405 0-15.656 4.828-19.194 11.864l7.154 5.551c1.7-5.082 6.45-8.854 12.04-8.854Z"/>`),
		g.Group(children),
	)
}