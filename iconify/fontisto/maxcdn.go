package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M32.72 9.694L29.644 24h-6.262L26.72 8.4a1.763 1.763 0 0 0-.285-1.654l.003.004a1.89 1.89 0 0 0-1.563-.618h.007h-3.17L17.886 24h-6.262l3.825-17.869h-5.36L6.262 24H0L3.825 6.131L.96 0h23.975c1.268 0 2.471.28 3.55.781l-.052-.022a7.586 7.586 0 0 1 2.756 2.114l.011.014a7.354 7.354 0 0 1 1.511 3.11l.009.049a8.565 8.565 0 0 1-.011 3.705z"/>`),
		g.Group(children),
	)
}