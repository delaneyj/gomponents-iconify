package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M44.426 2H19.574L2 19.574v24.852L19.574 62h24.852L62 44.426V19.574L44.426 2zM60 43.598L43.598 60H20.402L4 43.598V20.402L20.402 4h23.195L60 20.402v23.196z"/><path fill="currentColor" d="M22.473 9L9 22.474v19.051L22.473 55h19.052L55 41.525V22.474L41.525 9z"/>`),
		g.Group(children),
	)
}