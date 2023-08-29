package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm1.78 8.841a.75.75 0 0 1-1.06 0l-1.97-1.97v6.379a.75.75 0 0 1-1.5 0V8.871l-1.97 1.97a.75.75 0 1 1-1.06-1.06l3.25-3.25L8 6l.53.53l3.25 3.25a.75.75 0 0 1 0 1.061Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}