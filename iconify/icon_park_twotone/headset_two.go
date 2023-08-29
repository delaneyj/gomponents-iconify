package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHeadsetTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M21 16h-6v9a3 3 0 1 0 6 0v-9Z"/><path d="M18 28v10h-6m9-22V4H10.5C8 4 5 6 5 10s3.5 6 6 6h10Zm6 6V10h10.5c2.5 0 5.5 2 5.5 6s-3.5 6-6 6H27Z"/><path fill="#555" d="M33 22h-6v9a3 3 0 1 0 6 0v-9Z"/><path d="M30 34v10h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHeadsetTwo0)"/>`),
		g.Group(children),
	)
}