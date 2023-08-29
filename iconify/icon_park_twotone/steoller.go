package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steoller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSteoller0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M40 24H12l2.5 10H36l4-10Z"/><path d="m12 24l-4-9H3.5"/><circle cx="20" cy="41" r="3" fill="#555"/><circle cx="31" cy="41" r="3" fill="#555"/><path fill="#555" d="m23 8l5 16h12s3.5-8-2.5-14S25.667 6.667 23 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSteoller0)"/>`),
		g.Group(children),
	)
}