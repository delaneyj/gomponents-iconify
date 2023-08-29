package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 36V22H6V36C6 38.2091 7.79086 40 10 40H38C40.2091 40 42 38.2091 42 36Z"/><path stroke="#fff" d="M6 30H42"/><path stroke="#000" d="M6 26V34"/><path stroke="#000" d="M42 26V34"/><path stroke="#000" d="M6.06657 22H42C42 19.9844 30.1933 11.4184 24.0333 7.38735C17.8733 3.35628 12.7399 8.89906 8.6332 13.9379C5.34786 17.969 6.06659 19.9845 6.06657 22Z"/></g>`),
		g.Group(children),
	)
}