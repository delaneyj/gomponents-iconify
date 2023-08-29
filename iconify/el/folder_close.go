package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.033H0v144.653h1200V225.733H591.65L507.935 117.7H249.17zM0 410.67v671.63h1200V410.67H0z"/>`),
		g.Group(children),
	)
}