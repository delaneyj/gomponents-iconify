package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm112-360v208c0 13.3-10.7 24-24 24s-24-10.7-24-24v-80H192v80c0 13.3-10.7 24-24 24s-24-10.7-24-24V152c0-13.3 10.7-24 24-24s24 10.7 24 24v80h128v-80c0-13.3 10.7-24 24-24s24 10.7 24 24z"/>`),
		g.Group(children),
	)
}