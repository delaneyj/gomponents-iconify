package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.363 11.045a3.614 3.614 0 0 1 5.084 0a3.55 3.55 0 0 1 0 5.047L17 21.5l-5.447-5.408a3.55 3.55 0 0 1 0-5.047a3.614 3.614 0 0 1 5.084 0l.363.36l.363-.36Zm1.88-6.288a5.986 5.986 0 0 1 1.689 3.338A5.619 5.619 0 0 0 17 8.808a5.617 5.617 0 0 0-6.856.818a5.55 5.55 0 0 0-.178 7.701l.178.185l2.421 2.404L11 21.485l-8.48-8.492A6 6 0 0 1 11 4.529a5.998 5.998 0 0 1 8.242.228Z"/>`),
		g.Group(children),
	)
}