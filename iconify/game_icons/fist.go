package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m329.8 235.69l62.83-82.71l42.86 32.56l-62.83 82.75zm-12.86-9.53l66.81-88l-45-34.15l-66.81 88zm-27.48-97.78l-19.3 39.57l57-75l-42.51-32.3l-36.24 47.71zm-20.74-73.24l-46.64-35.43l-42 55.31l53.67 26.17zm107 235.52l-139-102.71l-9.92.91l4.56 2l62.16 138.43l-16.52 2.25l-57.68-128.5l-40-17.7l-4-30.84l39.41 19.42l36.36-3.33l17-34.83l-110.9-54.09l-80.68 112.51L177.6 346.67l-22.7 145.62H341V372.62l35.29-48.93L387 275.77z"/>`),
		g.Group(children),
	)
}