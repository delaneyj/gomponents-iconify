package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileSearch0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M40 20.842V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h16"/><path d="M14 17h6m8 0h6M14 28h6m-6 6h6m-3-14v-6m20.728 23.728l4.243 4.243"/><circle cx="33.485" cy="33.485" r="6" fill="#555" stroke-linejoin="round" transform="rotate(45 33.485 33.485)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileSearch0)"/>`),
		g.Group(children),
	)
}