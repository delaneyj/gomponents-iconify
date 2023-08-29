package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotationVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 28L34 24L30 28"/><path d="M33.1679 16C31.6248 8.93638 28.1006 4 24 4C18.4772 4 14 12.9543 14 24C14 35.0457 18.4772 44 24 44C29.5228 44 34 35.0457 34 24"/></g>`),
		g.Group(children),
	)
}