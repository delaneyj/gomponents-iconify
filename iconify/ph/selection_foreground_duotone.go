package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionForegroundDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 96v112a8 8 0 0 1-8 8H48a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h112a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M64 216a8 8 0 0 1-8 8h-8a16 16 0 0 1-16-16v-8a8 8 0 0 1 16 0v8h8a8 8 0 0 1 8 8Zm48-8H96a8 8 0 0 0 0 16h16a8 8 0 0 0 0-16Zm-72-40a8 8 0 0 0 8-8v-16a8 8 0 0 0-16 0v16a8 8 0 0 0 8 8Zm128 24a8 8 0 0 0-8 8v8h-8a8 8 0 0 0 0 16h8a16 16 0 0 0 16-16v-8a8 8 0 0 0-8-8Zm0-80a8 8 0 0 0 8-8v-8a16 16 0 0 0-16-16h-8a8 8 0 0 0 0 16h8v8a8 8 0 0 0 8 8ZM56 80h-8a16 16 0 0 0-16 16v8a8 8 0 0 0 16 0v-8h8a8 8 0 0 0 0-16Zm152-48H96a16 16 0 0 0-16 16v40a4.44 4.44 0 0 0 0 .55A8 8 0 0 0 88 96h24a8 8 0 0 0 0-16H96V48h112v112h-32v-16a8 8 0 0 0-16 0v24a8 8 0 0 0 8 8h40a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Z"/></g>`),
		g.Group(children),
	)
}