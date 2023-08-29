package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.426 17.839A8.94 8.94 0 0 1 13 18.945V23h-2v-4.055A9.004 9.004 0 0 1 3.056 11h2.016a7.002 7.002 0 0 0 9.87 5.354l-1.55-1.55A5 5 0 0 1 7.001 10V8.414L1.394 2.808l1.414-1.415l19.799 19.8l-1.415 1.414l-4.767-4.768Zm2.95-2.679l-1.443-1.442c.509-.81.856-1.73.997-2.718h2.016a8.95 8.95 0 0 1-1.57 4.16Zm-2.91-2.909l-8.78-8.78A5 5 0 0 1 17 6v4a4.959 4.959 0 0 1-.533 2.251Z"/>`),
		g.Group(children),
	)
}