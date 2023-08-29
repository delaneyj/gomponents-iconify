package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flappycrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.198 4.5h5.604a12 12 0 0 1 12 12v27h0H9.198h0v-27a12 12 0 0 1 12-12Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 28.015c0 1.743 1.919 3.156 4.285 3.156s4.285-1.413 4.285-3.156s-1.918-3.155-4.285-3.155h-3.33a.954.954 0 0 0-.955.954Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.526 24.86c0-1.363 2.59-4.665 6.097-4.694a15.574 15.574 0 0 1 4.803.348m-10.9 10.18c0 1.762 2.415 4.139 4.48 4.139h5.454a6.588 6.588 0 0 0 4.339-2.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.238 30.197h5.22a1.942 1.942 0 0 0-2.025-2.454c-2.532 0-5.571.39-5.571 2.454s1.987 2.376 3.819 2.376a2.775 2.775 0 0 0 2.648-2.376m-5.084-5.064v-1.286m2.436 1.286v-1.286"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.566 27.828a13.092 13.092 0 0 1-3.1-.085a3.26 3.26 0 0 1-2.941-3.565c0-2.26 1.052-3.772 2.57-3.772a2.103 2.103 0 0 1 2.319 1.318c1.168-.214 3.33-.935 3.876 1.383s-1.89 4.636-1.89 4.636m-9.33-7.583a3.945 3.945 0 0 1 .137-1.358m-2.591.098a5.783 5.783 0 0 1 1.088 1.349m-3.932-.415a14.589 14.589 0 0 1 1.387 1.467"/>`),
		g.Group(children),
	)
}