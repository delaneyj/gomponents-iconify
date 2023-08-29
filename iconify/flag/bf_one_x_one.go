package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BfOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#de0000" d="M512 511.6H.5V0H512z"/><path fill="#35a100" d="M511.8 512H0V256.2h511.7z"/></g><path fill="#fff300" fill-rule="evenodd" d="m389 223.8l-82.9 56.5l31.7 91.6l-82.7-56.7l-82.8 56.7l31.7-91.6l-82.8-56.6l102.3.2l31.6-91.7l31.5 91.6"/>`),
		g.Group(children),
	)
}