package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeMaskNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsPpeMaskNegative0)"><path d="m39.913 29.61l-4.303 1.56c1.735-2.762 3.053-6.642 3.334-12.09l1.315.318a2 2 0 0 1 1.52 2.118l-.556 6.387a2 2 0 0 1-1.31 1.706Zm-27.73 1.221c-1.628-2.73-2.851-6.501-3.125-11.708l-1.254.278a2 2 0 0 0-1.557 2.142l.573 6.021a2 2 0 0 0 1.266 1.674l4.097 1.593ZM24 18.47l-6.242 1.56a1 1 0 1 0 .485 1.94L24 20.53l5.758 1.44a1 1 0 1 0 .485-1.94L24 18.47Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm20.16 12.642S9 12.642 9 16.868v.219l-1.63.362a4 4 0 0 0-3.114 4.284l.573 6.02a4 4 0 0 0 2.532 3.35l7.192 2.796C18.651 38 23.516 38 23.516 38h.968s4.794 0 8.872-4.01a.995.995 0 0 0 .485-.05l6.754-2.45a4 4 0 0 0 2.62-3.413l.557-6.388a4 4 0 0 0-3.041-4.234L39 17.035v-.167c0-4.226-10.986-4.226-10.986-4.226l-.9-1.616A2 2 0 0 0 25.367 10h-2.56a2 2 0 0 0-1.746 1.026l-.9 1.616Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeMaskNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}