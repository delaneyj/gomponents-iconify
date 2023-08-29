package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM8 12c1.953 0 3.579-1.4 3.93-3.25h3.32a.75.75 0 0 0 0-1.5h-3.32a4.001 4.001 0 0 0-7.86 0H.75a.75.75 0 0 0 0 1.5h3.32A4.001 4.001 0 0 0 8 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}