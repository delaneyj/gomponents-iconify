package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.685 5.284l7.019 8.136c.24.277.339.595.296.954c-.043.359-.28.568-.71.626H2.744c-.351-.025-.582-.199-.693-.522c-.11-.324-.054-.651.17-.981l7.097-8.213a.962.962 0 0 1 1.367 0Z"/>`),
		g.Group(children),
	)
}