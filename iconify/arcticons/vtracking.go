package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vtracking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.928 19.964c.106-.8.163-1.593.163-2.373C37.091 10.361 31.231 4.5 24 4.5s-13.091 5.861-13.091 13.091c0 10.251 10 22.612 12.61 25.633a.8.8 0 0 0 1.211 0c1.448-1.704 5.236-6.43 8.24-12.075"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.5 26.997a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm12.001 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-12.001 0h8m-12 0h-2v-7m11-2v-2h-14m21.001 11h2v-5m0 0l-5-5m-.001-.001h-4v10.001m-11-9h-2m2 2h-1m14-1h4m-4 3.5h7.001m-7.001-3.5v3.5"/>`),
		g.Group(children),
	)
}