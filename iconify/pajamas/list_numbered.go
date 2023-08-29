package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h2v4H1V3H0V2Zm1.637 9.008H0v-1h1.637a1.382 1.382 0 0 1 .803 2.506L1.76 13H3v1H0v-.972L1.859 11.7a.382.382 0 0 0-.222-.693ZM4.75 3a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5h-9.5Zm0 4.25a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5h-9.5Zm-.75 5a.75.75 0 0 1 .75-.75h9.5a.75.75 0 0 1 0 1.5h-9.5a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}