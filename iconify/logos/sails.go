package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#28A3B2" d="M24.34 255.334S-73.092 100.291 127.341.457v254.877h-103m134.032 0V97.36s31.99 52.206 97.003 157.975h-97.003"/>`),
		g.Group(children),
	)
}