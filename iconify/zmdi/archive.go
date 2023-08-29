package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M374 48q10 11 10 27v266q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V75q0-16 10-27l29-36Q49 0 64 0h256q15 0 25 12zM192 309l117-117h-74v-43h-86v43H75zM45 43h294l-20-22H63z"/>`),
		g.Group(children),
	)
}