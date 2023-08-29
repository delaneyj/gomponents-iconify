package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperAirplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M32 272l128 48 16 160 80-112 112 112L480 32 32 272zm318.7 145.4L256 320l128-176-192 153.8-82.6-31 322-172.5-80.7 323.1z" fill="currentColor"/>`),
		g.Group(children),
	)
}