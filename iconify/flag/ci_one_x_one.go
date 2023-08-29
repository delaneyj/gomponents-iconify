package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CiOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#00cd00" d="M341.5 0H512v512H341.5z"/><path fill="#ff9a00" d="M0 0h170.3v512H0z"/><path fill="#fff" d="M170.3 0h171.2v512H170.3z"/></g>`),
		g.Group(children),
	)
}