package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symphony(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="21.5"/><circle cx="23.795" cy="23.729" r="5.09"/><path d="M11.134 21.008a13.087 13.087 0 0 1 10.087-9.767m-14.069 8.67A17.12 17.12 0 0 1 19.448 7.368m17.38 19.918a13.087 13.087 0 0 1-10.087 9.767m14.069-8.67a17.12 17.12 0 0 1-12.296 12.544"/><ellipse cx="8.931" cy="42.848" rx="3.338" ry="2.115"/><path d="M12.261 42.765V28.653m3.999 3.281a3.553 3.553 0 0 1-3.952-3.205M41.245 11.17H32.52l-2.512 2.512"/></g>`),
		g.Group(children),
	)
}