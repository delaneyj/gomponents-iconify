package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 16C3.595 16 .025 12.424.025 8.027c0-4.395 3.57-7.971 7.959-7.971c4.389 0 7.959 3.576 7.959 7.971c0 4.397-3.57 7.973-7.959 7.973zM7.977 1.904c-3.363 0-6.102 2.732-6.102 6.086c0 3.357 2.739 6.09 6.102 6.09c3.365 0 6.102-2.733 6.102-6.09c-.001-3.354-2.738-6.086-6.102-6.086z"/><path d="m5.047 8.051l2.984 3.904l2.916-3.905H8.944V4.786c0-.344-.324-.79-.955-.79c-.63 0-.976.483-.976.826v3.229H5.047z"/></g>`),
		g.Group(children),
	)
}