package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M152 341q62 0 105.5-43.5T301 192T257.5 86.5T152 43T46.5 86.5T3 192t43.5 105.5T152 341zm0-256q45 0 76 31t31 76t-31 76t-76 31t-76-31t-31-76t31-76t76-31zm43 107q0 18-12.5 30.5T152 235t-30.5-12.5T109 192t12.5-30.5T152 149t30.5 12.5T195 192z"/>`),
		g.Group(children),
	)
}