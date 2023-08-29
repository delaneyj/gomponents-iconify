package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13h-8a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V15a2 2 0 0 0-2-2Zm0 2v11h-8V15Zm-8 15v-2.4h8V30Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M20 22H4V6h24v5h2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M20 26H9a1 1 0 0 0 0 2h11Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}