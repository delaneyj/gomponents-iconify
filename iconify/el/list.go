package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 983.901h216.099V1200H0V983.901zm0 0h216.099V1200H0V983.901zm0-327.966h216.099v216.1H0v-216.1zm0-327.968h216.099v216.099H0V327.967zM0 0h216.099v216.098H0V0zm317.596 983.901H1200V1200H317.596V983.901zm0 0H1200V1200H317.596V983.901zm0-327.966H1200v216.1H317.596v-216.1zm0-327.968H1200v216.099H317.596V327.967zm0-327.967H1200v216.098H317.596V0z"/>`),
		g.Group(children),
	)
}