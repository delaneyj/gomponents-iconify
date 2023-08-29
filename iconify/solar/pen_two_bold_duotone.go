package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.25 22a.75.75 0 0 1 .75-.75h16a.75.75 0 0 1 0 1.5H4a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd" opacity=".5"/><path d="M19.08 7.372a3.147 3.147 0 0 0-4.45-4.45l-.71.71l.031.089c.26.75.751 1.733 1.675 2.656a7.004 7.004 0 0 0 2.745 1.705l.71-.71Z" opacity=".5"/><path d="m13.951 3.6l-.03.03l.03.09c.26.75.75 1.732 1.674 2.656A7.005 7.005 0 0 0 18.37 8.08l-6.85 6.85c-.462.462-.693.693-.948.891c-.3.234-.625.435-.969.6c-.291.138-.601.241-1.22.448l-3.268 1.09a.849.849 0 0 1-1.073-1.074l1.089-3.268c.206-.62.31-.93.448-1.22c.164-.344.365-.67.6-.97c.198-.254.429-.485.89-.947l6.882-6.88Z"/></g>`),
		g.Group(children),
	)
}