package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmwAppOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28 22h2v8h-8v-2h-2v4h12V20h-4v2z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M14 30H6v-8h2v-2H4v12h12v-4h-2v2z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M8 14H6V6h8v2h2V4H4v12h4v-2z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M11 11h6v6h-6z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M11 19h6v6h-6z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><path fill="currentColor" d="M19 19h6v6h-6z" class="clr-i-outline--badged clr-i-outline-path-6--badged"/><path fill="currentColor" d="M22 6h.5a7.49 7.49 0 0 1 .28-2H20v4h2Z" class="clr-i-outline--badged clr-i-outline-path-7--badged"/><path fill="currentColor" d="M30 13.5v.5h-2v2h4v-2.78a7.49 7.49 0 0 1-2 .28Z" class="clr-i-outline--badged clr-i-outline-path-8--badged"/><path fill="currentColor" d="M25 11.58a7.53 7.53 0 0 1-.58-.58H19v6h6Z" class="clr-i-outline--badged clr-i-outline-path-9--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-10--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}