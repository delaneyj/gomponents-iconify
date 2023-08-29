package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M302.933 407.885H512V467.2H302.933v-59.315zm-68.019-239.352H119.467V44.8H0v366.933h234.914c38.763 0 70.187-33.211 70.187-71.974V237.184c0-38.767-31.424-68.65-70.187-68.65z"/>`),
		g.Group(children),
	)
}