package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 8a4 4 0 0 1 4-4h6a4 4 0 0 1 0 8H9.75a.75.75 0 0 1 0-1.5H11a2.5 2.5 0 0 0 0-5H5a2.5 2.5 0 0 0-.735 4.89c-.16.468-.25.966-.263 1.484A4.002 4.002 0 0 1 1 8Zm14.998.126a4.988 4.988 0 0 1-.263 1.484A2.501 2.501 0 0 1 15 14.5H9a2.5 2.5 0 0 1 0-5h1.25a.75.75 0 0 0 0-1.5H9a4 4 0 1 0 0 8h6a4 4 0 0 0 .998-7.874Z"/>`),
		g.Group(children),
	)
}