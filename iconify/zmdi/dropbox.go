package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m126 3l87 72l88-72l126 81l-87 69l87 69l-126 82l-88-73l-87 73L0 222l87-69L0 84zm87 72L87 153l126 78l127-78zm0 172l89 73l37-25v27l-126 75l-125-75v-27l38 25z"/>`),
		g.Group(children),
	)
}