package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsSpinCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M8.254 17.596a.5.5 0 0 1 .707-.707A5.5 5.5 0 0 0 18.35 13a.5.5 0 0 1 .999.001a6.5 6.5 0 0 1-11.096 4.596Z"/><path d="M16.131 15.416a.5.5 0 0 1-.555-.832l3-2a.5.5 0 1 1 .555.832l-3 2Z"/><path d="M21.266 15.723a.5.5 0 1 1-.832.554l-2-3a.5.5 0 0 1 .832-.554l2 3Zm-3.912-7.518a.5.5 0 0 1-.708.707a5.5 5.5 0 0 0-9.389 3.89a.5.5 0 0 1-1-.001a6.5 6.5 0 0 1 11.097-4.596Z"/><path d="M9.476 10.385a.5.5 0 0 1 .555.832l-3 2a.5.5 0 1 1-.555-.832l3-2Z"/><path d="M4.341 10.078a.5.5 0 1 1 .832-.554l2 3a.5.5 0 0 1-.832.554l-2-3Zm-.191-5.2a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}