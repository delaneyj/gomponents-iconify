package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpsideDownFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5c15.163 0 27.5 12.336 27.5 27.5S47.163 59.5 32 59.5z"/><circle cx="43.5" cy="37.408" r="5" fill="currentColor"/><circle cx="20.5" cy="37.408" r="5" fill="currentColor"/><path fill="currentColor" d="M45.771 22.143c-2.457-4.047-7.417-7.65-13.771-7.65s-11.313 3.604-13.771 7.65c-.658 1.083.217 2.255 1.187 1.578c8.11-5.656 17.106-5.623 25.168 0c.97.676 1.845-.495 1.187-1.578"/>`),
		g.Group(children),
	)
}