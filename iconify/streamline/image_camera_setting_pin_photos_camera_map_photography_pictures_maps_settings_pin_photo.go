package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCameraSettingPinPhotosCameraMapPhotographyPicturesMapsSettingsPinPhoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 3.5a1 1 0 0 0-1-1h-2L9 .5H5l-1.5 2h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h3l2.5 4l2.5-4h3a1 1 0 0 0 1-1Z"/><circle cx="7" cy="5.5" r="2"/></g>`),
		g.Group(children),
	)
}