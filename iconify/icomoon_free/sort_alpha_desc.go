package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphaDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12H5zm9.5-5h-4a.5.5 0 0 1-.416-.777L13.566 1H10.5a.5.5 0 0 1 0-1h4a.502.502 0 0 1 .416.777L11.434 6H14.5a.5.5 0 0 1 0 1zm1.447 8.276l-3-6a.5.5 0 0 0-.894 0l-3 6a.5.5 0 0 0 .895.447l.862-1.724h3.382l.862 1.724a.5.5 0 0 0 .895-.447zM11.309 13l1.191-2.382L13.691 13h-2.382z"/>`),
		g.Group(children),
	)
}