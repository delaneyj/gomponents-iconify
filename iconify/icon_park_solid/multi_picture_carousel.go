package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiPictureCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMultiPictureCarousel0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" fill="#fff" rx="2"/><path d="M20 42h8m6 0h2M4 42h2m36 0h2m-32 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMultiPictureCarousel0)"/>`),
		g.Group(children),
	)
}