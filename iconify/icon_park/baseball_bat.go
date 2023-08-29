package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseballBat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-width="4"><circle cx="40" cy="40" r="3"/><path stroke-linejoin="round" d="M16.502 9.43095C16.502 9.43095 26.4998 22 37.4998 37.5C21.4998 26 9.43102 16.5021 9.43102 16.5021C9.43102 16.5021 3.11056 10.8894 6.99983 7.00032C10.8891 3.11124 16.502 9.43095 16.502 9.43095Z"/></g>`),
		g.Group(children),
	)
}