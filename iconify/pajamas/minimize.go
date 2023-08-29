package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 2.28a.75.75 0 0 0-1.06-1.06L10.5 4.44V2.75a.75.75 0 0 0-1.5 0V7h4.25a.75.75 0 0 0 0-1.5h-1.69l3.22-3.22ZM5.5 11.56v1.69a.75.75 0 0 0 1.5 0V9H2.75a.75.75 0 0 0 0 1.5h1.69l-3.22 3.22a.75.75 0 1 0 1.06 1.06l3.22-3.22Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}