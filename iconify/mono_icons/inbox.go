package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5zm2 9v5h14v-5h-2.28l-.771 2.316A1 1 0 0 1 15 17H9a1 1 0 0 1-.949-.684L7.28 14H5zm14-2V5H5v7h2.28a2 2 0 0 1 1.897 1.367L9.72 15h4.558l.544-1.633A2 2 0 0 1 16.721 12H19z"/>`),
		g.Group(children),
	)
}