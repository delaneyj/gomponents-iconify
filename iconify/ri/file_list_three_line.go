package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileListThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a3 3 0 0 1-3-3V3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v12h4v4a3 3 0 0 1-3 3Zm-1-5v2a1 1 0 1 0 2 0v-2h-2Zm-2 3V4H4v15a1 1 0 0 0 1 1h11ZM6 7h8v2H6V7Zm0 4h8v2H6v-2Zm0 4h5v2H6v-2Z"/>`),
		g.Group(children),
	)
}