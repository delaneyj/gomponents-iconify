package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015 331l-681 684q-10 9-23 9t-23-9q-83-84-123.5-187t-37-202T170 435t108-160q69-67 161-105t190.5-42t201 36T1015 286q9 9 9 22.5t-9 22.5zM61 753Q-20 592 7 426.5T150 147T432 7t326 54q-78 1-143 7T476.5 90t-138 49T225 220q-52 52-86.5 114.5T88 469.5T66.5 606T61 753z"/>`),
		g.Group(children),
	)
}