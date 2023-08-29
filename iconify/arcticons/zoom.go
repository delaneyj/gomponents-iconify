package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M34.86 18.63v10.53l8.64 7.16V11.68ZM6.42 12.31h21.7c3.16 0 6.53 2.53 6.53 5.9v15.16a2.11 2.11 0 0 1-2.1 2.11H10.22c-3.59 0-5.69-3.37-5.69-6.53V14.41a1.83 1.83 0 0 1 1.89-2.1Z"/>`),
		g.Group(children),
	)
}