package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackgroundColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M37 37a4 4 0 0 0 4-4c0-1.473-1.333-3.473-4-6c-2.667 2.527-4 4.527-4 6a4 4 0 0 0 4 4Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m20.854 5.504l3.535 3.536"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M23.682 8.333L8.125 23.889L19.44 35.203l15.556-15.557L23.682 8.333Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m12 20.073l16.961 5.577M4 43h40"/></g>`),
		g.Group(children),
	)
}