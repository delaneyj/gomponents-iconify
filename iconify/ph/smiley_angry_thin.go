package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyAngryThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 140a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm64-8a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm64-4A100 100 0 1 1 128 28a100.11 100.11 0 0 1 100 100Zm-8 0a92 92 0 1 0-92 92a92.1 92.1 0 0 0 92-92Zm-46.22-43.33L128 115.19L82.22 84.67a4 4 0 1 0-4.44 6.66l48 32a4 4 0 0 0 4.44 0l48-32a4 4 0 1 0-4.44-6.66Zm-19.57 96c-7.82-5.2-15.27-8.67-26.21-8.67s-18.39 3.47-26.21 8.67a4 4 0 1 0 4.42 6.66C113.2 182.69 119 180 128 180s14.8 2.69 21.79 7.33a4 4 0 1 0 4.42-6.66Z"/>`),
		g.Group(children),
	)
}