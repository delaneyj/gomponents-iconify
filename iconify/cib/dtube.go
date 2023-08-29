package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.135 7.333a13.218 13.218 0 0 0-5.177-5.385A14.89 14.89 0 0 0 17.437.052H.239l9.214 5.333h7.984c2.589 0 4.547.828 5.896 2.51c1.333 1.667 2 4.104 2 7.344v1.693c-.016 3.12-.708 5.521-2.068 7.203c-1.359 1.677-3.333 2.516-5.891 2.516H9.145l-9.146 5.281h17.505c2.786 0 5.297-.651 7.51-1.917c2.214-1.271 3.932-3.068 5.156-5.365c1.229-2.292 1.828-4.932 1.828-7.906V15.28c0-2.964-.625-5.599-1.865-7.948zM.172 5.281v21.464L18.761 16z"/>`),
		g.Group(children),
	)
}