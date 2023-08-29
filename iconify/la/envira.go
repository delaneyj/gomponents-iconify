package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5c2.614 8.976 2.362 13.181 6.896 16.693c3.97 3.026 7.94 2.237 10.112 1.914L25.398 27h2l-3.98-3.98C23.393 20.613 29.01 5 5 5zm3.084 2.002c.135.011.45.121 1.047.396c3.999 1.85 5.408 4.592 6.931 7.4c1.096 2.023 3.019 5.103 4.374 6.095c1.356.983 2.836 1.709-.288.398c-3.134-1.311-5.417-5.032-6.931-7.85c-1.164-2.162-2.163-4.153-4.336-5.613c0 0-1.203-.86-.797-.826z"/>`),
		g.Group(children),
	)
}