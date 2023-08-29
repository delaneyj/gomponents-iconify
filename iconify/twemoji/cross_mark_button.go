package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossMarkButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#77B255" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="m21.529 18.006l8.238-8.238a2.498 2.498 0 0 0 0-3.535a2.498 2.498 0 0 0-3.535 0l-8.238 8.238l-8.238-8.238a2.498 2.498 0 0 0-3.535 0a2.498 2.498 0 0 0 0 3.535l8.238 8.238l-8.258 8.258a2.498 2.498 0 0 0 1.768 4.267c.64 0 1.28-.244 1.768-.732l8.258-8.259l8.238 8.238c.488.488 1.128.732 1.768.732s1.279-.244 1.768-.732a2.498 2.498 0 0 0 0-3.535l-8.24-8.237z"/>`),
		g.Group(children),
	)
}