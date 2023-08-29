package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hibernate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m29.246 3.766l23.168 40.129l-23.18 40.19l-23.156-40.19Zm69.508 120.468L75.586 84.105l23.18-40.19l23.156 40.19ZM75.594 3.766H29.258L52.43 43.898h46.35ZM52.406 124.23H98.75L75.594 84.102H29.219Zm0 0"/>`),
		g.Group(children),
	)
}