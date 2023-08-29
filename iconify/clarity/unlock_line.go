package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M12 25.14V28h2v-2.77a2.42 2.42 0 1 0-2-.09Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M26 2a8.2 8.2 0 0 0-8 8.36V15H2v17a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V15h-4v-4.64A6.2 6.2 0 0 1 26 4a6.2 6.2 0 0 1 6 6.36v6.83a1 1 0 0 0 2 0v-6.83A8.2 8.2 0 0 0 26 2Zm-4 15v15H4V17Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}