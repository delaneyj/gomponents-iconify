package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarCodeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M5 7a1 1 0 0 0-1 1v22a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M9 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M17 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M21 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M25 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M29 7a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M33 7a1 1 0 0 0-1 1v22a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-8"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}