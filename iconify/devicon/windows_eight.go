package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00ADEF" d="m126 1.637l-67 9.834v49.831l67-.534zM1.647 66.709l.003 42.404l50.791 6.983l-.04-49.057zm56.82.68l.094 49.465l67.376 9.509l.016-58.863zM1.61 19.297l.047 42.383l50.791-.289l-.023-49.016z"/>`),
		g.Group(children),
	)
}