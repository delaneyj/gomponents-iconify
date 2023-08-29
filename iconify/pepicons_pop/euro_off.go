package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.489 4C9.43 4 7 6.213 7 9.5c0 3.386 2.527 6 5.489 6c.743 0 1.451-.161 2.098-.454a1 1 0 1 1 .826 1.821a7.061 7.061 0 0 1-2.924.633C8.283 17.5 5 13.845 5 9.5C5 5.055 8.38 2 12.489 2c1.237 0 2.428.393 3.574 1.174a1 1 0 1 1-1.126 1.652C14.08 4.242 13.274 4 12.489 4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 8a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm0 3.5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}