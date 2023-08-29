package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2h2ZM9 4.64V18.5H1.3L9 4.64Zm6 0l7.7 13.86H15V4.64ZM4.7 16.5H7v-4.14L4.7 16.5ZM17 12.36v4.14h2.3L17 12.36Z"/>`),
		g.Group(children),
	)
}