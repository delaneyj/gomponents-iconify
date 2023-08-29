package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organization(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M20 3v20H4V3h16ZM8.042 9h2V7h-2v2ZM14 9h2V7h-2v2Zm-5.958 6h2v-2h-2v2Zm2 8h4v-4h-4v4ZM14 15h2v-2h-2v2ZM2 3h20V1H2v2Z"/>`),
		g.Group(children),
	)
}