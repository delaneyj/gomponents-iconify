package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualRealityGlasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVirtualRealityGlasses0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M5 16h38a1 1 0 0 1 1 1v22a1 1 0 0 1-1 1H30l-5.992-5.999L18 40H5a1 1 0 0 1-1-1V17a1 1 0 0 1 1-1Z"/><path fill="#000" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm20 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path fill="#fff" fill-rule="evenodd" d="M6 10h36H6Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 10h36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVirtualRealityGlasses0)"/>`),
		g.Group(children),
	)
}