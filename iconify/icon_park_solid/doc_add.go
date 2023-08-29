package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDocAdd0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M38 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path stroke="#000" d="M17 30h14m-14 6h7m-5-19h10m-5 5V12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDocAdd0)"/>`),
		g.Group(children),
	)
}