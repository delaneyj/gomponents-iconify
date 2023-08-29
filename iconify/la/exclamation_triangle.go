package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3.219l-.875 1.5l-12 20.781l-.844 1.5H29.72l-.844-1.5l-12-20.781zm0 4L26.25 25H5.75zM15 14v6h2v-6zm0 7v2h2v-2z"/>`),
		g.Group(children),
	)
}