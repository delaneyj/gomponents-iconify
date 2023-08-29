package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M2 8h2v2H2z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7 10h24a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M2 14h2v2H2z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M31 14H7a1 1 0 0 0 0 2h24a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M2 20h2v2H2z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M31 20H7a1 1 0 0 0 0 2h24a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M2 26h2v2H2z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M31 26H7a1 1 0 0 0 0 2h24a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-8"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}