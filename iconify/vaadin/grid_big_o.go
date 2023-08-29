package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridBigO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 7h7V0H0v7zm1-6h5v5H1V1zm8-1v7h7V0H9zm6 6h-5V1h5v5zM0 16h7V9H0v7zm1-6h5v5H1v-5zm8 6h7V9H9v7zm1-6h5v5h-5v-5z"/>`),
		g.Group(children),
	)
}