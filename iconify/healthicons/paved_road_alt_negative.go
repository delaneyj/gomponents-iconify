package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PavedRoadAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPavedRoadAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 .04L0 0l-.04 48l48 .04l.04-48ZM12.993 7.97a2 2 0 0 1 2.002-1.998l18 .015a2 2 0 0 1 1.998 2.002l-.026 32a2 2 0 0 1-2.001 1.998l-18-.015a2 2 0 0 1-1.999-2.001l.026-32ZM24.012 10a1 1 0 0 0-1.001 1l-.003 4a1 1 0 1 0 2 0l.003-4a1 1 0 0 0-1-1Zm-.01 11a1 1 0 0 0-1 1l-.003 4a1 1 0 1 0 2 0l.003-4a1 1 0 0 0-1-1Zm-1.01 12a1 1 0 0 1 2 0l-.003 4a1 1 0 0 1-2 0l.004-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPavedRoadAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}