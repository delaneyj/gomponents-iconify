package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heartbeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 41.824l16.23-19.05A10.13 10.13 0 1 0 24 10.694"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 21.113h5l3-6l4 12l3-6h5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 41.824L7.77 22.774A10.13 10.13 0 1 1 24 10.693"/>`),
		g.Group(children),
	)
}