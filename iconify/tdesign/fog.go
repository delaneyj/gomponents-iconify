package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.1 4C8.634 4 6.7 5.916 6.7 8.2c0 .262.025.518.073.765l.229 1.19H5.79C4.839 10.155 4 11.018 4 12v1H2v-1c0-1.657 1.124-3.185 2.701-3.68L4.7 8.2C4.7 4.74 7.601 2 11.1 2c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.554 5.554 0 0 1 2.01.657C20.942 7.986 22 10.086 22 12v1h-2v-1c0-1.298-.756-2.669-1.741-3.216a3.612 3.612 0 0 0-1.818-.45l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.607 4.607 0 0 0 11.1 4ZM2 14h6v2H2v-2Zm8 0h12v2H10v-2Zm7 4h5v2h-5v-2ZM2 20v-2h13v2H2Z"/>`),
		g.Group(children),
	)
}