package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmStripLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 4H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Zm0 26H6V6h24Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M14.6 23.07a1.29 1.29 0 0 0 1.24.09l8.73-4a1.3 1.3 0 0 0 0-2.37l-8.73-4A1.3 1.3 0 0 0 14 14v8a1.29 1.29 0 0 0 .6 1.07Zm1-8.6L23.31 18l-7.71 3.51Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M8 7h2v3H8z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M14 7h2v3h-2z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M20 7h2v3h-2z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M26 7h2v3h-2z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M8 26h2v3H8z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M14 26h2v3h-2z" class="clr-i-outline clr-i-outline-path-8"/><path fill="currentColor" d="M20 26h2v3h-2z" class="clr-i-outline clr-i-outline-path-9"/><path fill="currentColor" d="M26 26h2v3h-2z" class="clr-i-outline clr-i-outline-path-10"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}