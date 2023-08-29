package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotonePercentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.501 3.5l-15 15.001l1.996 1.996l15-15z"/><circle cx="7" cy="7" r="1" fill="currentColor" fill-rule="evenodd" opacity=".3"/><circle cx="17" cy="17" r="1" fill="currentColor" fill-rule="evenodd" opacity=".3"/><path fill="currentColor" d="M17.003 14a3 3 0 1 1-.006 6a3 3 0 0 1 .006-6zM17 16a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM7.003 4a3 3 0 1 1-.006 6a3 3 0 0 1 .006-6zM7 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}