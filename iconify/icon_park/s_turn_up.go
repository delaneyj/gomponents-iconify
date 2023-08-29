package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 31L24 35C24 39 21 42 17 42C13 42 10 39 10 35L10 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 42L38 13C38 8.99999 35 5.99999 31 5.99999C27 5.99999 24 8.99999 24 13L24 31"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 37L38 42L43 37"/><circle cx="10" cy="11" r="5" fill="#2F88FF" transform="rotate(-180 10 11)"/></g>`),
		g.Group(children),
	)
}