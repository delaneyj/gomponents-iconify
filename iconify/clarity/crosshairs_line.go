package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 29a11 11 0 1 1 11-11a11 11 0 0 1-11 11Zm0-20a9 9 0 1 0 9 9a9 9 0 0 0-9-9Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 23a5 5 0 1 1 5-5a5 5 0 0 1-5 5Zm0-8a3 3 0 1 0 3 3a3 3 0 0 0-3-3Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 9a1 1 0 0 1-1-1V2.8a1 1 0 0 1 2 0V8a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M18 34a1 1 0 0 1-1-1v-5a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M8 19H3.17a1 1 0 0 1 0-2H8a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M33.1 19H28a1 1 0 0 1 0-2h5.1a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}