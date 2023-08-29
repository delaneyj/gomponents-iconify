package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridChartSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M15 17H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M32 17H21a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1.59v.056A7.396 7.396 0 0 0 34 12.269V15a2 2 0 0 1-2 2Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M15 30H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2Z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><path fill="currentColor" d="M32 30H21a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2Z" class="clr-i-solid--badged clr-i-solid-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}