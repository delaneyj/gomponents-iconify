package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 152a32 32 0 1 1-32-32a32 32 0 0 1 32 32Z" opacity=".2"/><path d="M212 152a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-4.55 39.29A48.08 48.08 0 0 1 160 232h-24a48.05 48.05 0 0 1-48-48v-40.51A64 64 0 0 1 32 80V40a16 16 0 0 1 16-16h16a8 8 0 0 1 0 16H48v40a48 48 0 0 0 48.64 48c26.11-.34 47.36-22.25 47.36-48.83V40h-16a8 8 0 0 1 0-16h16a16 16 0 0 1 16 16v39.17c0 32.84-24.53 60.29-56 64.31V184a32 32 0 0 0 32 32h24a32.06 32.06 0 0 0 31.22-25a40 40 0 1 1 16.23.27ZM224 152a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z"/></g>`),
		g.Group(children),
	)
}