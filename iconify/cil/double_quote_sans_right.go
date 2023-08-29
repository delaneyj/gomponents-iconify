package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleQuoteSansRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M280 185.143V416h216V16h-38.4ZM464 384H312V198.857L464 54.1ZM232 16h-38.4L16 185.143V416h216Zm-32 368H48V198.857L200 54.1Z"/>`),
		g.Group(children),
	)
}