package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 6a12 12 0 1 0 12 12A12 12 0 0 0 18 6Zm0 22a10 10 0 1 1 10-10a10 10 0 0 1-10 10Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 20.07a1.3 1.3 0 0 1-1.3-1.3v-6a1.3 1.3 0 1 1 2.6 0v6a1.3 1.3 0 0 1-1.3 1.3Z" class="clr-i-outline clr-i-outline-path-2"/><circle cx="17.95" cy="23.02" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}