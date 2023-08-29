package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 344V173h43v171h-43zm0 85v-42h43v42h-43zM0 429L427 3v128h-86v298H0z"/>`),
		g.Group(children),
	)
}