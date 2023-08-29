package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotspotWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0h1a9 9 0 0 1 9 9v1h-2V9a7 7 0 0 0-7-7h-1V0ZM4 2h6v2H6v18h12V12h2v12H4V2Zm8 2h1a5 5 0 0 1 5 5v1h-2V9a3 3 0 0 0-3-3h-1V4Zm0 4h2.004v2.004H12V8Z"/>`),
		g.Group(children),
	)
}