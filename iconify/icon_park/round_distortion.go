package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundDistortion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" d="M24 44C18.4772 44 14 39.5228 14 34C14 28.4772 18.4772 24 24 24C29.5228 24 34 19.5228 34 14C34 8.47715 29.5228 4 24 4"/><path stroke-linecap="round" d="M44 24C44 29.5228 39.5228 34 34 34C28.4772 34 24 29.5228 24 24C24 18.4772 19.5228 14 14 14C8.47715 14 4 18.4772 4 24"/></g>`),
		g.Group(children),
	)
}