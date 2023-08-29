package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RackServerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 9h2v2H6z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M10 9h14v2H10z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M6 17h2v2H6z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M10 17h14v2H10z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 7h28v6H4Zm0 8h28v6H4Zm0 14v-6h28v6Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M6 25h2v2H6z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M10 25h14v2H10z" class="clr-i-outline clr-i-outline-path-7"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}