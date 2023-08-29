package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivitySource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSActivitySource0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M42 5H6v8h36V5Zm0 15H6v8h36v-8Zm0 15H6v8h36v-8Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSActivitySource0)"/>`),
		g.Group(children),
	)
}