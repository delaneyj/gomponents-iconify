package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toys(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M235 240q0-48 34.5-82.5T352 123t82.5 34.5T469 240H235zm0 0q0 48-35 82.5T117 357t-82.5-34.5T0 240h235zm0 0q-48 0-83-34.5T117 123t35-83t83-35v235zm0 0q48 0 82.5 34.5T352 357t-34.5 83t-82.5 35V240z"/>`),
		g.Group(children),
	)
}