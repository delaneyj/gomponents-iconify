package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFinish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 0a3.5 3.5 0 0 0-3.5 3.5V93h-9.57a3.5 3.5 0 0 0-3.5 3.5a3.5 3.5 0 0 0 3.5 3.5h25a3.5 3.5 0 0 0 3.5-3.5a3.5 3.5 0 0 0-3.5-3.5H53.5V47h43a3.5 3.5 0 0 0 3.5-3.5v-40A3.5 3.5 0 0 0 96.5 0h-46a3.5 3.5 0 0 0-.254.01A3.5 3.5 0 0 0 50 0zm13.8 7h9.8v7.43h9.8V7H93v7.43h-9.6v9.799H93v9.8h-9.6V40h-9.8v-5.97h-9.8V40H54v-5.97h9.8v-9.801H54v-9.8h9.8V7zm0 7.43v9.799h9.8v-9.8h-9.8zm9.8 9.799v9.8h9.8v-9.8h-9.8z" color="currentColor"/>`),
		g.Group(children),
	)
}