package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneToMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOneToMany0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 7H6a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M11 20.58L15 18v12m14 0V18l8 12V18"/><path stroke="#000" stroke-linecap="round" d="M22 20v1m0 6v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOneToMany0)"/>`),
		g.Group(children),
	)
}