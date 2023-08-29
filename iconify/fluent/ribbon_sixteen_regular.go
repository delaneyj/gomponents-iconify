package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1a5 5 0 0 0-3 9v4.5a.5.5 0 0 0 .757.429L8 13.583l2.243 1.346A.5.5 0 0 0 11 14.5V10a5 5 0 0 0-3-9ZM4 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm6 4.584v3.033L8.257 12.57a.5.5 0 0 0-.514 0L6 13.617v-3.033A4.983 4.983 0 0 0 8 11c.711 0 1.388-.148 2-.416Z"/>`),
		g.Group(children),
	)
}