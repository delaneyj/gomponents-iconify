package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoreHole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0c-17.7 0-32 14.3-32 32v264.6c-19.1 11.1-32 31.7-32 55.4c0 35.3 28.7 64 64 64s64-28.7 64-64c0-23.7-12.9-44.4-32-55.4V32c0-17.7-14.3-32-32-32zM48 128c-26.5 0-48 21.5-48 48v288c0 26.5 21.5 48 48 48h416c26.5 0 48-21.5 48-48V176c0-26.5-21.5-48-48-48h-80c-17.7 0-32 14.3-32 32v192c0 53-43 96-96 96s-96-43-96-96V160c0-17.7-14.3-32-32-32H48z"/>`),
		g.Group(children),
	)
}