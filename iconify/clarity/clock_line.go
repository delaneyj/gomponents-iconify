package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm0 30a14 14 0 1 1 14-14a14 14 0 0 1-14 14Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18.92 18.4v-7.65a1 1 0 0 0-2 0v8.72l5.9 4a1 1 0 1 0 1.11-1.66Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M8 17.94a9.94 9.94 0 0 1 15.41-8.35l.85-1.36a11.55 11.55 0 1 0-8.53 21L16 27.7a10 10 0 0 1-8-9.76Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}