package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBedsideTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 18h40v24H4V18Z"/><path d="M22 24h4M4 30h40m-22 6h4M8 42v4m32-4v4M24 18v-8"/><path fill="#555" fill-rule="evenodd" d="M32 10a8 8 0 1 0-16 0h16Z" clip-rule="evenodd"/><path d="M44 26v8M4 26v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBedsideTwo0)"/>`),
		g.Group(children),
	)
}