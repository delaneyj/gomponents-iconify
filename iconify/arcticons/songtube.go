package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Songtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.132 4.5v20.198a8.26 8.26 0 0 0 8.261 8.26h4.01V4.5Zm27.736 0H26.01v39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.009 43.5h5.534a6.324 6.324 0 0 0 6.324-6.324V4.5"/>`),
		g.Group(children),
	)
}