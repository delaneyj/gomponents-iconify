package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPrinterCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#fff" fill-rule="evenodd" d="M8 10V6c0-.552.373-1 .833-1h8.334c.46 0 .833.448.833 1v4" clip-rule="evenodd"/><path fill="#000" d="M9 6v4a1 1 0 0 1-2 0V6c0-1.068.776-2 1.833-2h8.334C18.224 4 19 4.932 19 6v4a1 1 0 1 1-2 0V6H9Z"/><path fill="#000" d="M7 9h12a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-5a3 3 0 0 1 3-3Z"/><path fill="#fff" fill-rule="evenodd" d="M8 15v4.8c0 .663.373 1.2.833 1.2h8.334c.46 0 .833-.537.833-1.2V15" clip-rule="evenodd"/><path fill="#000" d="M9 19.8V15a1 1 0 1 0-2 0v4.8c0 1.154.727 2.2 1.833 2.2h8.334C18.273 22 19 20.954 19 19.8V15a1 1 0 1 0-2 0v4.8a.661.661 0 0 1-.029.2H9.029A.66.66 0 0 1 9 19.8Z"/><path fill="#000" d="M10 19a.5.5 0 0 1 0-1h6a.5.5 0 0 1 0 1h-6Zm0-2a.5.5 0 0 1 0-1h6a.5.5 0 0 1 0 1h-6Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPrinterCircleFilled0)"/></g>`),
		g.Group(children),
	)
}