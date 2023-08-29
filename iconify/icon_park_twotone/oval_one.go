package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOvalOne0"><ellipse cx="24" cy="24" fill="#555" stroke="#fff" stroke-width="4" rx="14" ry="20"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOvalOne0)"/>`),
		g.Group(children),
	)
}