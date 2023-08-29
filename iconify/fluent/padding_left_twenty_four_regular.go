package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 4.5a.75.75 0 0 0-.75.75v.867a.75.75 0 0 0 1.5 0V5.25a.75.75 0 0 0-.75-.75Zm0 3.467a.75.75 0 0 0-.75.75v1.733a.75.75 0 0 0 1.5 0V8.717a.75.75 0 0 0-.75-.75Zm0 4.333a.75.75 0 0 0-.75.75v1.733a.75.75 0 0 0 1.5 0V13.05a.75.75 0 0 0-.75-.75Zm0 4.333a.75.75 0 0 0-.75.75v.867a.75.75 0 0 0 1.5 0v-.867a.75.75 0 0 0-.75-.75ZM21.25 4.5a.75.75 0 0 0-.75.75v13a.75.75 0 0 0 1.5 0v-13a.75.75 0 0 0-.75-.75ZM5.22 12.28a.75.75 0 0 1 0-1.06l5-5a.75.75 0 1 1 1.06 1.06L7.56 11h10.69a.75.75 0 0 1 0 1.5H7.56l3.72 3.72a.75.75 0 1 1-1.06 1.06l-5-5Z"/>`),
		g.Group(children),
	)
}