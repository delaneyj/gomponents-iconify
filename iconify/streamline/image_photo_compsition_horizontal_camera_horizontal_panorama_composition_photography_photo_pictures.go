package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotoCompsitionHorizontalCameraHorizontalPanoramaCompositionPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.81 11.2a.52.52 0 0 0 .47 0a.49.49 0 0 0 .22-.41V3.26a.49.49 0 0 0-.22-.41a.52.52 0 0 0-.47-.05a14.67 14.67 0 0 1-11.62 0a.52.52 0 0 0-.47.05a.49.49 0 0 0-.22.41v7.48a.49.49 0 0 0 .22.41a.52.52 0 0 0 .47 0a14.67 14.67 0 0 1 11.62.05Z"/>`),
		g.Group(children),
	)
}