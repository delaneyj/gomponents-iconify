package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M19.63 12.12C17.51 9.28 14.88 2 14.88 2S12 2.83 12 5.25V15H2.23a29.46 29.46 0 0 0 1.44 9.74C5.61 30.27 7.8 32 7.8 32h6.86C16.9 32 21 30.06 24 28.31v-12.8a10.84 10.84 0 0 1-4.37-3.39Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M27 13a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h7V13Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}