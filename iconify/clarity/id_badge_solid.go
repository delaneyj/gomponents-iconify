package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdBadgeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="17.77" r="4.23" fill="currentColor" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M21 4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6h6Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M10.26 27a1.13 1.13 0 0 0-.26.73V30h16v-2.3a1.12 1.12 0 0 0-.26-.73A9.9 9.9 0 0 0 18 23.69A9.9 9.9 0 0 0 10.26 27Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="currentColor" d="M28 6h-5v2h5v24H8V8h5V6H8a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2Z" class="clr-i-solid clr-i-solid-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}