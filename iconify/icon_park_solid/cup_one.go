package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCupOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M10 17h28c2 0 6 0 6 3s-4 3-6 3v18a2 2 0 0 1-2 2H12a2 2 0 0 1-2-2V23c-2 0-6 0-6-3s4-3 6-3Z"/><path d="M41 17c0-5.873-5.541-7.681-13-9V6a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2c-7.459 1.319-13 3.127-13 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCupOne0)"/>`),
		g.Group(children),
	)
}