package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSlave0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 38v-7.5M10 38V6a2 2 0 0 1 2-2h24a2 2 0 0 1 2 2v1"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25 13h-1a6 6 0 0 0-6 6v0a6 6 0 0 0 6 6h1m8-12h1a6 6 0 0 1 6 6v0a6 6 0 0 1-6 6h-1m-8-6h8"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 38a6 6 0 0 1 6-6h16a6 6 0 0 1 0 12H16a6 6 0 0 1-6-6Z"/><circle cx="32" cy="38" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSlave0)"/>`),
		g.Group(children),
	)
}