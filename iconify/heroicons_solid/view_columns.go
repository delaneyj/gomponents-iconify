package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 18h2.75A2.25 2.25 0 0 0 19 15.75V4.25A2.25 2.25 0 0 0 16.75 2H14v16ZM12.5 2h-5v16h5V2ZM3.25 2H6v16H3.25A2.25 2.25 0 0 1 1 15.75V4.25A2.25 2.25 0 0 1 3.25 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}