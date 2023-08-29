package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyNervousLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180.24 163.76a6 6 0 1 1-8.48 8.48L160 160.49l-11.76 11.75a6 6 0 0 1-8.48 0L128 160.49l-11.76 11.75a6 6 0 0 1-8.48 0L96 160.49l-11.76 11.75a6 6 0 0 1-8.48-8.48l16-16a6 6 0 0 1 8.48 0L112 159.51l11.76-11.75a6 6 0 0 1 8.48 0L144 159.51l11.76-11.75a6 6 0 0 1 8.48 0ZM230 128A102 102 0 1 1 128 26a102.12 102.12 0 0 1 102 102Zm-12 0a90 90 0 1 0-90 90a90.1 90.1 0 0 0 90-90ZM92 118a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm72 0a10 10 0 1 0-10-10a10 10 0 0 0 10 10Z"/>`),
		g.Group(children),
	)
}