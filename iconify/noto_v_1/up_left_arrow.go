package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m29.75 29.82l53.73 8.28l-14.17 14.18l28.94 28.94l-16.96 16.96l-28.94-28.94l-14.33 14.33l-8.27-53.75z"/>`),
		g.Group(children),
	)
}