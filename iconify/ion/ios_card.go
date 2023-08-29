package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M480 128c0-17.6-14.4-32-32-32H64c-17.6 0-32 14.4-32 32v48h448v-48z" fill="currentColor"/><path d="M32 384c0 17.6 14.4 32 32 32h384c17.6 0 32-14.4 32-32V224H32v160zm142-48h164c7.7 0 14 6.3 14 14s-6.3 14-14 14H174c-7.7 0-14-6.3-14-14s6.3-14 14-14zm-72 0h12c7.7 0 14 6.3 14 14s-6.3 14-14 14h-12c-7.7 0-14-6.3-14-14s6.3-14 14-14z" fill="currentColor"/>`),
		g.Group(children),
	)
}