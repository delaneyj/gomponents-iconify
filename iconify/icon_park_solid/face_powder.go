package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacePowder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFacePowder0"><g fill="none"><ellipse cx="24" cy="30" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="16" ry="6"/><ellipse cx="24" cy="14" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="16" ry="10"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m23 10l-5 3m12 1l-5 3"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 38c0 3.314-7.163 6-16 6S8 41.314 8 38m32 0v-8M8 38v-8"/><ellipse cx="24" cy="30" fill="#000" rx="7" ry="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFacePowder0)"/>`),
		g.Group(children),
	)
}