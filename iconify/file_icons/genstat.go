package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Genstat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0h512v395.201c-170.667 16.732-228.671-316.7-328.347-316.7C99.747 78.501 67.705 266.416 0 288.674V0zm0 512h512v-84.822c-175.872 18.777-264.66-315.946-329.886-316.7c-57.925 0-89.678 189.848-182.114 210.173V512z"/>`),
		g.Group(children),
	)
}