package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Readerforselfoss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.79 43.5h6.4A17.79 17.79 0 0 0 8.4 25.7v6.41A11.39 11.39 0 0 1 19.79 43.5Zm-6.41 0h0a5 5 0 0 0-5-5h0v3a2 2 0 0 0 1.95 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.65 4.5h-27.3a2 2 0 0 0-1.95 2v12.79A24.21 24.21 0 0 1 32.61 43.5h5a2 2 0 0 0 2-2V6.45a2 2 0 0 0-1.96-1.95ZM13.57 15.02h14.14M13.57 9.17h20.86"/>`),
		g.Group(children),
	)
}