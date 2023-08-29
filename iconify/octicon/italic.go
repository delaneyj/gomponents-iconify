package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M2.81 5h1.98L3 14H1l1.81-9zm.36-2.7c0-.7.58-1.3 1.33-1.3c.56 0 1.13.38 1.13 1.03c0 .75-.59 1.3-1.33 1.3c-.58 0-1.13-.38-1.13-1.03z" fill="currentColor"/>`),
		g.Group(children),
	)
}