package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#673AB7" d="M16.5 18L0 42h33z"/><path fill="#9575CD" d="M33.6 24L19.2 42H48z"/><path fill="#40C4FF" d="M42.9 6.3C43.6 7.4 44 8.6 44 10c0 3.9-3.1 7-7 7c-.7 0-1.3-.1-1.9-.3c1.2 2 3.4 3.3 5.9 3.3c3.9 0 7-3.1 7-7c0-3.2-2.1-5.9-5.1-6.7z"/>`),
		g.Group(children),
	)
}