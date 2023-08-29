package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVideoOne0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 10a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v9l8-6v23l-8-6v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V10Z"/><circle cx="17" cy="21" r="5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVideoOne0)"/>`),
		g.Group(children),
	)
}