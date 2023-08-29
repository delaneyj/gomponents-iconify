package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneShower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7c-2.76 0-5 2.24-5 5h10c0-2.76-2.24-5-5-5z" opacity=".3"/><circle cx="8" cy="20" r="1" fill="currentColor"/><circle cx="16" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M13 5.08V3h-2v2.08C7.61 5.57 5 8.47 5 12v2h14v-2c0-3.53-2.61-6.43-6-6.92zM7 12c0-2.76 2.24-5 5-5s5 2.24 5 5H7z"/><circle cx="16" cy="20" r="1" fill="currentColor"/><circle cx="12" cy="17" r="1" fill="currentColor"/><circle cx="8" cy="17" r="1" fill="currentColor"/><circle cx="12" cy="20" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}