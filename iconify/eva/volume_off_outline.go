package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVolumeOffOutline0"><g id="evaVolumeOffOutline1"><g id="evaVolumeOffOutline2" fill="currentColor"><path d="M4.71 3.29a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm12.2 10.79l1.44 1.44a6 6 0 0 0-.07-7.15a1 1 0 1 0-1.56 1.26a4 4 0 0 1 .19 4.45Z"/><path d="M21 12a6.51 6.51 0 0 1-1.78 4.39l1.42 1.42A8.53 8.53 0 0 0 23 12a8.75 8.75 0 0 0-3.36-6.77a1 1 0 1 0-1.28 1.54A6.8 6.8 0 0 1 21 12Zm-7.5 6.1l-5.1-3.5a1 1 0 0 0-.57-.17H3.5V9.57h3.24l-2-2H2.5a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4a1.06 1.06 0 0 0 .57.17a1 1 0 0 0 1-1v-1.67l-2-2Zm0-12.2v4.77l2 2V4a1 1 0 0 0-1.57-.83L9.23 6.4l1.44 1.44Z"/></g></g></g>`),
		g.Group(children),
	)
}