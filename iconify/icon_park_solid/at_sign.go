package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAtSign0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M44 24c0-11.046-8.954-20-20-20S4 12.954 4 24s8.954 20 20 20v0c4.989 0 9.55-1.827 13.054-4.847"/><path fill="#fff" d="M24 32a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path stroke-linecap="round" d="M32 24a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6m-12 1v-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAtSign0)"/>`),
		g.Group(children),
	)
}