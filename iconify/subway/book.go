package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 124.7L256 18L0 124.7l256 106.7l256-106.7zM256 274l-144.9-67.6L0 252.7l256 106.7l256-106.7l-111.1-46.3L256 274zm0 128l-139.6-69.8L0 380.7l256 106.7l256-106.7l-116.4-48.5L256 402z"/>`),
		g.Group(children),
	)
}