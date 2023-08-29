package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrGlasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVrGlasses0"><g fill="none" fill-rule="evenodd" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"><path fill="#fff" stroke="#fff" d="M2 12h44v24H30l-6-6l-6 6H2V12Z"/><path fill="#000" stroke="#000" d="M13 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm22 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVrGlasses0)"/>`),
		g.Group(children),
	)
}