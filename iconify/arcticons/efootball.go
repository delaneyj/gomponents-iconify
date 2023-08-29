package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Efootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.543 14.407c5.295-5.301 13.88-5.301 19.174 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".974" d="M5.903 14.407a21.385 21.385 0 0 1 9.595-9.634c10.596-5.32 23.523-.974 28.84 9.634m-28.795 0h-9.64m28.814 0h9.62"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.65 33.593c-5.295 5.301-13.88 5.301-19.175 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".974" d="M44.29 33.593a21.385 21.385 0 0 1-9.596 9.634c-10.596 5.32-23.523.974-28.839-9.634m28.795 0h9.64m-28.815 0h-9.62m32.55-13.546H4.03m34.375 7.919H4.03m0-7.919c-.434 2.502-.543 5.099 0 7.919m34.375-7.919c.434 2.502.543 5.099 0 7.919"/>`),
		g.Group(children),
	)
}