package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 6L6 12"/><path d="M42 36L36 42"/><path d="M22 6L6 22"/><path d="M32 6L6 32"/><path d="M42 6L6 42"/><path d="M42 16L16 42"/><path d="M42 26L26 42"/></g>`),
		g.Group(children),
	)
}