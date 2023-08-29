package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormTrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7.5 9h9v10h-9V9ZM5 9h14M9.364 6h5v3h-5V6Zm1.181 5v6m3-6v6"/>`),
		g.Group(children),
	)
}