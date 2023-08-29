package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m0 3l5-2v12l-5 2zM6 .5l5 3V15l-5-2.5zm6 3l4-3v12l-4 3z"/>`),
		g.Group(children),
	)
}