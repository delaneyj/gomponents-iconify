package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.93 19.2L17.33 4.6a2 2 0 0 0-1.41-.6H6a2 2 0 0 0-2 2v9.92a2 2 0 0 0 .59 1.41l14.6 14.6a2 2 0 0 0 2.83 0l9.9-9.9a2 2 0 0 0 .01-2.83ZM9.65 11.31a1.66 1.66 0 1 1 1.66-1.66a1.66 1.66 0 0 1-1.66 1.66Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6.33" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}