package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnarchiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 32H7V22H5v10a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V22h-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M14 24a1 1 0 0 0 1 1h6a1 1 0 0 0 0-2h-6a1 1 0 0 0-1 1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M15 18H6v-4h9v-2H5.5A1.5 1.5 0 0 0 4 13.5V20h11.78a3 3 0 0 1-.78-2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M30.5 12H21v2h9v4h-9a3 3 0 0 1-.78 2H32v-6.5a1.5 1.5 0 0 0-1.5-1.5Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="m13 9.55l4-3.95V18a1 1 0 1 0 2 0V5.6l4 3.95a1 1 0 1 0 1.41-1.42L18 1.78l-6.39 6.35A1 1 0 0 0 13 9.55Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}