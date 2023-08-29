package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePictureFlowerPhotosPhotoPictureCameraPhotographyPicturesFlower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M10 6.5a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3V4.31a.5.5 0 0 1 .72-.45L7 5l2.28-1.14a.5.5 0 0 1 .72.45Zm-3 3v4m0 0L9.5 11M7 13.5L4.5 11"/></g>`),
		g.Group(children),
	)
}