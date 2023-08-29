package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacePowder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFacePowder0"><g fill="none"><ellipse cx="24" cy="30" fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="16" ry="6"/><ellipse cx="24" cy="14" fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="16" ry="10"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m23 10l-5 3m12 1l-5 3m15 21c0 3.314-7.163 6-16 6S8 41.314 8 38m32 0v-8M8 38v-8"/><ellipse cx="24" cy="30" fill="#fff" rx="7" ry="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFacePowder0)"/>`),
		g.Group(children),
	)
}