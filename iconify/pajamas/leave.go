package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.16 4.28a.75.75 0 1 1 1.06-1.06l3.25 3.25L16 7l-.53.53l-3.25 3.25a.75.75 0 0 1-1.06-1.06l1.97-1.97H3.25a1.75 1.75 0 1 0 0 3.5h2a.75.75 0 0 1 0 1.5h-2a3.25 3.25 0 0 1 0-6.5h9.88l-1.97-1.97Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}