package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.07 7L6 11.606V20h12V7H9.07ZM8 5h11a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V11l4-6Zm5-4h5a1 1 0 0 1 1 1v2h-7V2a1 1 0 0 1 1-1ZM8 12h2v6H8v-6Z"/>`),
		g.Group(children),
	)
}