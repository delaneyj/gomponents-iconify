package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoisonNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPoisonNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM17 5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2v4a7 7 0 0 1 7 7v20a3 3 0 0 1-3 3H16a3 3 0 0 1-3-3V21a7 7 0 0 1 7-7v-4h-2a1 1 0 0 1-1-1V5Zm5 5v5a1 1 0 0 1-1 1h-1a5.002 5.002 0 0 0-4.9 4h17.8a5.002 5.002 0 0 0-4.9-4h-1a1 1 0 0 1-1-1v-5h-4Zm5.536 19.829l.707.707l1.414-1.415l-2.829-2.828l-1.414 1.414l.708.708L24 30.537l-2.121-2.122l.707-.707l-1.414-1.414l-.708.707l-1.414 1.414l-.707.707l1.414 1.414l.707-.707l2.122 2.122l-2.121 2.12l-.708-.707l-1.414 1.415l.708.707l1.414 1.415l.707.706l1.414-1.414l-.707-.707L24 33.365l2.121 2.121l-.707.707l1.414 1.415l2.829-2.829l-1.414-1.414l-.707.707l-2.122-2.121l2.122-2.122Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPoisonNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}