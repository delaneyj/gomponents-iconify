package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M609 448q0 35 9 83q1 5 16.5 35.5t27 73.5t11.5 96h-64q0-1-1.5-18t-4-35.5t-6-40t-8.5-40t-12-26.5q-23-24-27.5-47.5T545 448q0-80 37.5-136t90.5-56q15 0 32 6q-42 17-69 68t-27 118zM417 204q0 38 9 89q55 62 55 168q0 175-32 243q-32-145-32-243q0-77-9-89q-55-62-55-168q0-84 37.5-144T481 0q15 0 32 7q-42 17-69 72t-27 125zm429 697q-21 50-54.5 86.5T721 1024H305q-37 0-71-37t-55-88q-17-10-51-29t-51-29t-38-24.5t-30-26T1 768h1024q0 15-9 29t-29.5 28t-38.5 23.5t-50.5 26T846 901z"/>`),
		g.Group(children),
	)
}