package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhosphorLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 96a64 64 0 0 1-64 64h-8V32h8a64 64 0 0 1 64 64ZM64 160h72L64 32Z" opacity=".2"/><path d="M144 24H64a8 8 0 0 0-8 8v128a80.09 80.09 0 0 0 80 80a8 8 0 0 0 8-8v-64a72 72 0 0 0 0-144ZM72 62.54L122.32 152H72Zm56 161A64.14 64.14 0 0 1 72.51 168H128Zm0-94L77.68 40H128ZM144 152V40a56 56 0 0 1 0 112Z"/></g>`),
		g.Group(children),
	)
}