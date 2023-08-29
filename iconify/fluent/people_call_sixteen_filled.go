package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleCallSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm8-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM8 8a1.5 1.5 0 0 1 1.5 1.5v.094c-.009.25-.188 2.906-4 2.906c-4 0-4-2.925-4-2.925V9.5A1.5 1.5 0 0 1 3 8h5Zm4.584.581l.283-.75c.258-.68 1.062-1.016 1.74-.727l.388.166c.473.202.865.568.947 1.06c.457 2.725-1.908 6.601-4.63 7.59c-.492.178-1.023.04-1.445-.246l-.346-.235a1.184 1.184 0 0 1-.204-1.79l.545-.607a1.066 1.066 0 0 1 1.034-.323l1.225.29c.971-.607 1.492-1.46 1.562-2.56l-.878-.86a.937.937 0 0 1-.221-1.008Z"/>`),
		g.Group(children),
	)
}