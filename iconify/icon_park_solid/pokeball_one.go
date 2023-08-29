package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PokeballOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPokeballOne0"><g fill="none"><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><circle cx="24" cy="24" r="6" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 24h14M4 24h14"/><circle cx="24" cy="24" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPokeballOne0)"/>`),
		g.Group(children),
	)
}