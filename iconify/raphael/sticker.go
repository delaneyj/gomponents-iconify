package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sticker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 2c-1.042 0-1.916.376-2.57 1.087L2.895 13.137C2.302 13.785 2 14.58 2 15.5C2 22.943 8.056 29 15.5 29S29 22.943 29 15.5S22.943 2 15.5 2zm0 26C8.596 28 3 22.404 3 15.5c0-3.452 5.24-2.737 7.5-5c2.262-2.26 1.548-7.5 5-7.5C22.404 3 28 8.597 28 15.5S22.404 28 15.5 28z"/>`),
		g.Group(children),
	)
}