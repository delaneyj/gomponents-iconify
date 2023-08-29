package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotoFourPhotosCameraPicturePhotographyPicturesFourPhoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="5" x=".5" y=".5" rx=".5"/><rect width="5" height="5" x="8.5" y=".5" rx=".5"/><rect width="5" height="5" x=".5" y="8.5" rx=".5"/><rect width="5" height="5" x="8.5" y="8.5" rx=".5"/></g>`),
		g.Group(children),
	)
}