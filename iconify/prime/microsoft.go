package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h7.5v7.5H4Zm8.5 0H20v7.5h-7.5ZM4 12.5h7.5V20H4Zm8.5 0H20V20h-7.5Z"/>`),
		g.Group(children),
	)
}