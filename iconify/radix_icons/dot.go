package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 9.125a1.625 1.625 0 1 0 0-3.25a1.625 1.625 0 0 0 0 3.25Zm0 1a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}