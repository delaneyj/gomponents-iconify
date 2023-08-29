package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightRazor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="6" x="3.609" y="36.534" rx="2" transform="rotate(-10 3.61 36.534)"/><path d="m44 40l-4-4M8 4l18.385 18.385l-4.243 4.242L9.414 13.9c-2.828-2.83-2.828-4.244-2.828-5.658C6.586 6.828 8 4 8 4Zm0 0l18 18l9 9"/></g>`),
		g.Group(children),
	)
}