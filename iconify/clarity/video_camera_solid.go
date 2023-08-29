package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32.3 9.35L26 12.9V8a2 2 0 0 0-2-2H6a4 4 0 0 0-4 4v16a4 4 0 0 0 4 4h18a2 2 0 0 0 2-2v-4.92l6.3 3.55a1.1 1.1 0 0 0 1.7-.86V10.2a1.1 1.1 0 0 0-1.7-.85Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}