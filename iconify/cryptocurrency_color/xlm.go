package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xlm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#000"/><path fill="#FFF" d="m23.13 9.292l-2.4 1.224l-11.598 5.907A6.909 6.909 0 0 1 19.35 9.498l1.374-.7l.205-.105a8.439 8.439 0 0 0-13.371 7.472a1.535 1.535 0 0 1-.834 1.484l-.725.37v1.724l2.134-1.088l.691-.353l.681-.347l12.226-6.23l1.374-.699l2.84-1.447V7.856L23.13 9.292zm2.816 2.012L10.201 19.32l-1.374.7L6 21.463v1.723l2.808-1.43l2.401-1.224l11.61-5.916a6.909 6.909 0 0 1-10.229 6.93l-.085.045l-1.49.76a8.439 8.439 0 0 0 13.372-7.475a1.536 1.536 0 0 1 .833-1.483l.726-.37v-1.718z"/></g>`),
		g.Group(children),
	)
}