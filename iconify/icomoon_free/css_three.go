package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m2.381.758l-.537 2.686h10.934l-.342 1.735H1.496l-.53 2.686h10.933l-.61 3.063l-4.406 1.46l-3.819-1.46l.261-1.329H.639L0 12.823l6.316 2.417l7.281-2.417L16 .757z"/>`),
		g.Group(children),
	)
}