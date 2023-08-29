package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MlOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="red" d="M340.6 0H512v512H340.6z"/><path fill="#009a00" d="M0 0h170.3v512H0z"/><path fill="#ff0" d="M170.3 0h171.2v512H170.3z"/></g>`),
		g.Group(children),
	)
}