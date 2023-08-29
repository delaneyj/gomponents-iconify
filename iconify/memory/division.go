package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Division(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 8h-2V7H9V5h1V4h2v1h1v2h-1Zm5 4H5v-2h12Zm-5 6h-2v-1H9v-2h1v-1h2v1h1v2h-1Z"/>`),
		g.Group(children),
	)
}