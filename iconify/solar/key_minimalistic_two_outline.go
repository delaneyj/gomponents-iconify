package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 2.75a6.25 6.25 0 1 0 0 12.5a6.25 6.25 0 0 0 0-12.5ZM7.25 9a7.75 7.75 0 1 1 2.824 5.983a.711.711 0 0 1-.044.047L7.56 17.5l.97.97a.75.75 0 1 1-1.06 1.06l-.97-.97l-.94.94l.97.97a.75.75 0 1 1-1.06 1.06l-.97-.97l-.47.47a.75.75 0 0 1-1.06-1.06l1-1l5-5a.73.73 0 0 1 .047-.044A7.718 7.718 0 0 1 7.25 9ZM15 7.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5ZM12.25 9a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}