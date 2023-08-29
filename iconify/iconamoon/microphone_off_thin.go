package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOffThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}