package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiPictureCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" fill="#2F88FF" rx="2"/><path d="M20 42H28"/><path d="M34 42H36"/><path d="M4 42H6"/><path d="M42 42H44"/><path d="M12 42H14"/></g>`),
		g.Group(children),
	)
}