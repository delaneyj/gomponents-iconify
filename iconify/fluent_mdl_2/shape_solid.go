package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 128h1152v1152h-234l-517-896l-103 177q-55-85-131-152T896 297V128zm205 595l-468 810q-29 3-57 3q-119 0-224-45t-183-124t-123-183T0 960q0-119 45-224t124-183t183-123t224-46q85 0 164 24t148 68t123 106t90 141zM595 1920l702-1216l702 1216H595z"/>`),
		g.Group(children),
	)
}