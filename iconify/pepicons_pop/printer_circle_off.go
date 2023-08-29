package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M9 6v4a1 1 0 0 1-2 0V6c0-1.068.776-2 1.833-2h8.334C18.224 4 19 4.932 19 6v4a1 1 0 1 1-2 0V6H9Z"/><path d="M7 9h12a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-5a3 3 0 0 1 3-3Z"/><path d="M9 19.8V15a1 1 0 1 0-2 0v4.8c0 1.154.727 2.2 1.833 2.2h8.334C18.273 22 19 20.954 19 19.8V15a1 1 0 1 0-2 0v4.8a.661.661 0 0 1-.029.2H9.029A.66.66 0 0 1 9 19.8Z"/><path d="M10 19a.5.5 0 0 1 0-1h6a.5.5 0 0 1 0 1h-6Zm0-2a.5.5 0 0 1 0-1h6a.5.5 0 0 1 0 1h-6Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}