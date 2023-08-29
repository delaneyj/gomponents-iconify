package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoisonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2v4a7 7 0 0 1 7 7v20a3 3 0 0 1-3 3H16a3 3 0 0 1-3-3V21a7 7 0 0 1 7-7v-4h-2a1 1 0 0 1-1-1V5Zm5 5h4v5a1 1 0 0 0 1 1h1a5.002 5.002 0 0 1 4.9 4H15.1a5.002 5.002 0 0 1 4.9-4h1a1 1 0 0 0 1-1v-5Zm-7 12v19a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V22H15Zm4-16v2h10V6H19Zm1.464 23.83l2.122 2.12l-2.121 2.122l-.708-.708l-1.414 1.415l.708.707l1.414 1.415l.707.706l1.414-1.414l-.707-.707L24 33.365l2.121 2.121l-.707.707l1.414 1.415l2.829-2.829l-1.414-1.414l-.707.707l-2.122-2.121l2.122-2.122l.707.707l1.414-1.415l-2.829-2.828l-1.414 1.414l.708.708L24 30.537l-2.121-2.122l.707-.707l-1.414-1.414l-.708.707l-1.414 1.414l-.707.707l1.414 1.414l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}