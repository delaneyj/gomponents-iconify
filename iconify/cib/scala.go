package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scala(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.12 32c6.047 0 18.411-2.021 19.76-4v-7.641c-1.276 1.88-13.714 3.885-19.76 3.885zm0-10.182c6.047 0 18.411-2.021 19.76-4v-7.635c-1.276 1.875-13.714 3.88-19.76 3.88zm0-10.177c6.047 0 18.411-2.021 19.76-4V0C24.604 1.875 12.166 3.88 6.12 3.88z"/>`),
		g.Group(children),
	)
}