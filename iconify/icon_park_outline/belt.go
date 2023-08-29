package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Belt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m10.565 30.364l-6.364 6.364l7.071 7.071l6.364-6.364m4.243-18.385l14.85-14.849l7.07 7.071l-14.849 14.85"/><path d="M9.859 29.657a2 2 0 0 1 0-2.829l8.485-8.485a2 2 0 0 1 2.828 0l8.485 8.485a2 2 0 0 1 0 2.829l-8.485 8.485a2 2 0 0 1-2.828 0l-8.485-8.485Zm15.556-7.071l-7.071 7.07M31.779 9.15l7.07 7.072M26.828 14.1l7.071 7.072m-18.384 0L26.83 32.485"/></g>`),
		g.Group(children),
	)
}