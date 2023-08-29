package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RackServerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M2 22h32v-8H2Zm8-5h14v2H10Zm-4 0h2v2H6Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M32 4H4a2 2 0 0 0-2 2v6h32V6a2 2 0 0 0-2-2ZM8 9H6V7h2Zm16 0H10V7h14Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M2 30a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-6H2Zm8-3h14v2H10Zm-4 0h2v2H6Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}