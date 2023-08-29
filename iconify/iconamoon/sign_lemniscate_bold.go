package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignLemniscateBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/>`),
		g.Group(children),
	)
}