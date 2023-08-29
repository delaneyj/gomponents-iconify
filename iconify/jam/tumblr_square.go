package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.441 12.953s-.307.34-.89.34c-.584 0-.845-.34-.845-.845V9.59h1.887V7.857h-1.887V5H9.285C9.062 6.282 8.187 7.35 7 7.874V9.59h1.386v3.291c0 .457.446 2.119 2.72 2.119c1.335 0 1.888-.83 1.888-.83l-.553-1.217z"/><path d="M4 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4zm0-2h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/></g>`),
		g.Group(children),
	)
}