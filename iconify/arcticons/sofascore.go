package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sofascore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.532 5.5h-37v18.347h9.047v-9.402h27.953V5.5Zm-37 37h37V24.153h-9.047v9.402H5.532V42.5Z"/>`),
		g.Group(children),
	)
}