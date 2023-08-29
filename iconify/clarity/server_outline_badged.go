package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M15.2 27.8c0 1.5 1.2 2.8 2.8 2.8s2.8-1.2 2.8-2.8S19.5 25 18 25s-2.8 1.2-2.8 2.8zm4 0c0 .7-.6 1.2-1.2 1.2s-1.2-.6-1.2-1.2s.6-1.2 1.2-1.2s1.2.5 1.2 1.2z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M13 21h10v1.6H13z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M24 10.5c-.1-.2-.2-.3-.3-.5H12v1.6h12v-1.1z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M12 6v1.6h10.7c-.1-.5-.2-1.1-.2-1.6H12z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M26 12.3V32H10V4h12.8c.2-.7.5-1.4.9-2H9.5C8.7 2 8 2.7 8 3.5V34h20V13.2c-.7-.2-1.4-.5-2-.9z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-6--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}