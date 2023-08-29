package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 206.2L326.8 21v108.9C-10.9 129.9 0 391.4 0 500.3c76.2-217.9 174.3-217.9 326.8-217.9v108.9L512 206.2z"/>`),
		g.Group(children),
	)
}