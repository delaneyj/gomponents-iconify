package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletListLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="5.21" cy="9.17" r="2" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><circle cx="5.21" cy="17.17" r="2" fill="currentColor" class="clr-i-outline clr-i-outline-path-2"/><circle cx="5.21" cy="25.17" r="2" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M32.42 9a1 1 0 0 0-1-1H10v2h21.42a1 1 0 0 0 1-1Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M31.42 16H10v2h21.42a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M31.42 24H10v2h21.42a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}