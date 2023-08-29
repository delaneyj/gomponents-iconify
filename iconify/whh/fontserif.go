package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontserif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 192q0-53-28-90.5T736 64H608q-40 0-68 37.5T512 192v672q0 40 37.5 68t90.5 28v64H256v-64q53 0 90.5-28t37.5-68V192q0-53-28-90.5T288 64H160q-40 0-68 37.5T64 192H0V0h896v192h-64z"/>`),
		g.Group(children),
	)
}