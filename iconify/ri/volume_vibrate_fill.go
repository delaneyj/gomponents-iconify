package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeVibrateFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.39 3.161l1.413 1.414l-2.475 2.475l2.475 2.475L18.328 12l2.475 2.476l-2.475 2.475l2.475 2.475l-1.414 1.414l-3.889-3.89l2.475-2.474L15.5 12l2.475-2.475L15.5 7.05l3.89-3.889Zm-6.503.577a.5.5 0 0 1 .113.317v15.89a.5.5 0 0 1-.817.387L6.89 15.999L3 16a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h3.889l5.294-4.332a.5.5 0 0 1 .704.07Z"/>`),
		g.Group(children),
	)
}