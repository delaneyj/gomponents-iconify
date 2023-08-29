package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneToOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOneToOne0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 7H6a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 20.58L16 18v12m15-9.42L35 18v12"/><path stroke="#000" stroke-linecap="round" d="M24 20v1m0 6v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOneToOne0)"/>`),
		g.Group(children),
	)
}