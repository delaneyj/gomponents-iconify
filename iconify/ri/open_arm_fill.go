package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenArmFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 12a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm6 5v5h-2v-5c0-4.451 2.644-8.285 6.447-10.016l.828 1.82A9.002 9.002 0 0 0 18.001 17Zm-10 0v5h-2v-5A9.002 9.002 0 0 0 .727 8.805l.827-1.821A11.002 11.002 0 0 1 8.001 17Z"/>`),
		g.Group(children),
	)
}