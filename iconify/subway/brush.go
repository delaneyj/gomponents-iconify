package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M79.1 372.5C59.4 431.9 49.5 451.7 0 451.7c168.2 69.3 197.9-49.5 197.9-49.5l-49.5-49.5s-49.5-9.8-69.3 19.8zM504.6 46.1c-14-14-38-3.9-38-3.9c-35 7-229.1 201.8-229.1 201.8l69.3 69.3S501.6 119.2 508.6 84.2c-.1-.1 10-24.2-4-38.1zM168.2 333l49.5 49.5L287 333l-69.3-69.3l-49.5 69.3z"/>`),
		g.Group(children),
	)
}