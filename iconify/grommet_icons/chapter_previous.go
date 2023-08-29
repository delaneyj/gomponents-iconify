package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChapterPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M23 3.5V20l-9-6v6L3 12l11-8v6l9-6.5ZM2 2v20V2Z"/>`),
		g.Group(children),
	)
}