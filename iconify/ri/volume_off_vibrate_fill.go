package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffVibrateFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.39 3.161l1.414 1.414L18.33 7.05l2.475 2.475L18.33 12l2.475 2.476l-2.475 2.475l2.475 2.475l-1.414 1.414l-3.889-3.89l2.475-2.474L15.5 12l2.475-2.475L15.5 7.05l3.89-3.889Zm-6.389 16.784a.5.5 0 0 1-.817.387L6.89 15.999L3.001 16a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1l2.584-.002l-3.776-3.776l1.414-1.414L13 12.586v7.359Zm-.113-16.206a.5.5 0 0 1 .113.316v5.702L9.283 6.04l2.901-2.372a.5.5 0 0 1 .704.07Z"/>`),
		g.Group(children),
	)
}