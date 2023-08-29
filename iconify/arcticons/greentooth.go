package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greentooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.425 24l-.017 19.5l10.416-9.779L11.5 13.788m13.704 7.614l7.62-7.123L22.408 4.5v14.29M11.5 34.212L22.425 24m10.399 9.721l3.676 3.433M22.408 18.79l2.796 2.612"/>`),
		g.Group(children),
	)
}