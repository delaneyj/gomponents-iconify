package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 6v12h16V6H4Zm11.5 10a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm-4 0a2.5 2.5 0 0 1 0-5c.543 0 1.07.18 1.5.512a2.476 2.476 0 0 0 0 3.975a2.46 2.46 0 0 1-1.5.513Z"/>`),
		g.Group(children),
	)
}