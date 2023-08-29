package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagekit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#212121" d="M0 0v320h128v-48H48V45.333h160V224h-80v48h128V0z"/>`),
		g.Group(children),
	)
}