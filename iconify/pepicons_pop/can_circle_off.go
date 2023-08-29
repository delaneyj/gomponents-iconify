package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CanCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.5 5h-5c-1.667.667-2.667 1.667-3 3v10a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3V8c-.333-1.333-1.333-2.333-3-3Zm-6 13V8.294c.19-.502.624-.926 1.411-1.294h4.178c.787.368 1.221.792 1.411 1.294V18a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 12.5a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm2.805 1.405c-.248.061-.653.185-1.221.373a3.081 3.081 0 0 1-.061-.085c1.258-.378 2.272-.955 3.032-1.734c.682-.697 1.145-1.449 1.382-2.25c.149.126.285.266.407.419c.124.835.03 1.457-.259 1.884c-.355.526-.95.855-1.505 1c-.071.019-.2.048-.402.092l-.398.085l-.213.045l-.113.024c-.274.06-.483.107-.65.148Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}