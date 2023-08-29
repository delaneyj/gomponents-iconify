package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ACane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="svgIDb"><g fill="none"><g clip-path="url(#svgIDa)"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19.557 44.768L33.641 18.28c1.174-2.207 3.812-9.299-3.252-13.055C23.326 1.47 19.157 7.181 17.75 9.83"/></g><defs><clipPath id="svgIDa"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#svgIDb)"/>`),
		g.Group(children),
	)
}