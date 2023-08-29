package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorReset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 235q0 14-3 28L112 79q14-19 28.5-37.5T163 14l8-10q5 6 13.5 16.5t30.5 40t39 56.5t31 60.5t14 57.5zm-19 66l58 59l-27 27l-56-56q-36 32-84 32q-53 0-90.5-37.5T43 235q0-35 28-88L0 76l27-28l154 155z"/>`),
		g.Group(children),
	)
}