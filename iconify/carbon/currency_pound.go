package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.453 25A8.777 8.777 0 0 0 14 20a10.6 10.6 0 0 0-.18-2H22v-2h-8.783a47.4 47.4 0 0 0-.255-.624A9.859 9.859 0 0 1 12 11a4.792 4.792 0 0 1 5-5a6.123 6.123 0 0 1 5.222 2.628l1.556-1.256A8.11 8.11 0 0 0 17 4a6.778 6.778 0 0 0-7 7a11.65 11.65 0 0 0 1.056 5H8v2h3.773A8.209 8.209 0 0 1 12 20c0 2.523-1.486 5-3 5v2h15v-2Z"/>`),
		g.Group(children),
	)
}