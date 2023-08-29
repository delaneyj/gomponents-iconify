package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMoreThree0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#fff" fill-rule="evenodd" d="M21.5 14a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Zm0 10a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0ZM24 36.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMoreThree0)"/>`),
		g.Group(children),
	)
}