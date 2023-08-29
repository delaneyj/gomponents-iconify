package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirbudsCaseBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.002 9.75C3 10.142 3 10.558 3 11v2c0 3.75 0 5.625.955 6.939a5 5 0 0 0 1.106 1.106C6.375 22 8.251 22 12 22c3.75 0 5.625 0 6.939-.955a5 5 0 0 0 1.106-1.106C21 18.625 21 16.749 21 13v-2c0-.442 0-.858-.002-1.25h-3.352a2.751 2.751 0 0 1-2.646 2H9a2.751 2.751 0 0 1-2.646-2H3.002Z"/><path d="M3.021 8.25h3.333A2.751 2.751 0 0 1 9 6.25h6c1.259 0 2.32.846 2.647 2h3.332c-.055-2.01-.248-3.245-.934-4.189a5 5 0 0 0-1.106-1.106C17.625 2 15.75 2 12 2s-5.624 0-6.939.955A5 5 0 0 0 3.955 4.06c-.686.944-.88 2.178-.934 4.189Z" opacity=".5"/><path d="M7.75 9c0-.69.56-1.25 1.25-1.25h6a1.25 1.25 0 1 1 0 2.5H9c-.69 0-1.25-.56-1.25-1.25Z" opacity=".7"/></g>`),
		g.Group(children),
	)
}