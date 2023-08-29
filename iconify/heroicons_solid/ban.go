package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.477 14.89A6 6 0 0 1 5.11 6.524l8.367 8.368Zm1.414-1.414L6.524 5.11a6 6 0 0 1 8.367 8.367ZM18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}