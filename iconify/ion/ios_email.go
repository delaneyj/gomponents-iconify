package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosEmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M448 384V141.8l-131.1 99.8L385 319l-2 2-78.9-69.6L256 288l-48.1-36.6L129 321l-2-2 68-77.4L64 142v242z" fill="currentColor"/><path d="M439.7 128H72l184 139.9z" fill="currentColor"/>`),
		g.Group(children),
	)
}