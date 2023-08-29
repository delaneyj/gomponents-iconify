package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphaAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12H5zm9.5 4h-4a.5.5 0 0 1-.416-.777L13.566 10H10.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 .416.777L11.434 15H14.5a.5.5 0 0 1 0 1zm1.447-9.724l-3-6a.5.5 0 0 0-.894 0l-3 6a.5.5 0 0 0 .895.447l.862-1.724h3.382l.862 1.724a.5.5 0 0 0 .895-.447zM11.309 4L12.5 1.618L13.691 4h-2.382z"/>`),
		g.Group(children),
	)
}