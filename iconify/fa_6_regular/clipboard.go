package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M280 64h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h49.6C121 27.5 153.3 0 192 0s71 27.5 78.4 64h9.6zM64 112c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h256c8.8 0 16-7.2 16-16V128c0-8.8-7.2-16-16-16h-16v24c0 13.3-10.7 24-24 24H104c-13.3 0-24-10.7-24-24v-24H64zm128-8a24 24 0 1 0 0-48a24 24 0 1 0 0 48z"/>`),
		g.Group(children),
	)
}