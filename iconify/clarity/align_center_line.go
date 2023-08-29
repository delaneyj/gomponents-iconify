package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 20H19v-4h6a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-6V2a1 1 0 0 0-2 0v4h-6a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h6v4H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12v4a1 1 0 0 0 2 0v-4h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1Zm-19-6V8h12v6Zm18 14H6v-6h24Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}