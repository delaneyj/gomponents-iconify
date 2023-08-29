package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2c4.42 0 8 3.58 8 8s-3.58 8-8 8s-8-3.58-8-8s3.58-8 8-8zm1 4c0-.55-.45-1-1-1s-1 .45-1 1s.45 1 1 1s1-.45 1-1zm0 9V9H9v6h2z"/>`),
		g.Group(children),
	)
}