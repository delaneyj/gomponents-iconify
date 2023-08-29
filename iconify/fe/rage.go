package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feRage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRage1" fill="currentColor"><path id="feRage2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-5a4 4 0 1 0-8 0h1.333s.423-2.667 2.667-2.667c2.244 0 2.661 2.667 2.661 2.667H16Zm-1.5-6A1.5 1.5 0 0 0 16 9.5V8c-.828 0-3 .672-3 1.5a1.5 1.5 0 0 0 1.5 1.5Zm-5 0A1.5 1.5 0 0 0 11 9.5C11 8.672 8.828 8 8 8v1.5A1.5 1.5 0 0 0 9.5 11Z"/></g></g>`),
		g.Group(children),
	)
}