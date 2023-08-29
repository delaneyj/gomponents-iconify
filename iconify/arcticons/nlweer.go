package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nlweer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.244 23.179c-1.53-12.35 14.209-13.553 15.957-2.186c3.826-2.951 8.635.328 7.76 5.356c8.963.109 7.651 10.601 2.405 10.82H10.95c-7.76.219-9.29-13.553 2.295-13.99l-.328-.219"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.997 24.71C5.921 23.397 4.5 20.992 4.5 18.37c0-4.153 3.388-7.542 7.542-7.542c2.404 0 4.48 1.093 5.902 2.842"/>`),
		g.Group(children),
	)
}