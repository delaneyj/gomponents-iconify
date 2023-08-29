package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChapterNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3.5V20l9-6v6l11-8l-11-8v6L1 3.5ZM22 2v20V2Z"/>`),
		g.Group(children),
	)
}