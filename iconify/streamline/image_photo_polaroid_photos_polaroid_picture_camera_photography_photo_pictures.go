package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotoPolaroidPhotosPolaroidPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="9" x=".5" y=".5" rx="1" transform="rotate(180 5 5)"/><path d="m12 4.7l.78.25a1 1 0 0 1 .64 1.27l-2.18 6.59a1 1 0 0 1-1.26.64L5.56 12M.5 7h9"/></g>`),
		g.Group(children),
	)
}