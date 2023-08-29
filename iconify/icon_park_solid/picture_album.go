package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureAlbum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPictureAlbum0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="30" height="30" x="6" y="6" fill="#fff" stroke="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M42 12v27a3 3 0 0 1-3 3H12"/><path stroke="#000" stroke-linecap="round" d="m6 25l7.656-6.806a2 2 0 0 1 2.674.015L26 27m-4-4l4.785-3.988a2 2 0 0 1 2.48-.063L36 24"/><path stroke="#fff" stroke-linecap="round" d="M6 19v8m30-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPictureAlbum0)"/>`),
		g.Group(children),
	)
}