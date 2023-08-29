package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path d="M0 0h213.3v480H0z"/><path fill="#ffd90c" d="M213.3 0h213.4v480H213.3z"/><path fill="#f31830" d="M426.7 0H640v480H426.7z"/></g>`),
		g.Group(children),
	)
}