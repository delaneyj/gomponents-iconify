package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M518.474 105.344C305.831 120.286.168 154.236 0 570.687v523.97h474.504v-560.61H316.946C306.965 384.354 430.23 345.701 564.274 316.03l-45.8-210.686zm635.724 0c-212.643 14.942-518.306 48.893-518.473 465.343v523.97h474.505v-560.61H952.672C942.689 384.354 1065.956 345.701 1200 316.03l-45.802-210.686z"/>`),
		g.Group(children),
	)
}