package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17 2v20V2Zm-5 0v20V2ZM7 2v20V2ZM2 22h20V2H2v20Z"/>`),
		g.Group(children),
	)
}