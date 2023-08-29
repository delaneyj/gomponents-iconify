package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandscapeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.27 12.216L15 6l8 15H2L9 8l2.27 4.216Zm1.12 2.022L14.987 19h4.68l-4.77-8.942l-2.507 4.18ZM5.348 19h7.304L9 12.219L5.348 19ZM5.5 8a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}