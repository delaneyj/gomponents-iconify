package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.774 1.747a6.5 6.5 0 1 0 1.142 12.062a.75.75 0 0 1 .673 1.34A8 8 0 1 1 16 8v1.25a2.75 2.75 0 0 1-5.072 1.475A4 4 0 1 1 12 8v1.25a1.25 1.25 0 0 0 2.5 0V8a6.5 6.5 0 0 0-4.726-6.253ZM10.5 8a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}