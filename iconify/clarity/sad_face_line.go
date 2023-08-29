package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadFaceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm0 30a14 14 0 1 1 14-14a14 14 0 0 1-14 14Z" class="clr-i-outline clr-i-outline-path-1"/><circle cx="25.16" cy="14.28" r="1.8" fill="currentColor" class="clr-i-outline clr-i-outline-path-2"/><circle cx="11.41" cy="14.28" r="1.8" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M18.16 20a9 9 0 0 0-7.33 3.78a1 1 0 1 0 1.63 1.16a7 7 0 0 1 11.31-.13a1 1 0 0 0 1.6-1.2A9 9 0 0 0 18.16 20Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}