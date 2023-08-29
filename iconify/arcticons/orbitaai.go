package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orbitaai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 33.462L14.539 24L24 14.539L33.462 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.23L8.23 24L24 39.77L39.77 24Zm1.577-1.576L28.73 3.5m12.616 18.923L44.5 19.27m-6.308 0l3.154-3.155M28.73 9.808l3.155-3.154m9.461 18.923L44.5 28.73m-6.308 0l3.154 3.155m-15.769 9.461L28.73 44.5m0-6.308l3.155 3.154M22.423 6.654L19.27 3.5m0 6.308l-3.155-3.154m6.308 34.692L19.27 44.5m0-6.308l-3.155 3.154M6.654 25.577L3.5 28.73m6.308 0l-3.154 3.155m0-9.462L3.5 19.27m6.308 0l-3.154-3.155m9.461-3.154l-3.153-3.153m-.001 6.307l-3.153-3.153m22.077-.001l3.153-3.153m0 6.307l3.154-3.153h0m-6.307 22.076l3.153 3.154h0m0-6.307l3.154 3.153m-22.077 0l-3.153 3.154m-.001-6.307l-3.153 3.153"/>`),
		g.Group(children),
	)
}