package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTKeyOne0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="15" cy="33" r="8" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m29 16l6.5 6M20 26L37 7m-2 4l7 6.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTKeyOne0)"/>`),
		g.Group(children),
	)
}