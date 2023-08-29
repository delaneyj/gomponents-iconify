package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 7a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}