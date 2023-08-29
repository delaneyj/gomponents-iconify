package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DateOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13.22V30H4V8h3V6H3.75A1.78 1.78 0 0 0 2 7.81v22.38A1.78 1.78 0 0 0 3.75 32h28.5A1.78 1.78 0 0 0 34 30.19V12.34a7.45 7.45 0 0 1-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M8 14h2v2H8z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M14 14h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M20 14h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M26 14h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><path fill="currentColor" d="M8 19h2v2H8z" class="clr-i-outline--badged clr-i-outline-path-6--badged"/><path fill="currentColor" d="M14 19h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-7--badged"/><path fill="currentColor" d="M20 19h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-8--badged"/><path fill="currentColor" d="M26 19h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-9--badged"/><path fill="currentColor" d="M8 24h2v2H8z" class="clr-i-outline--badged clr-i-outline-path-10--badged"/><path fill="currentColor" d="M14 24h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-11--badged"/><path fill="currentColor" d="M20 24h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-12--badged"/><path fill="currentColor" d="M26 24h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-13--badged"/><path fill="currentColor" d="M10 10a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v6a1 1 0 0 0 1 1Z" class="clr-i-outline--badged clr-i-outline-path-14--badged"/><path fill="currentColor" d="M22.5 6H13v2h9.78a7.49 7.49 0 0 1-.28-2Z" class="clr-i-outline--badged clr-i-outline-path-15--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-16--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}