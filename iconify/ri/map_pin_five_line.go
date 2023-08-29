package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinFiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18.485l4.243-4.242a6 6 0 1 0-8.486 0L12 18.485Zm5.657-2.828L12 21.314l-5.657-5.657a8 8 0 1 1 11.314 0ZM5 22h14v2H5v-2Z"/>`),
		g.Group(children),
	)
}