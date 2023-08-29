package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileWord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileWord0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 4h20l10 10v28a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/><path stroke="#000" stroke-linecap="round" d="m16.008 20l3 14l5-10l5 10l3-14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileWord0)"/>`),
		g.Group(children),
	)
}