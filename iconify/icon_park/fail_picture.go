package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FailPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="34" height="34" x="7" y="7" rx="3"/><path stroke-linecap="round" d="M41 26L34.3877 21.0408C33.5914 20.4436 32.4772 20.5228 31.7735 21.2265L27 26"/><path stroke-linecap="round" d="M7 28L18 18"/><path stroke-linecap="round" d="M6 6L42 42"/></g>`),
		g.Group(children),
	)
}