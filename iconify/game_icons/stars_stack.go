package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarsStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256.65 38.984l-49.697 100.702l-111.13 16.146l80.413 78.385l-18.982 110.68l99.396-52.256l99.397 52.256l-18.983-110.68l80.413-78.384l-111.127-16.146l-49.7-100.702zM112 308.826l-26.674 54.05l-59.646 8.665l43.16 42.073l-10.188 59.403L112 444.97l53.348 28.046l-10.188-59.403l43.16-42.072l-59.646-8.665L112 308.825zm288 0l-26.674 54.05l-59.646 8.665l43.16 42.073l-10.188 59.403L400 444.97l53.348 28.046l-10.188-59.403l43.16-42.072l-59.646-8.665L400 308.825z"/>`),
		g.Group(children),
	)
}