package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCrownThree0"><g fill="#555" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 42h22l6-21l-10 5l-7-14l-7 14l-10-5l6 21Z"/><circle cx="7" cy="18" r="3"/><circle cx="24" cy="9" r="3"/><circle cx="41" cy="18" r="3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCrownThree0)"/>`),
		g.Group(children),
	)
}