package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 5h3v1H0V5zm5-5h1v3H5V0zm1 11H5V8.5l1 1zm5-5H9.5l-1-1H11zM3.131 7.161l.707.707l-2.97 2.97l-.707-.707l2.97-2.97zm7-7l.707.707l-2.97 2.97l-.707-.707l2.97-2.97zM.836.199l3.465 3.465l-.707.707L.129.906L.836.199zM6.1 4.1L4 6.1l9.8 9.9l2.2-2.1l-9.9-9.8zm0 1.4L8.5 8l-.6.6l-2.5-2.5l.7-.6z"/>`),
		g.Group(children),
	)
}