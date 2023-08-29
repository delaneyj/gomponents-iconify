package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivityManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h15.857v15.857H5.5zm19.07 7.928l10-10l10 10l-10 10zM5.5 26.643h15.857V42.5H5.5zm21.143 0H42.5V42.5H26.643z"/>`),
		g.Group(children),
	)
}