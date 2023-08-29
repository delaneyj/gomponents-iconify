package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m192 356l157-123l35 27l-192 149L0 260l35-27zm0-55L35 179L0 152L192 3l192 149l-35 27z"/>`),
		g.Group(children),
	)
}