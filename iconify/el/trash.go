package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M451.376 0v91.743H104.587V260.55h990.826V91.743H748.624V0H451.376zM157.798 388.991V1200h884.404V388.991H157.798z"/>`),
		g.Group(children),
	)
}