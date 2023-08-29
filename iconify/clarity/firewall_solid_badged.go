package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirewallSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M24 10.49V12h1.51A7.53 7.53 0 0 1 24 10.49Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M32 13.22V14H4v-2h8V8h2v4h8V8h.78a7.49 7.49 0 0 1-.28-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34a7.45 7.45 0 0 1-2 .88ZM14 28h-2v-4h2Zm10 0h-2v-4h2Zm8-6H4v-2h3v-4h2v4h8v-4h2v4h8v-4h2v4h3Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}