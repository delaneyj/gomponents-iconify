package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 13.5V30h-3v-8a2 2 0 0 0-2-2H11a2 2 0 0 0-2 2v8H6V6h4v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-1.68l-.43-.3H12V6h10.5a7.49 7.49 0 0 1 .28-2H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V13.22a7.49 7.49 0 0 1-2 .28ZM25 30H11v-8h14Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}