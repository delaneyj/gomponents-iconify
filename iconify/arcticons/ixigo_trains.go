package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IxigoTrains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.12 26.61s-8.54 2.634-18.904-16.63h-2.09V7.16H9.398v3.107H8.227s-1.786-.74-3.508 8.208v10.317s.143 4.6 2.848 4.933s34.266 0 34.266 0s2.735-2.863.288-7.117Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.785 21.187a1.42 1.42 0 0 1 1.42 1.42v3.154a1.42 1.42 0 0 1-2.84 0v-3.154a1.42 1.42 0 0 1 1.42-1.42Zm4.385 3.393a1.222 1.222 0 0 1 1.222 1.221v2.284a1.222 1.222 0 0 1-2.443 0v-2.283a1.221 1.221 0 0 1 1.221-1.222Zm3.999 2.256a1.042 1.042 0 0 1 1.042 1.042v1.534a1.042 1.042 0 1 1-2.084 0v-1.534a1.042 1.042 0 0 1 1.042-1.042Zm3.803 1.36a.838.838 0 0 1 .838.838v1.274a.838.838 0 0 1-1.676 0v-1.274a.838.838 0 0 1 .838-.838Z"/><circle cx="9.015" cy="25.963" r="2.353" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.081" cy="25.963" r="2.353" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.148 13.265a3.176 3.176 0 1 1 0 6.352m-7.33-6.352a3.176 3.176 0 1 0 0 6.352h7.33m0-6.352h-7.33"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M42.411 35.895a.985.985 0 0 1 1.089 1.09c0 1.197-.46 1.06-.46 1.06h-1.18c-12.598 0-18.042 2.793-18.042 2.793H4.5v-2.227h2.021v-2.42h5.438v2.321h7.976v-2.617Z"/>`),
		g.Group(children),
	)
}