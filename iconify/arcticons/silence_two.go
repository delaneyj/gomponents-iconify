package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SilenceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.706 30.274a61.99 61.99 0 0 1-6.298-1.11a3.109 3.109 0 0 0-2.999.58c-.53.54-2.059 2.05-3.648 3.609a28.68 28.68 0 0 1-13.096-13.105c1.55-1.59 3-3.11 3.589-3.639a3.109 3.109 0 0 0 .58-2.999a61.821 61.821 0 0 1-1.11-6.308A2 2 0 0 0 15.53 5.52a1.95 1.95 0 0 0-.105.014H6.908a1.5 1.5 0 0 0-1.36 1.37c-.55 7.687 3.739 15.914 4.608 17.494h0v.06l.12.23h0a35.428 35.428 0 0 0 12.996 12.995h0l.44.25h0c2 1.06 9.947 5.058 17.374 4.508a1.5 1.5 0 0 0 1.39-1.36v-8.507a2 2 0 0 0-1.665-2.286a2.587 2.587 0 0 0-.105-.013Z"/><circle cx="34.503" cy="13.506" r="7.997" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.214 13.506h10.578"/>`),
		g.Group(children),
	)
}