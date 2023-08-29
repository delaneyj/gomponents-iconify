package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.8 1.2 16 2.667 16h10.666C14.801 16 16 14.8 16 13.333V2.667A2.674 2.674 0 0 0 13.333 0zM4 4h8c.143 0 .281.031.409.088L8 9.231L3.591 4.088A.982.982 0 0 1 4 4zm-1 7V5l.002-.063l2.932 3.421l-2.9 2.9A.967.967 0 0 1 3 11zm9 1H4c-.088 0-.175-.012-.258-.034L6.588 9.12l1.413 1.648L9.414 9.12l2.846 2.846a.967.967 0 0 1-.258.034zm1-1c0 .088-.012.175-.034.258l-2.9-2.9l2.932-3.421L13 5v6z"/>`),
		g.Group(children),
	)
}