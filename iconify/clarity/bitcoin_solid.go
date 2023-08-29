package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M21.18 18.47H14.5v6h6.68a2.7 2.7 0 0 0 2.63-2.77v-.48a2.71 2.71 0 0 0-2.63-2.75Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M23 13.75a2.24 2.24 0 0 0-2.23-2.25H14.5V16h6.3a2.22 2.22 0 0 0 2.2-2.25Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm8.31 19.73A5.22 5.22 0 0 1 21.18 27H21v1.9a1 1 0 0 1-2 0V27h-2v1.9a1 1 0 0 1-2 0V27h-1.75A1.25 1.25 0 0 1 12 25.75V10.23A1.25 1.25 0 0 1 13.25 9H15V7.07a1 1 0 0 1 2 0V9h2V7.07a1 1 0 0 1 2 0V9a4.72 4.72 0 0 1 3.2 8a5.31 5.31 0 0 1 2.11 4.24Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}