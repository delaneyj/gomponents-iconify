package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageAltCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4v12c0 1.103.897 2 2 2h3.586L12 21.414L15.414 18H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2zm2 0h14v12h-4.414L12 18.586L9.414 16H5V4z"/><path fill="currentColor" d="m17.207 7.207l-1.414-1.414L11 10.586L8.707 8.293L7.293 9.707L11 13.414z"/>`),
		g.Group(children),
	)
}