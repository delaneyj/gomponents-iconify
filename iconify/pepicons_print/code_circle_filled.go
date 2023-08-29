package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M4.962 12.666a.5.5 0 0 1 .706-.038l3.333 3a.5.5 0 1 1-.669.744l-3.333-3a.5.5 0 0 1-.037-.706Z"/><path d="M9.038 9.666a.5.5 0 0 1-.037.706l-3.333 3a.5.5 0 0 1-.67-.744l3.334-3a.5.5 0 0 1 .706.038Zm12 3a.5.5 0 0 1-.037.706l-3.333 3a.5.5 0 0 1-.67-.744l3.334-3a.5.5 0 0 1 .706.038Z"/><path d="M16.962 9.666a.5.5 0 0 1 .706-.038l3.333 3a.5.5 0 0 1-.669.744l-3.333-3a.5.5 0 0 1-.037-.706Zm-2.33-2.648a.5.5 0 0 1 .35.614l-3 11a.5.5 0 0 1-.964-.264l3-11a.5.5 0 0 1 .614-.35Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}