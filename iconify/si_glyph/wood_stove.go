package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WoodStove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.965 3.027V.072h-.949v2.955H3v11.957h2.016L5 15.978h.969v-.994h6.047v.994h.969v-.994H15V3.027h-3.035zm1.05 7.02V12c0 1.002-.074 1.055-1.016 1.055H6.044c-.94 0-1.045-.053-1.045-1.055V6c0-1.001.104-1.033 1.045-1.033H12c.941 0 1.016.032 1.016 1.033v2.953h1.016v1.094h-1.017z"/><path d="M11.974 7.031c0-.539-.401-.975-.896-.975H6.981c-.494 0-.897.436-.897.975v3.979c0 .537.403.975.897.975h4.097c.495 0 .896-.438.896-.975V7.031zM8 10H7V8h1v2zm2 0H8.964V8H10v2z"/></g>`),
		g.Group(children),
	)
}