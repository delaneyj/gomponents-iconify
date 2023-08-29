package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 18A12 12 0 1 1 18 6a12 12 0 0 1 12 12Zm-4.77-2.16a1.4 1.4 0 0 0-2-2l-6.77 6.77L13 17.16a1.4 1.4 0 0 0-2 2l5.45 5.45Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}