package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzureSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.97.67a.5.5 0 0 0-.824-.524l-4 4a.5.5 0 0 0-.106.157l-3 7A.5.5 0 0 0 .5 12h3a.5.5 0 0 0 .47-.33l4-11Zm1.985 1.623a.5.5 0 0 0-.92.021l-2 5A.5.5 0 0 0 7.1 7.8l2.584 3.445l-5.342 1.78A.5.5 0 0 0 4.5 14h10a.5.5 0 0 0 .455-.707l-5-11Z"/>`),
		g.Group(children),
	)
}