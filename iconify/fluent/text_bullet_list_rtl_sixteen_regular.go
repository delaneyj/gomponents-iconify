package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRtlSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 4.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM14 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-1 3.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0ZM10.5 3a.5.5 0 0 1 0 1h-9a.5.5 0 0 1 0-1h9Zm.5 5a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0 0 1h9A.5.5 0 0 0 11 8Zm-.5 4a.5.5 0 0 1 0 1h-9a.5.5 0 0 1 0-1h9Z"/>`),
		g.Group(children),
	)
}