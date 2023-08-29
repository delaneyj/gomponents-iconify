package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NervousNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsNervousNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm0 2c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18ZM14 31.5a4.5 4.5 0 0 1 4.5-4.5h11a4.5 4.5 0 1 1 0 9h-11a4.5 4.5 0 0 1-4.5-4.5Zm3 1.5a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H18a1 1 0 0 1-1-1Zm1-4a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H18Zm2-8c0 2.21-1.12 4-2.5 4S15 23.21 15 21s1.12-4 2.5-4s2.5 1.79 2.5 4Zm13 0c0 2.21-1.12 4-2.5 4S28 23.21 28 21s1.12-4 2.5-4s2.5 1.79 2.5 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsNervousNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}