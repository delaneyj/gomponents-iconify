package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCameraTripodTripodPhotosPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 8.5v5m0-5l-4.5 5m4.5-5l4.5 5"/><rect width="11" height="8" x="1.5" y=".5" rx="1"/><circle cx="7" cy="4.5" r="1.5"/></g>`),
		g.Group(children),
	)
}