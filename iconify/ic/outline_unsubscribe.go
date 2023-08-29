package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineUnsubscribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.99 14.04V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h10.05c.28 1.92 2.1 3.35 4.18 2.93c1.34-.27 2.43-1.37 2.7-2.71c.25-1.24-.16-2.39-.94-3.18zm-2-9.04L12 8.5L5 5h13.99zm-3.64 10H5V7l7 3.5L19 7v6.05c-.16-.02-.33-.05-.5-.05c-1.39 0-2.59.82-3.15 2zm5.15 2h-4v-1h4v1z"/>`),
		g.Group(children),
	)
}