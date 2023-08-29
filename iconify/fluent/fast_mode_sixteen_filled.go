package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastModeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10.202 4.004l3.185 3.185a2.092 2.092 0 0 1-1.962 3.515v.12c0 .648-.526 1.174-1.174 1.174H8.063v-.499c0-.8-.64-1.499-1.563-1.499h-1a.5.5 0 1 0 0 1h1c.373 0 .563.253.563.5v.498h-2.89A1.174 1.174 0 0 1 3 10.824v-1.53c0-.036 0-.07.002-.105A1.783 1.783 0 1 1 5.43 7.008c.078-.005.157-.008.237-.008h3.143c.354 0 .692.057 1.001.183a1.84 1.84 0 0 0 .119-.309c.06-.182.125-.382.207-.465l-1.17-1.17a.874.874 0 0 1 1.236-1.235z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}