package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M45 512h26q34 0 43-36l38-224l41 226q3 16 14.5 26t27.5 10h24q17 0 29.5-12.5T301 471V43q0-18-12.5-30.5T259 0H45Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512zm188-43l-55-320q-5-21-26-21t-26 21L71 469H45V149q41 0 62.5-20T131 87v-2h64v2q2 19 18 37t46 23v322h-26zm26-364q-18-6-22-20h22v20zM45 43h214v21H45V43zm0 42h43q-6 22-43 22V85z"/>`),
		g.Group(children),
	)
}