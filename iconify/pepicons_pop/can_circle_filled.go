package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CanCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCanCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.5 5h-5c-1.667.667-2.667 1.667-3 3v10a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3V8c-.333-1.333-1.333-2.333-3-3Zm-6 13V8.294c.19-.502.624-.926 1.411-1.294h4.178c.787.368 1.221.792 1.411 1.294V18a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1Z"/><path d="M10 12.5a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm2.805 1.405c-.248.061-.653.185-1.221.373a3.081 3.081 0 0 1-.061-.085c1.258-.378 2.272-.955 3.032-1.734c.682-.697 1.145-1.449 1.382-2.25c.149.126.285.266.407.419c.124.835.03 1.457-.259 1.884c-.355.526-.95.855-1.505 1c-.071.019-.2.048-.402.092l-.398.085l-.213.045l-.113.024c-.274.06-.483.107-.65.148Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCanCircleFilled0)"/></g>`),
		g.Group(children),
	)
}