package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RampRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-6.3q-.825 1.125-1.975 2.138T6.45 18.725L5 17.275q.75-.425 1.775-1.175t1.963-1.788q.937-1.037 1.6-2.374T11 9V6.825L9.4 8.4L8 7l4-4l4 4l-1.4 1.4L13 6.825V21h-2Z"/>`),
		g.Group(children),
	)
}