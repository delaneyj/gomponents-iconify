package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoSlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.042 19.166a2.521 2.521 0 1 1-2.52-2.521h2.52zm1.27 0a2.521 2.521 0 0 1 5.043 0v6.313a2.521 2.521 0 1 1-5.042 0zm2.522-10.124a2.521 2.521 0 1 1 2.521-2.52v2.52zm0 1.27a2.521 2.521 0 0 1 0 5.043H6.52a2.521 2.521 0 1 1 0-5.042zm10.124 2.522a2.521 2.521 0 1 1 2.52 2.521h-2.52zm-1.27 0a2.521 2.521 0 0 1-5.043 0V6.52a2.521 2.521 0 1 1 5.042 0zm-2.522 10.124a2.521 2.521 0 1 1-2.521 2.52v-2.52zm0-1.27a2.521 2.521 0 0 1 0-5.043h6.313a2.521 2.521 0 1 1 0 5.042z"/>`),
		g.Group(children),
	)
}