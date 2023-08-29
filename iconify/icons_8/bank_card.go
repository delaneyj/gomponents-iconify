package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3H5zm0 2h22c.565 0 1 .435 1 1v2H5v2h23v10c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1z"/>`),
		g.Group(children),
	)
}