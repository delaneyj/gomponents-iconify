package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M3.5 2A1 1 0 0 0 3 3.719l20 20a1 1 0 1 0 1.406-1.407L17 14.907V3.312c0-1.265-1.105-1.582-1.969-.718L9.812 7.719L4.407 2.312A1 1 0 0 0 3.594 2A1 1 0 0 0 3.5 2zM5 9.063c-.551 0-1 .448-1 1v6c0 .55.449 1 1 1h3.438L15 23.468c1 1 2 .488 2-.875V20.03L6.031 9.063H5z"/>`),
		g.Group(children),
	)
}