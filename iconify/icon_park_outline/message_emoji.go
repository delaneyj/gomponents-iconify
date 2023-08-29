package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageEmoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 7H4v30h7v5l10-5h23V7Zm-13 9v1m-14-1v1"/><path d="M31 25s-2 4-7 4s-7-4-7-4"/></g>`),
		g.Group(children),
	)
}