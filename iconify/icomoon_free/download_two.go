package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 8h-2.5L8 11.5L4.5 8H2l-2 4v1h16v-1l-2-4zM0 14h16v1H0v-1zm9-9V1H7v4H3.5L8 9.5L12.5 5H9z"/>`),
		g.Group(children),
	)
}