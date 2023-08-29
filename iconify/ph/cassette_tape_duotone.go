package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CassetteTapeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M170.4 171.2L192 200H64l21.6-28.8A8 8 0 0 1 92 168h72a8 8 0 0 1 6.4 3.2ZM176 88a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm-72 24a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z" opacity=".2"/><path d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM80 192l12-16h72l12 16Zm144 0h-28l-19.2-25.6A16.07 16.07 0 0 0 164 160H92a16.07 16.07 0 0 0-12.8 6.4L60 192H32V64h192v128ZM176 80H80a32 32 0 0 0 0 64h96a32 32 0 0 0 0-64Zm-27.7 16a31.92 31.92 0 0 0 0 32h-40.6a31.92 31.92 0 0 0 0-32ZM64 112a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm112 16a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/></g>`),
		g.Group(children),
	)
}