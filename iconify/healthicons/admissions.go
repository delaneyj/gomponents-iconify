package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Admissions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm15.52 5.582H17v18.817h4.52V14.582Zm0-2H15v22.817h6.52v1.814a1 1 0 0 0 1.315.948l8.865-2.955a1 1 0 0 0 .684-.948V13.969a1 1 0 0 0-.684-.949l-8.865-2.955a1 1 0 0 0-1.316.95v1.567Zm4.345 10.807c0 .8-.324 1.449-.724 1.449c-.4 0-.725-.649-.725-1.449s.325-1.449.725-1.449c.4 0 .724.649.724 1.45Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}