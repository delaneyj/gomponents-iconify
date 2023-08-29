package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aerospike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h24V0zm19.295 5.386v1.64l-3.576 1.586v7.363l3.576 1.602v1.565L5.672 12.98l-1.607-.688l1.607-.743zm-4.948 3.825L7.45 12.283l6.897 3.092Z"/>`),
		g.Group(children),
	)
}