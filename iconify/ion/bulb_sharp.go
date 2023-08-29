package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulbSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M208 464h96v32h-96zm-16-48h128v32H192zM369.42 62.69C339.35 32.58 299.07 16 256 16A159.62 159.62 0 0 0 96 176c0 46.62 17.87 90.23 49 119.64l4.36 4.09C167.37 316.57 192 339.64 192 360v40h48V269.11L195.72 244L214 217.72L256 240l41.29-22.39l19.1 25.68l-44.39 26V400h48v-40c0-19.88 24.36-42.93 42.15-59.77l4.91-4.66C399.08 265 416 223.61 416 176a159.16 159.16 0 0 0-46.58-113.31Z"/>`),
		g.Group(children),
	)
}