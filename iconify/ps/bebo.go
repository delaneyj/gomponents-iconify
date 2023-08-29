package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bebo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M231 462h-46q-77 0-130.5-54T1 278V48q0-20 13.5-33T47 2t32.5 13.5T93 48v230q0 38 27 65t65 27h46q38 0 65-23.5t27-56.5t-14.5-57t-31.5-24h-92q-19 0-32.5-13.5T139 163t13.5-32.5T185 117h92q57 0 97.5 50.5T415 290q0 71-54 121.5T231 462z"/>`),
		g.Group(children),
	)
}