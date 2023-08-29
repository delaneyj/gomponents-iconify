package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-73.14 14.86a8 8 0 0 1 11.32 0l45.25 45.26a8 8 0 0 1-11.31 11.31l-45.26-45.25a8 8 0 0 1 0-11.32Zm-34.68 51.91a8 8 0 0 1 10.45-4.33l59.13 24.49a8 8 0 0 1-3.06 15.4a7.89 7.89 0 0 1-3.06-.62l-59.13-24.49a8 8 0 0 1-4.33-10.45ZM96 152h64a8 8 0 0 1 0 16H96a8 8 0 0 1 0-16Zm104 40a8 8 0 0 1-8 8H64a8 8 0 0 1-8-8v-48a8 8 0 0 1 16 0v40h112v-40a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}