package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 416a288 288 0 1 0-576 0c0 118.144 94.528 272.128 288 456.576C705.472 688.128 800 534.144 800 416zM512 960C277.312 746.688 160 565.312 160 416a352 352 0 0 1 704 0c0 149.312-117.312 330.688-352 544z"/><path fill="currentColor" d="M512 448a64 64 0 1 0 0-128a64 64 0 0 0 0 128zm0 64a128 128 0 1 1 0-256a128 128 0 0 1 0 256zm345.6 192L960 960H672v-64H352v64H64l102.4-256h691.2zm-68.928 0H235.328l-76.8 192h706.944l-76.8-192z"/>`),
		g.Group(children),
	)
}