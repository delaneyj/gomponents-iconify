package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Affectscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M512 105.026H170.667L0 406.974h341.333L252.29 249.436h177.002l30.195-52.513H222.61l-22.261-39.385h283.676L512 105.026zM83.89 357.48l86.777-153.495l86.777 153.495H83.889z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}