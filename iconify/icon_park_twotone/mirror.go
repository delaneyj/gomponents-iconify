package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMirror0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M34 44H14c-5 0-7-1-7-4v-6h11l1 4h11l1-4h10v6c0 3-2 4-7 4Z"/><path d="M12 34c-2.045-3.118-5-7-5-12.146C7 11.422 14.611 4 24 4s17 7.422 17 17.854C41 27 38.044 30.882 36 34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMirror0)"/>`),
		g.Group(children),
	)
}