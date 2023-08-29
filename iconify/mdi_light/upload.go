package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.16V5.91l5.25 5.25l.75-.66L11.5 4L5 10.5l.75.66L11 5.91v12.25h1M3 19h1v2h15v-2h1v3H3v-3Z"/>`),
		g.Group(children),
	)
}