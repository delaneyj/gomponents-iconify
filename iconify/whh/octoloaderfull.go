package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octoloaderfull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 583V441l256-106v354zM643 280L749 25l250 250l-255 107zm-202-24L335 0h354L583 256H441zM281 381L25 275L276 25l106 255zm-25 202L0 689V335l256 106v142zm126 161L276 999L25 749l256-106zm201 24l106 256H335l106-256h142zm161-125l255 106l-250 250l-106-255z"/>`),
		g.Group(children),
	)
}