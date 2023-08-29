package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M961.12 1024q-26.5 0-45.5-19t-19-45V571l-386 442q-24 27-62-13V570l-386 443q-25 27-62-13V24q37-40 62-13l386 444V24q37-40 62-13l386 443V64q0-27 19-45.5T961.12 0t45 18.5t18.5 45.5v896q0 26-18.5 45t-45 19z"/>`),
		g.Group(children),
	)
}