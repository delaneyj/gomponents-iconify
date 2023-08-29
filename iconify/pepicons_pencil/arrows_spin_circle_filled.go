package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsSpinCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowsSpinCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8.254 17.596a.5.5 0 0 1 .707-.707A5.5 5.5 0 0 0 18.35 13a.5.5 0 0 1 .999.001a6.5 6.5 0 0 1-11.096 4.596Z"/><path d="M16.131 15.416a.5.5 0 0 1-.555-.832l3-2a.5.5 0 1 1 .555.832l-3 2Z"/><path d="M21.266 15.723a.5.5 0 1 1-.832.554l-2-3a.5.5 0 0 1 .832-.554l2 3Zm-3.912-7.518a.5.5 0 0 1-.708.707a5.5 5.5 0 0 0-9.389 3.89a.5.5 0 0 1-1-.001a6.5 6.5 0 0 1 11.097-4.596Z"/><path d="M9.476 10.385a.5.5 0 0 1 .555.832l-3 2a.5.5 0 1 1-.555-.832l3-2Z"/><path d="M4.341 10.078a.5.5 0 1 1 .832-.554l2 3a.5.5 0 0 1-.832.554l-2-3Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowsSpinCircleFilled0)"/></g>`),
		g.Group(children),
	)
}