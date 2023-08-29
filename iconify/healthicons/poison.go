package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2v4a7 7 0 0 0-7 7v20a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V21a7 7 0 0 0-7-7v-4h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H18Zm4 11v-5h4v5a1 1 0 0 0 1 1h1a5.002 5.002 0 0 1 4.9 4H15.1a5.002 5.002 0 0 1 4.9-4h1a1 1 0 0 0 1-1Zm6.243 15.536l-.707-.707l-2.122 2.122l2.122 2.121l.707-.707l1.414 1.414l-2.829 2.829l-1.414-1.415l.707-.707L24 33.365l-2.12 2.12l.706.708l-1.414 1.414l-.707-.707l-1.415-1.414l-.707-.707l1.414-1.415l.708.708l2.12-2.121l-2.12-2.122l-.708.707l-1.414-1.414l.707-.707l1.414-1.414l.708-.707l1.414 1.414l-.707.707L24 30.537l2.122-2.122l-.708-.708l1.414-1.414l2.829 2.828l-1.414 1.415Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}