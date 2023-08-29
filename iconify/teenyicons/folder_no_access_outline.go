package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderNoAccessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m9.5 6.5l-4 4m-5-8v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Zm7 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}