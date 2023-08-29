package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShareOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M35 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM13 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="m30 13.575l-12.66 7.67m-.002 5.319l13.34 7.883"/><path fill="#555" d="M35 32a5 5 0 1 1 0 10a5 5 0 0 1 0-10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShareOne0)"/>`),
		g.Group(children),
	)
}