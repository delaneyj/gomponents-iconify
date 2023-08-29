package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherSunPhotosLightCameraModeBrightnessSunPhotoFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="m13.5 7l-2.34 1.72l.44 2.88l-2.88-.44L7 13.5l-1.72-2.34l-2.88.44l.44-2.88L.5 7l2.34-1.72L2.4 2.4l2.88.44L7 .5l1.72 2.34l2.88-.44l-.44 2.88L13.5 7z"/></g>`),
		g.Group(children),
	)
}