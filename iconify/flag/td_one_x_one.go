package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TdOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#002664" d="M0 0h171.2v512H0z"/><path fill="#c60c30" d="M340.8 0H512v512H340.8z"/><path fill="#fecb00" d="M171.2 0h169.6v512H171.2z"/></g>`),
		g.Group(children),
	)
}