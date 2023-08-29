package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#321B41" d="M28.21 2.4a3.5 3.5 0 0 1 1.344 4.763l-11.51 20.554a4 4 0 0 1-6.135 1.047L3.186 21.08a3.5 3.5 0 0 1 4.628-5.252l5.936 5.23l9.696-17.314A3.5 3.5 0 0 1 28.21 2.4Z"/>`),
		g.Group(children),
	)
}