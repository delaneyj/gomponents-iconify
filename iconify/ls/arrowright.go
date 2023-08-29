package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 251v215c0 26 19 45 46 45h250v172c0 16 7 26 22 33c4 1 9 2 13 2c10 0 18-3 25-10l323-325c14-12 14-36 0-50L356 11c-20-22-60-8-60 25v171H46c-27 0-46 18-46 44z"/>`),
		g.Group(children),
	)
}