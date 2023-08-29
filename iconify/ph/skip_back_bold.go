package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.75 30.52a20 20 0 0 0-20.3.53L68 102V40a12 12 0 0 0-24 0v176a12 12 0 0 0 24 0v-62l113.45 71A20 20 0 0 0 212 208.12V47.88a19.86 19.86 0 0 0-10.25-17.36ZM188 200.73L71.7 128L188 55.27Z"/>`),
		g.Group(children),
	)
}