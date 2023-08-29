package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrotliOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M374 250.698H95.623V96.433H374v154.265zM44.174 469.588h36.015V344.562l-36.015-36.015v161.04zM327.306 48.07H44.174V208.9l36.015 37.73V80.999h247.117V48.071zM95.623 346.386V512H374V346.386H95.623zM28.74 192.732V32.638h262.11V0H0v162.623l28.74 30.11zm0 100.381L0 264.373v162.802h28.74V293.113z"/>`),
		g.Group(children),
	)
}