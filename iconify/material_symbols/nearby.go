package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.4L7.6 12L12 7.6l4.4 4.4l-4.4 4.4Zm1.4 4.975q-.275.275-.65.425t-.75.15q-.375 0-.75-.15t-.65-.425L2.625 13.4q-.275-.275-.425-.65T2.05 12q0-.375.15-.75t.425-.65l7.95-7.95q.3-.3.663-.45T12 2.05q.4 0 .762.15t.663.45l7.95 7.95q.275.275.425.65t.15.75q0 .375-.15.75t-.425.65L13.4 21.375ZM12 19.2l7.2-7.2L12 4.8L4.8 12l7.2 7.2Z"/>`),
		g.Group(children),
	)
}