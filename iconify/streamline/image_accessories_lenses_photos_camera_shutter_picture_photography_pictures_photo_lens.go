package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAccessoriesLensesPhotosCameraShutterPicturePhotographyPicturesPhotoLens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r="2.5"/><path d="M4.5 7V1M7 4.5h6M9.5 7v6M7 9.5H1"/></g>`),
		g.Group(children),
	)
}