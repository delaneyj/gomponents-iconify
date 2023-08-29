package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logstash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M8 0v80c0 26.508 21.492 48 48 48h4V48C60 21.488 38.508 0 12 0Zm64 80v48h44V80Zm0 0"/>`),
		g.Group(children),
	)
}