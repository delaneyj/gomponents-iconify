package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 13h-4a.499.499 0 0 1-.457-.297L6.175 4H1.5a.5.5 0 0 1 0-1h5c.198 0 .377.116.457.297L10.825 12H14.5a.5.5 0 0 1 0 1zm0-9h-5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}