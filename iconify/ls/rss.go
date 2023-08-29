package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 126V27c39-8 80-13 122-13c300 0 544 244 544 544c0 42-6 83-14 122h-98c11-39 16-80 16-122c0-248-200-449-448-449c-42 0-83 6-122 17zm0 228V248c38-15 79-23 122-23c183 0 333 150 333 333c0 43-9 84-24 122H326c22-36 33-78 33-122c0-131-106-238-237-238c-45 0-86 12-122 34zm122 82c68 0 122 54 122 122s-54 122-122 122S0 626 0 558s54-122 122-122z"/>`),
		g.Group(children),
	)
}