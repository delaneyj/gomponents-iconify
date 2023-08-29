package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircleTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 14C3.5 8.201 8.201 3.5 14 3.5S24.5 8.201 24.5 14S19.799 24.5 14 24.5S3.5 19.799 3.5 14ZM14 2C7.373 2 2 7.373 2 14s5.373 12 12 12s12-5.373 12-12S20.627 2 14 2Zm3.25 7.25a3.25 3.25 0 1 1-6.5 0a3.25 3.25 0 0 1 6.5 0ZM14 21.5c3.314 0 6-2.143 6-5.357C20 14.959 19.04 14 17.857 14h-7.714C8.959 14 8 14.96 8 16.143c0 3.214 2.686 5.357 6 5.357Z"/>`),
		g.Group(children),
	)
}