package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningRoomTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 7.5a4 4 0 0 1 4-4h1v-3h-11v3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2v10h11v-7l-.145-.087C14 14 12.5 11.5 12.5 7.5Zm0 0h-4m7-3.874V5A2.5 2.5 0 0 0 18 7.5M20.5 0v4"/>`),
		g.Group(children),
	)
}