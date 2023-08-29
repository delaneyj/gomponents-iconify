package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JwLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.824 12.433v10.258a3.416 3.416 0 0 1-3.412 3.419h0A3.416 3.416 0 0 1 12 22.69v-1.135m24-9.122L32.581 26.11l-3.419-13.677l-3.42 13.677l-3.419-13.677M8.5 29.556v6.011h3.006m1.543-6.011v6.011m7.081 0v-6.011h1.97a2.019 2.019 0 0 1 0 4.037h-1.97m1.969 0l1.968 1.973m6.68.001v-6.011h1.968a2.019 2.019 0 0 1 0 4.037h-1.968m1.968 0l1.968 1.973m4.817-6.01l-1.991 3.005l-1.992-3.005m1.992 6.011v-3.006m-20.354 0a1.503 1.503 0 1 1 0 3.006h-2.48v-6.011h2.48a1.503 1.503 0 1 1 0 3.005Zm0 0h-2.479m10.72 2.988l1.953-5.994m1.955 6.012l-1.955-6.012m1.301 4.001h-2.604"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}