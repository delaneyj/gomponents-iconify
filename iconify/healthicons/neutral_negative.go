package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsNeutralNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm0 2c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Zm-6.5-17c1.38 0 2.5-1.79 2.5-4s-1.12-4-2.5-4s-2.5 1.79-2.5 4s1.12 4 2.5 4ZM33 21c0 2.21-1.12 4-2.5 4S28 23.21 28 21s1.12-4 2.5-4s2.5 1.79 2.5 4Zm-16 9a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H17Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsNeutralNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}