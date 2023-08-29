package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.083 14H14V4.083H4.083V14zM17 4.083V14h9.917V4.083H17zm0 22.834h9.917V17H17v9.917zm-12.917 0H14V17H4.083v9.917z"/>`),
		g.Group(children),
	)
}