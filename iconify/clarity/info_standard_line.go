package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoStandardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="17.97" cy="10.45" r="1.4" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M21 25h-2V14.1h-3a1 1 0 0 0 0 2h1V25h-2a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 34a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm0-30a14 14 0 1 0 14 14A14 14 0 0 0 18 4Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}