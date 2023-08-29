package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForAlandIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-.983 36.884v21.091a27.845 27.845 0 0 1-2.95-.256V35.934h31.652a28.113 28.113 0 0 1-.577 2.95H31.017m-26.154.022a28.042 28.042 0 0 1-.582-2.973H20.2v21.45a27.922 27.922 0 0 1-2.95-1.599V38.906H4.863M17.25 25.117V8.215a28.021 28.021 0 0 1 2.95-1.599v21.45H4.281c.142-1 .332-1.985.577-2.949H17.25m10.816 2.949V4.281a27.845 27.845 0 0 1 2.95-.256v21.092h28.125c.245.964.436 1.949.577 2.949H28.066"/>`),
		g.Group(children),
	)
}