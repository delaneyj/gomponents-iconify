package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.058 1.043C16.744-2.841 6.018 4.682 6.104 14.38a.459.459 0 0 1-.45.451a.459.459 0 0 1-.388-.221a10.387 10.387 0 0 1-.412-7.634a.42.42 0 0 0-.712-.412a10.284 10.284 0 0 0-2.784 7.033A10.284 10.284 0 0 0 11.76 23.999c14.635-.332 11.257-19.491 8.298-22.956z"/>`),
		g.Group(children),
	)
}