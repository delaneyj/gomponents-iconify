package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCameraTwoPhotosPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="3.75" rx="1"/><circle cx="7" cy="8.25" r="2"/><path d="M9.5 3.75v-1.5a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v1.5"/></g>`),
		g.Group(children),
	)
}