package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurveChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7 11.8a.8.8 0 1 1 0-1.6h6c2.404 0 3.368 1.707 4.653 6.278l.182.651l.184.651c1.313 4.595 2.53 6.42 4.981 6.42h6a.8.8 0 1 1 0 1.6h-6c-3.465 0-5.019-2.331-6.519-7.58a80.53 80.53 0 0 1-.186-.66l-.182-.649C15.043 13.105 14.305 11.8 13 11.8H7Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}