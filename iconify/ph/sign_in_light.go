package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignInLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m140.24 132.24l-40 40a6 6 0 0 1-8.48-8.48L121.51 134H24a6 6 0 0 1 0-12h97.51L91.76 92.24a6 6 0 0 1 8.48-8.48l40 40a6 6 0 0 1 0 8.48ZM192 34h-56a6 6 0 0 0 0 12h56a2 2 0 0 1 2 2v160a2 2 0 0 1-2 2h-56a6 6 0 0 0 0 12h56a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14Z"/>`),
		g.Group(children),
	)
}