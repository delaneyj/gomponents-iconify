package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logstash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00bfb3" d="M72 128h44V80H72Zm0 0"/><path fill="#fec514" d="M12 0H8v80h52V48C60 21.488 38.508 0 12 0"/><path fill="#343741" d="M8 80c0 26.508 21.492 48 48 48h4V80Zm0 0"/>`),
		g.Group(children),
	)
}