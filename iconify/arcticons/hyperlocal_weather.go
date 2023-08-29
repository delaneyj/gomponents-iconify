package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HyperlocalWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.265 25.691c0-6.87-9.156-21.191-9.156-21.191S5.953 18.82 5.953 25.691a9.156 9.156 0 1 0 18.312 0ZM42.047 12.45c0-2.404-3.203-7.413-3.203-7.413s-3.202 5.009-3.202 7.412a3.203 3.203 0 1 0 6.405 0Zm-3.948 15.678c0-3.395-4.525-10.473-4.525-10.473s-4.525 7.078-4.525 10.473a4.525 4.525 0 0 0 9.05 0Zm-8.74 12.319c0-2.29-3.053-7.065-3.053-7.065s-3.052 4.775-3.052 7.065a3.053 3.053 0 1 0 6.105 0Z"/>`),
		g.Group(children),
	)
}