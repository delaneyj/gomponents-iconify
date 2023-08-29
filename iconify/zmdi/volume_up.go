package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128h85L192 21v342L85 256H0V128zm288 64q0 28-14.5 51T235 278V106q24 12 38.5 35t14.5 51zM235 5q64 15 106.5 67T384 192t-42.5 120T235 379v-44q46-14 76-53.5t30-89.5t-30-89.5T235 49V5z"/>`),
		g.Group(children),
	)
}