package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 9H13V6h.5a.5.5 0 0 0 0-1H13V3l-.286-.573a.249.249 0 0 0-.424-.006L12 3v2h-1V3l-.286-.573a.249.249 0 0 0-.424-.006L10 3v2H9V3l-.286-.573a.25.25 0 0 0-.424-.006L8 3v2H7V3l-.286-.573a.25.25 0 0 0-.424-.006L6 3v2H5V3l-.286-.573a.25.25 0 0 0-.424-.006L4 3v2H3V3l-.286-.573a.25.25 0 0 0-.424-.006L2 3v2h-.5a.5.5 0 1 0 0 1H2v3h-.5a.5.5 0 1 0 0 1H2v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h.5a.5.5 0 0 0 0-1zM3 9V6h1v3zm2 0V6h1v3zm2 0V6h1v3zm2 0V6h1v3zm2 0V6h1v3z"/>`),
		g.Group(children),
	)
}