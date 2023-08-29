package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 3c2.669 0 5.226 2.258 4.475 5.206a17.028 17.028 0 0 1-12.269 12.27C5.258 21.225 3 18.668 3 16v-1c0-1.127.901-1.945 1.9-2.044a8.942 8.942 0 0 0 2.389-.575a1 1 0 0 1 1.072.223l1.003 1.002a11.056 11.056 0 0 0 4.242-4.242L12.604 8.36a1 1 0 0 1-.223-1.072a8.942 8.942 0 0 0 .575-2.39C13.055 3.902 13.873 3 15 3h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}