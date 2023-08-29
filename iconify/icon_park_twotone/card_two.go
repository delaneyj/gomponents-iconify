package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCardTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M28 12V4L8 14v28l12-6"/><path fill="#555" d="M20 16L40 6v28L20 44V16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCardTwo0)"/>`),
		g.Group(children),
	)
}