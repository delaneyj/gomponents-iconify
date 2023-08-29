package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M18 10H8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2ZM8 9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3H8Z" clip-rule="evenodd"/><path d="M13 15a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path fill-rule="evenodd" d="M17.447 4.474a.5.5 0 0 1-.223.67l-10 5a.5.5 0 1 1-.448-.894l10-5a.5.5 0 0 1 .671.224ZM14 13.65a.35.35 0 0 1 .35-.35h3.3a.35.35 0 1 1 0 .7h-3.3a.35.35 0 0 1-.35-.35Zm0 1.5a.35.35 0 0 1 .35-.35h3.3a.35.35 0 1 1 0 .7h-3.3a.35.35 0 0 1-.35-.35Zm0 1.5a.35.35 0 0 1 .35-.35h3.3a.35.35 0 1 1 0 .7h-3.3a.35.35 0 0 1-.35-.35Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}