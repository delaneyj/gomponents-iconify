package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotoFocusFramePhotosFramePhotoCameraPhotographyPicturesFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx="1"/><path d="M11 6.5V9H8.5M3 7.5V5h2.5"/></g>`),
		g.Group(children),
	)
}