package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailymotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.068 11.313h-.023a3.1 3.1 0 1 0 .183 6.195h-.024a2.969 2.969 0 0 0 2.926-2.968l-.001-.076v.004l.002-.103a3.054 3.054 0 0 0-3.054-3.054h-.01h.001z"/><path fill="currentColor" d="M0 0v24h24V0zm20.693 20.807h-3.576V19.4a4.837 4.837 0 0 1-3.727 1.47h.011l-.104.001a5.646 5.646 0 0 1-3.83-1.49l.004.004A6.357 6.357 0 0 1 7.27 14.57l.001-.127v-.03a6.34 6.34 0 0 1 2.007-4.635l.003-.003a5.815 5.815 0 0 1 4.147-1.73h.041h-.002a4.186 4.186 0 0 1 3.526 1.578l.007.009V4.157l3.693-.765z"/>`),
		g.Group(children),
	)
}