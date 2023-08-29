package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlag0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 44h8m-4 0V4"/><path fill="#fff" d="M40 6H12v16h28l-4-8l4-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlag0)"/>`),
		g.Group(children),
	)
}