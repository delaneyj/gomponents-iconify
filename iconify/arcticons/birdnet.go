package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Birdnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 25.445l27.908 5.49l5.317 4.372m2.36-1.561l-2.36 1.561l-3.299 2.183M20.588 10.51l-5.669 3.232l.032.85l-1.894 1.462l2.81.46v4.648l2.433 2.989l3.85 3.617l8.03-.778l1.261.025L43.5 33.869l-10.136-8.186l3.892.396l-2.83-2.459l4.303-.268l-17.499-9.563Zm.547 16.43l-.83 1.49m5.435-.896l-1.246 1.664"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.068 20.312l-7.715-.26l-.501-3.858Z"/>`),
		g.Group(children),
	)
}