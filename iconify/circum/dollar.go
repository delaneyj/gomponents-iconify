package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.934Zm0-18.868A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.066Z"/><path fill="currentColor" d="M14.5 13.5a2.006 2.006 0 0 1-2 2v1.01a.5.5 0 0 1-1 0V15.5h-1.25a.5.5 0 0 1 0-1h2.25a1 1 0 0 0 0-2h-1a2 2 0 0 1 0-4V7.49a.5.5 0 0 1 1 0V8.5h1.25a.5.5 0 0 1 0 1H11.5a1 1 0 0 0 0 2h1a2.006 2.006 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}