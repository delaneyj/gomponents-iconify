package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Game(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGame0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M38.512 9.536A20.429 20.429 0 0 0 24.5 4C13.178 4 4 13.178 4 24.5S13.178 45 24.5 45a20.435 20.435 0 0 0 14.405-5.914L24 24L38.512 9.536Z"/><path fill="#555" d="M40 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M17 13v8m-4-4h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGame0)"/>`),
		g.Group(children),
	)
}