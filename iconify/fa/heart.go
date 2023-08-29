package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M896 1536q-26 0-44-18L228 916q-10-8-27.5-26T145 824.5T77 727T23.5 606T0 468q0-220 127-344T478 0q62 0 126.5 21.5t120 58T820 148t76 68q36-36 76-68t95.5-68.5t120-58T1314 0q224 0 351 124t127 344q0 221-229 450l-623 600q-18 18-44 18z"/>`),
		g.Group(children),
	)
}