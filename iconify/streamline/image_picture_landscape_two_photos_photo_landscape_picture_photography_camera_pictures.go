package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePictureLandscapeTwoPhotosPhotoLandscapePicturePhotographyCameraPictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M.5 11h13m-9.66 0l5.21-4.88a.5.5 0 0 1 .64 0l3.81 2.73"/><circle cx="4.5" cy="4.5" r="1.5"/></g>`),
		g.Group(children),
	)
}