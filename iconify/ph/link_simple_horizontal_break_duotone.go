package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleHorizontalBreakDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 128a48 48 0 0 1-48 48H64a48 48 0 0 1-48-48a48 48 0 0 1 48-48h128a48 48 0 0 1 48 48Z" opacity=".2"/><path d="M24 128a40 40 0 0 0 40 40h40a8 8 0 0 1 0 16H64a56 56 0 0 1 0-112h40a8 8 0 0 1 0 16H64a40 40 0 0 0-40 40Zm168-56h-40a8 8 0 0 0 0 16h40a40 40 0 0 1 0 80h-40a8 8 0 0 0 0 16h40a56 56 0 0 0 0-112Z"/></g>`),
		g.Group(children),
	)
}