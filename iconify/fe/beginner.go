package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBeginner0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBeginner1" fill="currentColor" fill-rule="nonzero"><path id="feBeginner2" d="M12 7.529L5 5.09v10.372l7 3.251V7.529ZM5.632 3.108L12 5.326l6.368-2.218c1.047-.365 2.18.227 2.53 1.322c.067.213.102.436.102.66v10.372c0 .826-.465 1.574-1.188 1.91L12 21l-7.812-3.628C3.465 17.036 3 16.288 3 15.462V5.09C3 3.936 3.895 3 5 3c.215 0 .429.037.632.108Z"/></g></g>`),
		g.Group(children),
	)
}