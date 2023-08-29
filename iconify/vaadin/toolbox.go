package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 8h6v2h4V8h6v6H0z"/><path fill="currentColor" d="M7 7h2v2H7V7zm4-3V2H5v2H0v3h6V6h4v1h6V4h-5zM6 4V3h4v1H6z"/>`),
		g.Group(children),
	)
}