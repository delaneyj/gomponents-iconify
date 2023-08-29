package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosSkipforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M337 96v141.8L96 96v320l241-141.8V416h79V96h-79z" fill="currentColor"/>`),
		g.Group(children),
	)
}