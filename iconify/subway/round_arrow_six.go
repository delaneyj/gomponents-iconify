package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundArrowSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M110.5 131.8C145.6 90.5 197.5 64 256 64c106.1 0 192 85.9 192 192h64C512 114.6 397.4 0 256 0C179.8 0 111.7 33.4 64.9 86.2L0 21.3V192h170.7l-60.2-60.2zm291 248.4c-35.2 41.3-87 67.8-145.5 67.8c-106.1 0-192-85.9-192-192H0c0 141.4 114.6 256 256 256c76.2 0 144.3-33.4 191.1-86.2l64.9 64.9V320H341.3l60.2 60.2zM213.3 256c0 23.6 19.1 42.7 42.7 42.7s42.7-19.1 42.7-42.7s-19.1-42.7-42.7-42.7s-42.7 19.1-42.7 42.7z"/>`),
		g.Group(children),
	)
}