package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tokopedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.043 12.942c-3.43-2.897-16.85-2.247-16.85-2.247l-.473 32.65s17.855.134 23.353 0s9.341-4.508 9.4-7.878s0-24.18 0-24.18c-6.858-.829-11.942-.178-15.43 1.655Z"/><circle cx="19.531" cy="24.172" r="6.976" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.043 29.33a6.271 6.271 0 1 0-2.3-1.786m-19.55-16.849l-4.494 3.252L5.5 39.369l4.22 3.977m23.975-32.251a7.796 7.796 0 0 0-15.318-.299"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.396 19.662a2.36 2.36 0 0 1-3.878 2.59a4.194 4.194 0 1 0 3.878-2.59Zm-13.872.345a2.424 2.424 0 0 1-4.251 2.211a4.31 4.31 0 1 0 4.25-2.21Zm3.837 11.411c0-2.817 2.031-3.962 4.721-3.962c2.395 0 3.755 3.252 3.755 3.252a18.179 18.179 0 0 1-7.45 1.448a9.904 9.904 0 0 0 5.321 2.543s-.827.62-3.665.62c-2.306.001-2.682-2.453-2.682-3.902Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.317 31.569a10.394 10.394 0 0 1-.258 3.008"/>`),
		g.Group(children),
	)
}