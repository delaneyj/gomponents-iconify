package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm0-148a58.07 58.07 0 0 0-58 58a6 6 0 0 1-12 0a70.08 70.08 0 0 1 70-70a6 6 0 0 1 0 12Zm70 58a70.08 70.08 0 0 1-70 70a6 6 0 0 1 0-12a58.07 58.07 0 0 0 58-58a6 6 0 0 1 12 0Zm-40 0a30 30 0 1 0-30 30a30 30 0 0 0 30-30Zm-48 0a18 18 0 1 1 18 18a18 18 0 0 1-18-18Z"/>`),
		g.Group(children),
	)
}