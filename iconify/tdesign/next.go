package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 4.336L14.414 12L6.75 19.664V4.336ZM17.5 5v14h-2V5h2ZM8.75 9.164v5.672L11.586 12L8.75 9.164Z"/>`),
		g.Group(children),
	)
}