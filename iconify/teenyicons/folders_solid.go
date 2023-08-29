package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldersSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 0A1.5 1.5 0 0 0 3 1.5v7A1.5 1.5 0 0 0 4.5 10h9A1.5 1.5 0 0 0 15 8.5v-5A1.5 1.5 0 0 0 13.5 2H9.707l-2-2H4.5Z"/><path fill="currentColor" d="M12 11H4.5A2.5 2.5 0 0 1 2 8.5V5h-.5A1.5 1.5 0 0 0 0 6.5v7A1.5 1.5 0 0 0 1.5 15h9a1.5 1.5 0 0 0 1.5-1.5V11Z"/>`),
		g.Group(children),
	)
}