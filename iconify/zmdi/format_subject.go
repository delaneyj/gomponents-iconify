package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSubject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 299v42H0v-42h213zm128-171v43H0v-43h341zM0 256v-43h341v43H0zM0 43h341v42H0V43z"/>`),
		g.Group(children),
	)
}