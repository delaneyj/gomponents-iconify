package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdSave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M362.7 64h-256C83 64 64 83.2 64 106.7v298.7c0 23.5 19 42.7 42.7 42.7h298.7c23.5 0 42.7-19.2 42.7-42.7v-256L362.7 64zM256 405.3c-35.4 0-64-28.6-64-64s28.6-64 64-64 64 28.6 64 64-28.6 64-64 64zM320 192H106.7v-85.3H320V192z" fill="currentColor"/>`),
		g.Group(children),
	)
}