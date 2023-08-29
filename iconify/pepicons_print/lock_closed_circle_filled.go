package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M10.5 15.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 10.5h-5A3.5 3.5 0 0 0 6.5 14v3a3.5 3.5 0 0 0 3.5 3.5h5a3.5 3.5 0 0 0 3.5-3.5v-3a3.5 3.5 0 0 0-3.5-3.5ZM7.5 14a2.5 2.5 0 0 1 2.5-2.5h5a2.5 2.5 0 0 1 2.5 2.5v3a2.5 2.5 0 0 1-2.5 2.5h-5A2.5 2.5 0 0 1 7.5 17v-3Z" clip-rule="evenodd"/><path d="M10 11a.5.5 0 0 1-1 0V7.5a3.5 3.5 0 1 1 7 0V11a.5.5 0 0 1-1 0V7.5a2.5 2.5 0 0 0-5 0V11Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}