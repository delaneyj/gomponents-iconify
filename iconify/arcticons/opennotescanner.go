package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opennotescanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 33.75v7.8a2 2 0 0 0 1.95 2h7.8m11.7-.05h7.8a2 2 0 0 0 2-2v-7.8M18.15 4.5h-7.8a2 2 0 0 0-1.95 2v7.8m31.2-.05v-7.8a2 2 0 0 0-2-1.95h-7.8m-14.58 8.2h17.55m0 7.53H15.22m0 7.54h9.75m-9.75 7.53h5.85M26 32.84v2.46h2.4a1 1 0 0 0 .68-.29l8.1-8.11l-2.85-2.84l-8.1 8.1a1 1 0 0 0-.23.68Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.37 24.7a.74.74 0 0 0 0-1.06h0l-1.72-1.78a.77.77 0 0 0-1.06 0l-2.26 2.21l2.85 2.84Z"/>`),
		g.Group(children),
	)
}