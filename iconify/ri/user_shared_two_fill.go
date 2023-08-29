package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSharedTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14.252V22H4a8 8 0 0 1 10-7.748ZM12 13c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm6 4v-3.5l5 4.5l-5 4.5V19h-3v-2h3Z"/>`),
		g.Group(children),
	)
}