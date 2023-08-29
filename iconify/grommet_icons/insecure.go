package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insecure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 6.919V6a4.724 4.724 0 0 1 5-5a4.724 4.724 0 0 1 5 5v5.052M12 23a7 7 0 1 0-7-7a7 7 0 0 0 7 7zm2.985-7h-5.97"/>`),
		g.Group(children),
	)
}