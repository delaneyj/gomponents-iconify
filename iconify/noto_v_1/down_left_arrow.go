package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m29.82 98.25l8.28-53.74l14.18 14.18l28.94-28.94l16.96 16.97l-28.94 28.93l14.32 14.33l-53.74 8.27z"/>`),
		g.Group(children),
	)
}