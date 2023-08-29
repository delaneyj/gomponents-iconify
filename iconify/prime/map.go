package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.9 4.09a1.75 1.75 0 0 0-1.66-.15l-3.57 1.53l-4.75-2a1.51 1.51 0 0 0-1.18 0L4.38 5.31a1.88 1.88 0 0 0-1.13 1.75v11.25a1.91 1.91 0 0 0 .85 1.6a1.75 1.75 0 0 0 1.66.15l3.57-1.53l4.75 2a1.45 1.45 0 0 0 .59.12a1.52 1.52 0 0 0 .59-.12l4.36-1.87a1.88 1.88 0 0 0 1.13-1.75V5.69a1.91 1.91 0 0 0-.85-1.6Zm-9.82 1l3.84 1.64v12.13l-3.84-1.64ZM5.17 18.68a.25.25 0 0 1-.25 0a.4.4 0 0 1-.17-.35V7.06A.39.39 0 0 1 5 6.69l3.58-1.55v12.08Zm14.08-1.74a.39.39 0 0 1-.22.37l-3.61 1.55V6.78l3.41-1.46a.25.25 0 0 1 .25 0a.4.4 0 0 1 .17.35Z"/>`),
		g.Group(children),
	)
}