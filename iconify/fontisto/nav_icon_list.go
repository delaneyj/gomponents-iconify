package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavIconList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.184 0H27.65v5.219H7.184zm0 9.39H27.65v5.219H7.184zm0 9.391H27.65V24H7.184zM0 0h5.219v5.219H0zm0 9.39h5.219v5.219H0zm0 9.391h5.219V24H0z"/>`),
		g.Group(children),
	)
}