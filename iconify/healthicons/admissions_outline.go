package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdmissionsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22.835 10.065a1 1 0 0 0-1.316.95v1.567H15v22.817h6.52v1.814a1 1 0 0 0 1.315.948l8.865-2.955a1 1 0 0 0 .684-.948V13.969a1 1 0 0 0-.684-.949l-8.865-2.955Zm3.03 13.324c0 .8-.324 1.449-.724 1.449c-.4 0-.725-.649-.725-1.449s.325-1.449.725-1.449c.4 0 .724.649.724 1.45Zm-4.346-8.807H17v18.817h4.52V14.582Z"/><path d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm31 3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v30a1 1 0 0 0 1 1h30a1 1 0 0 0 1-1V9Z"/></g>`),
		g.Group(children),
	)
}