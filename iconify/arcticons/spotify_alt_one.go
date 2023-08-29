package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpotifyAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.86 33.51a1.34 1.34 0 0 1-1.844.443c-5.048-3.083-11.402-3.781-18.886-2.072a1.34 1.34 0 0 1-.597-2.613c8.19-1.871 15.215-1.066 20.882 2.398a1.34 1.34 0 0 1 .445 1.843Zm2.631-5.855a1.676 1.676 0 0 1-2.305.552c-5.78-3.552-14.589-4.58-21.424-2.506a1.679 1.679 0 0 1-2.092-1.116a1.679 1.679 0 0 1 1.117-2.09c7.809-2.37 17.516-1.223 24.152 2.856a1.676 1.676 0 0 1 .552 2.304v0Zm.226-6.096c-6.93-4.116-18.362-4.494-24.978-2.486a2.01 2.01 0 1 1-1.167-3.849c7.595-2.305 20.22-1.86 28.197 2.876a2.009 2.009 0 0 1 .704 2.756a2.01 2.01 0 0 1-2.755.703h0Z"/>`),
		g.Group(children),
	)
}