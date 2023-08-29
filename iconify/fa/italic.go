package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m0 1534l17-85q22-7 61.5-16.5t72-19T210 1390q28-35 41-101q1-7 62-289t114-543.5T479 160v-25q-24-13-54.5-18.5t-69.5-8t-58-5.5L316 0q33 2 120 6.5t149.5 7T706 16q48 0 98.5-2.5t121-7T1024 0q-5 39-19 89q-30 10-101.5 28.5T795 151q-8 19-14 42.5t-9 40t-7.5 45.5t-6.5 42q-27 148-87.5 419.5T593 1096q-2 9-13 58t-20 90t-16 83.5t-6 57.5l1 18q17 4 185 31q-3 44-16 99q-11 0-32.5 1.5T643 1536q-29 0-87-10t-86-10q-138-2-206-2q-51 0-143 9T0 1534z"/>`),
		g.Group(children),
	)
}