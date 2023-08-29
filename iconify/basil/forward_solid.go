package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.554 16.06a.5.5 0 0 0 .656.68l2.499-1.153c2.1-.97 4.393-1.413 6.681-1.309c.027.735.068 1.47.121 2.204l.069.938a.754.754 0 0 0 1.158.581a19.55 19.55 0 0 0 5.351-5.068l.46-.64a.5.5 0 0 0 0-.584l-.46-.64A19.55 19.55 0 0 0 13.738 6a.754.754 0 0 0-1.158.58l-.069.94a61.91 61.91 0 0 0-.108 1.89h-.644a9.5 9.5 0 0 0-8.475 5.209l-.73 1.442Z"/>`),
		g.Group(children),
	)
}