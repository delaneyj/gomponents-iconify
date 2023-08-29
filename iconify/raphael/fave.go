package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.132 7.97c-2.203-2.204-5.916-2.097-8.25.236l-.382.382l-.382-.382c-2.334-2.333-6.047-2.44-8.25-.235c-2.204 2.204-2.098 5.917.235 8.25l8.396 8.396l8.395-8.396c2.334-2.333 2.44-6.046.237-8.25z"/>`),
		g.Group(children),
	)
}