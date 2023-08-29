package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstructionWorkerNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsConstructionWorkerNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16.096 18.76c-2.664-.934-4.017-2.461-2.393-4.582c.015.047.03.094.048.14a1 1 0 1 0 1.882-.675c-.128-.356-.177-.874-.088-1.503l.02-.02l.024-.14c.252-1.339 1.04-2.867 2.411-4.045v4.545a1 1 0 0 0 2 0V6.69A8.928 8.928 0 0 1 23 6v4.48a1 1 0 0 0 2 0V6.058c.64.073 1.24.195 1.797.359c.424.125.825.275 1.203.445v5.618a1 1 0 0 0 2 0V8.112c1.404 1.178 2.226 2.673 2.442 4.025l.017.013c.08.57.037 1.086-.107 1.5a1 1 0 1 0 1.887.661a3.56 3.56 0 0 0 .048-.146c1.64 2.128.286 3.66-2.383 4.595a8 8 0 1 1-15.809 0ZM18 20c0-.232.013-.461.039-.686c3.616.73 8.306.73 11.922 0A6 6 0 1 1 18 20Zm1 8.49C13.013 29.535 6 32.203 6 36v6h36v-6c0-3.797-7.013-6.465-13-7.51V33H19v-4.51ZM13 42v-9h2v9h-2Zm21-9v9h2v-9h-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsConstructionWorkerNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}