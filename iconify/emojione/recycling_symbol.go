package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecyclingSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#68963b" d="M16.2 58.1h14.2V41.7H10.3s-.7 5.7-.7 7.7c-.3 8.6 6.6 8.7 6.6 8.7"/><path fill="#83bf4f" d="m19.5 34.8l6.5 3.8l-8.6-15.2l-17.4.1l6.5 3.8l-5.2 9C.8 37.2.2 38.7 1 40l8.9 15.3c1.6 2.8 4.4 2.7 6.4 2.7c0 0-9.1-1.9-3.3-11.9l6.5-11.3"/><path fill="#68963b" d="m62.5 32.7l-7.1-12.3l-14.2 8.2l10 17.4s5.3-2.3 7-3.2c7.6-4.1 4.3-10.1 4.3-10.1"/><path fill="#83bf4f" d="M40.6 41.5V34l-8.9 15l8.9 15v-7.5H51c1.1 0 2.6-.3 3.4-1.6l8.9-15.3c1.6-2.8.2-5.1-.8-6.9c0 0 2.9 8.8-8.7 8.8H40.6"/><path fill="#68963b" d="m18 4.1l-7.1 12.3l14.2 8.2l10-17.4s-4.6-3.5-6.3-4.5C21.6-1.7 18 4.1 18 4.1z"/><path fill="#83bf4f" d="m36.6 18.7l-6.5 3.8l17.4.2l8.6-15.2l-6.5 3.8l-5.2-9C43.8 1.2 42.8 0 41.3 0H23.6C20.4 0 19 2.4 18 4.1c0 0 6.2-6.9 12 3.1l6.6 11.5"/>`),
		g.Group(children),
	)
}