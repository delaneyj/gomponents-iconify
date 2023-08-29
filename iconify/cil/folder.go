package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 472H40a24.028 24.028 0 0 1-24-24V64a24.028 24.028 0 0 1 24-24h186.667a23.935 23.935 0 0 1 22.154 14.77L269.333 104H472a24.028 24.028 0 0 1 24 24v320a24.028 24.028 0 0 1-24 24ZM48 440h416V136H248l-26.667-64H48Z"/>`),
		g.Group(children),
	)
}