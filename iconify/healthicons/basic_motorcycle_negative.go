package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasicMotorcycleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBasicMotorcycleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM34 31.5a3.5 3.5 0 1 0 3.067-3.474l1.349 3.072a1 1 0 1 1-1.832.804l-1.348-3.071A3.492 3.492 0 0 0 34 31.5ZM27.847 12l4.665 10.627a.997.997 0 0 0-.288.242l-6.2 7.631h-1.936L22.7 28H25a3 3 0 0 0 3-3H17.5a1.5 1.5 0 0 0 0 3h2.912l1.389 2.5h-4.877A6.502 6.502 0 0 0 4 31.5a6.5 6.5 0 0 0 12.924 1H26.5a1 1 0 0 0 .776-.37l6.108-7.517l.618 1.407a6.5 6.5 0 1 0 1.831-.805l-6.417-14.617A1 1 0 0 0 28.5 10H23v2h4.847ZM10.5 28a3.502 3.502 0 0 1 3.355 2.5H10.5a1 1 0 1 0 0 2h3.355A3.502 3.502 0 0 1 7 31.5a3.5 3.5 0 0 1 3.5-3.5Zm17.442-11h-4.257c-3.142 0-5.447 2.952-4.685 6h10.382a1 1 0 0 0 .894-1.447l-2.25-4.501a.094.094 0 0 0-.084-.052ZM7 18a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h9.42a1 1 0 0 0 .78-1.625l-1.199-1.498A5 5 0 0 0 12.097 18H7Zm27-1.5l6.067-2.5l.933 6l-6.067-1L34 16.5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBasicMotorcycleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}