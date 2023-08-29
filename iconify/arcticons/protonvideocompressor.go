package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Protonvideocompressor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.468 16.592A14.149 14.149 0 0 1 24.182 5.35a17.335 17.335 0 0 1 2.28 0a14.089 14.089 0 0 1 3.077.594m5.07 2.837a14.11 14.11 0 0 1 4.809 9.525a17.3 17.3 0 0 1 0 2.28a14.1 14.1 0 0 1-4.378 9.132m-4.947 3.043a14.073 14.073 0 0 1-3.631.78a17.3 17.3 0 0 1-2.28 0q-.497-.04-.985-.113c-.623-.094-1.612-.384-2.24-.422a1.912 1.912 0 0 0-2.083 1.107a6.717 6.717 0 0 0-.213 2.252v5.996a1.14 1.14 0 0 1-1.141 1.14h-5.198a1.14 1.14 0 0 1-1.141-1.14V22.128"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.056 20.584L21.33 27.355a1.317 1.317 0 0 1-1.976-1.14V12.673a1.317 1.317 0 0 1 1.976-1.141l11.727 6.77a1.317 1.317 0 0 1 0 2.282Z"/><circle cx="32.291" cy="7.137" r="2.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.623" cy="31.558" r="2.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.18" cy="19.444" r="2.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}