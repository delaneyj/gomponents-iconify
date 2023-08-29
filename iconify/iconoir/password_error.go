package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.121 20.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L15.12 16.12m2.122 2.122l2.121 2.121M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6m1-4.99l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/>`),
		g.Group(children),
	)
}