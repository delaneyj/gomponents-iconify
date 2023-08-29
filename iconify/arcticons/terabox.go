package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terabox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.388 11.012c5.289-.285 10.791 2.763 12.617 9.736a8.159 8.159 0 0 1 7.467 8.796c-.323 3.959-3.402 7.462-7.407 7.462H13.26c-11.125 0-13.06-18.263 0-19.52a12.025 12.025 0 0 1 10.127-6.474Zm1.155 25.994V25.767"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.286 30.509l-4.743-4.742l-4.742 4.742"/>`),
		g.Group(children),
	)
}