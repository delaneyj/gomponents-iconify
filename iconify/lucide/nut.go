package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592A7.003 7.003 0 0 0 19 14v-4"/><path d="M12 4C8 4 4.5 6 4 8c-.243.97-.919 1.952-2 3c1.31-.082 1.972-.29 3-1c.54.92.982 1.356 2 2c1.452-.647 1.954-1.098 2.5-2c.595.995 1.151 1.427 2.5 2c1.31-.621 1.862-1.058 2.5-2c.629.977 1.162 1.423 2.5 2c1.209-.548 1.68-.967 2-2c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4Z"/></g>`),
		g.Group(children),
	)
}