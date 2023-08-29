package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm4.207-9L7.5 10.207L3.293 6h8.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}