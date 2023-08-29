package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lovoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.483a12.277 12.277 0 0 1 6.167-1.65h0A12.333 12.333 0 0 1 42.5 30.167V42.5H30.167a12.333 12.333 0 0 1-12.334-12.333h0a12.276 12.276 0 0 1 1.652-6.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.167 42.5H17.833A12.333 12.333 0 0 1 5.5 30.167h0a12.333 12.333 0 0 1 12.333-12.334h0a12.333 12.333 0 0 1 12.334 12.334"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.167 5.5h0A12.333 12.333 0 0 1 42.5 17.833v12.334h0h-12.333a12.333 12.333 0 0 1-12.334-12.334v0A12.333 12.333 0 0 1 30.167 5.5Z"/>`),
		g.Group(children),
	)
}