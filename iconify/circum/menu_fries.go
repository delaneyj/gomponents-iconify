package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.437 19.937a.5.5 0 0 1 0 1l-16.874.002a.5.5 0 0 1 0-1l16.874-.002Zm0-8.437a.5.5 0 0 1 0 1l-10 .001a.5.5 0 0 1 0-1l10-.001Zm0-8.438a.5.5 0 0 1 0 1l-16.874.001a.5.5 0 0 1 0-1l16.874-.001Z"/>`),
		g.Group(children),
	)
}