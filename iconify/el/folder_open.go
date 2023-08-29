package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.032H0v609.972l123.572-465.318H993.75V225.732h-402.1L507.935 117.7H249.17zM0 1016.612v65.688h993.75L1200 410.669H175.134L0 1016.612z"/>`),
		g.Group(children),
	)
}