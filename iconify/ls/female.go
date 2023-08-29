package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M245.116 76c0-42-34-76-76-76c-41 0-75 34-75 76s34 76 75 76c42 0 76-34 76-76zm42 124l50 157c5 13-3 27-16 31c-13 5-27-3-32-16l-27-83c-4-13-16-24-27-24c-10 0-15 11-11 24l58 179c4 13-4 24-18 24h-25v203c0 14-12 25-25 25c-14 0-26-11-26-25V492h-38v203c0 14-11 25-25 25s-25-11-25-25V492h-25c-14 0-22-11-18-24l58-179c4-13-1-24-12-24c-10 0-22 11-27 24l-27 83c-4 13-18 21-31 16c-14-4-21-18-17-31l51-157c4-13 19-24 33-24h169c14 0 28 11 33 24z"/>`),
		g.Group(children),
	)
}