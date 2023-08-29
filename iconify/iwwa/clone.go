package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M32.167 32.31H7.833v-8.734h24.335v8.734zm-23.334-1h22.335v-6.734H8.833v6.734z"/><path fill="currentColor" d="M29.822 35.393H10.178V31.31h19.644v4.083zm-18.644-1h17.644V32.31H11.178v2.083zm11.161-9.856H17.66v-8.308a6.057 6.057 0 0 1-3.707-5.576c0-3.334 2.712-6.046 6.046-6.046s6.046 2.712 6.046 6.046a6.056 6.056 0 0 1-3.707 5.576v8.308zm-3.678-1h2.679v-8.005l.334-.117a5.053 5.053 0 0 0 3.373-4.762c0-2.782-2.264-5.046-5.046-5.046s-5.046 2.264-5.046 5.046a5.053 5.053 0 0 0 3.373 4.762l.334.117v8.005z"/>`),
		g.Group(children),
	)
}