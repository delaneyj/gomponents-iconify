package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backgroundrestrictor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.311 41.09a17.417 17.417 0 0 0 8.991-29.392M20.675 6.924a17.397 17.397 0 0 0-8.977 29.379"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.4 13.6l-4.628 4.628l.013-12.727H42.5l-4.629 4.628L34.4 13.6zM13.599 34.401l4.615-4.615v12.715H5.5l4.628-4.629l3.471-3.471zM42.5 42.5l-37-37"/>`),
		g.Group(children),
	)
}