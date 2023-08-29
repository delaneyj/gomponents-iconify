package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M20 24h24h-24Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 24h24"/><path d="M20 38h24h-24Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 38h24"/><path d="M20 10h24h-24Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 10h24"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M4 34h8v8H4zm0-14h8v8H4zM4 6h8v8H4z"/></g>`),
		g.Group(children),
	)
}