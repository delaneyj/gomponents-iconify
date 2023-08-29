package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotoCompositionOvalCameraFrameCompositionPhotographyPicturesLandscapePhotoOval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 11a10.34 10.34 0 0 0 1-4a10.34 10.34 0 0 0-1-4A13.27 13.27 0 0 0 7 2a13.27 13.27 0 0 0-5.5 1a10.34 10.34 0 0 0-1 4a10.34 10.34 0 0 0 1 4A13.27 13.27 0 0 0 7 12a13.27 13.27 0 0 0 5.5-1Z"/>`),
		g.Group(children),
	)
}