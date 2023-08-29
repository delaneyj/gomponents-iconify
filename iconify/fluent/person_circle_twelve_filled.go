package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircleTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 6A5 5 0 1 1 1 6a5 5 0 0 1 10 0Zm-6.783.264A.84.84 0 0 0 4 6.81v.16c0 .817.817 1.53 2 1.53s2-.713 2-1.53v-.16a.84.84 0 0 0-.217-.546A.753.753 0 0 0 7.225 6h-2.45a.753.753 0 0 0-.558.264ZM6 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}