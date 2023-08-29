package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 128a48 48 0 0 1-48 48H48a32 32 0 0 1 0-64h32V80h112a48 48 0 0 1 48 48Z" opacity=".2"/><path d="M248 128a56.06 56.06 0 0 1-56 56H48a40 40 0 0 1 0-80h144a24 24 0 0 1 0 48H80a8 8 0 0 1 0-16h112a8 8 0 0 0 0-16H48a24 24 0 0 0 0 48h144a40 40 0 0 0 0-80H80a8 8 0 0 1 0-16h112a56.06 56.06 0 0 1 56 56Z"/></g>`),
		g.Group(children),
	)
}