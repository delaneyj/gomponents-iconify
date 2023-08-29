package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureAlbum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPictureAlbum0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="30" height="30" x="6" y="6" fill="#555" rx="2"/><path stroke-linecap="round" d="M42 12v27a3 3 0 0 1-3 3H12M6 25l7.656-6.806a2 2 0 0 1 2.674.015L26 27"/><path stroke-linecap="round" d="m22 23l4.785-3.988a2 2 0 0 1 2.48-.063L36 24M6 19v8m30-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPictureAlbum0)"/>`),
		g.Group(children),
	)
}