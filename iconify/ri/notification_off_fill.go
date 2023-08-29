package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.587 20H4a.5.5 0 0 1-.4-.8l.4-.533V10c0-1.33.325-2.584.899-3.687L1.395 2.808l1.414-1.414l19.799 19.798l-1.415 1.415L18.587 20ZM20 15.786L7.561 3.345A8 8 0 0 1 20 10v5.786ZM9.501 21h5a2.5 2.5 0 0 1-5 0Z"/>`),
		g.Group(children),
	)
}