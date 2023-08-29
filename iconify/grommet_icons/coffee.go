package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="2" d="M6.264 20.192c4.096 2.868 8.602-.081 11.47-4.177c2.868-4.095 4.097-9.338.002-12.206c-4.096-2.868-8.602.08-11.47 4.176c-2.868 4.096-4.098 9.339-.002 12.207Z"/><path fill="currentColor" d="M16.589 5.447c-1.393.246-5.326 2.375-5.408 5.98c-.083 3.604-2.787 6.594-3.77 7.126c1.803.042 5.326-2.375 5.408-5.98c.083-3.604 2.786-6.594 3.77-7.126Z"/></g>`),
		g.Group(children),
	)
}