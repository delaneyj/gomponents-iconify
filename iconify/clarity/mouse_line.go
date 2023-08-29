package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 34A10 10 0 0 1 8 24V12a10 10 0 0 1 20 0v12a10 10 0 0 1-10 10Zm0-30a8 8 0 0 0-8 8v12a8 8 0 0 0 16 0V12a8 8 0 0 0-8-8Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 15a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}