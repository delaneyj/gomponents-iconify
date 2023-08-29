package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCocosKeelingIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m52.233 39.354l.927-3l.927 3L51.66 37.5h3l-2.427 1.854zM42.5 35.831l-1.113 1.054l.111-1.549L40 35.101l1.25-.878l-.755-1.347l1.449.455l.556-1.446l.556 1.446l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.549l-1.113-1.054zm15-3l-1.113 1.054l.111-1.549L55 32.101l1.25-.878l-.755-1.347l1.449.455l.556-1.446l.556 1.446l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.549l-1.113-1.054zm-8-6l-1.113 1.054l.111-1.549L47 26.101l1.25-.878l-.755-1.347l1.449.455l.556-1.446l.556 1.446l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.549l-1.113-1.054zm0 21l-1.113 1.054l.111-1.549L47 47.101l1.25-.878l-.755-1.347l1.449.455l.556-1.446l.556 1.446l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.549l-1.113-1.054z"/><circle cx="13" cy="25" r="4" fill="#fcea2b"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="M32.98 42.596A6.596 6.596 0 1 1 36 30.13a8.395 8.395 0 1 0 0 11.74a6.652 6.652 0 0 1-3.02.726Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}