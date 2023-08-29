package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5 8.5h5M.5.5h14v4H.5v-4Zm1 4v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9h-12Z"/>`),
		g.Group(children),
	)
}