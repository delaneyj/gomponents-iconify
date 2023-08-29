package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18.977 14.79a6.907 6.907 0 1 0-11.573-3.159c.095.369.01.768-.258 1.037L3.433 16.38a1.48 1.48 0 0 0-.424 1.21l.232 2.089c.025.223.125.43.283.589l.208.208a.987.987 0 0 0 .589.283l2.089.232a1.48 1.48 0 0 0 1.21-.424l.71-.71l1.06-1.061l1.942-1.942c.27-.27.668-.353 1.037-.258a6.904 6.904 0 0 0 6.608-1.806Z" opacity=".5"/><path d="M15.414 8.586a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828ZM6.583 18.13l1.746 1.727l1.06-1.061l-1.751-1.733a.75.75 0 1 0-1.055 1.066Z"/></g>`),
		g.Group(children),
	)
}