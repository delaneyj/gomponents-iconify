package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 336v-48h432v48H0zm0-120v-48h432v48H0zM0 48h432v48H0V48z"/>`),
		g.Group(children),
	)
}