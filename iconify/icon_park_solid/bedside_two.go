package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBedsideTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 18h40v24H4V18Z"/><path stroke="#000" d="M22 24h4M4 30h40m-22 6h4"/><path stroke="#fff" d="M8 42v4m32-4v4M24 18v-8"/><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M32 10a8 8 0 1 0-16 0h16Z" clip-rule="evenodd"/><path stroke="#fff" d="M44 26v8M4 26v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBedsideTwo0)"/>`),
		g.Group(children),
	)
}