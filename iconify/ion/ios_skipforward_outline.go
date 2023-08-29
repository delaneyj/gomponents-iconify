package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosSkipforwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M337 96v141.8L96 96v320l241-141.8V416h79V96h-79zm-8.1 164.4L112 388V124l216.9 127.6 7.6 4.4-7.6 4.4zM400 400h-47V112h47v288z" fill="currentColor"/>`),
		g.Group(children),
	)
}