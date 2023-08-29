package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDirham(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 19H5m3.599-2.521A1.5 1.5 0 1 0 7.5 19M7 4v9m8 0h1.888a1.5 1.5 0 0 0 1.296-2.256L16 7m-5 6.01V13"/>`),
		g.Group(children),
	)
}