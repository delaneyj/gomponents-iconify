package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 102q0-29-20-49t-48-20H70q-28 0-48 20T2 102v185q0 29 20 49t48 20h324q28 0 48-20t20-49V102zM186 273V99l132 87z"/>`),
		g.Group(children),
	)
}