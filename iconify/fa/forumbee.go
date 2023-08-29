package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forumbee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M934 22Q617 143 378 384.5T20 945Q0 856 0 769q0-208 102.5-384.5t278.5-279T765 3q82 0 169 19zm269 119q93 65 164 155q-389 113-674.5 400.5T296 1373q-93-72-155-162q112-386 395-671t667-399zM470 1475q115-356 379.5-622T1469 469q40 92 54 195q-292 120-516 345t-343 518q-103-14-194-52zm1066 58q-193-50-367-115q-135 84-290 107q109-205 274-370.5T1522 879q-21 152-101 284q65 175 115 370z"/>`),
		g.Group(children),
	)
}