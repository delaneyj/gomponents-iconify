package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteBlankBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h108.69a19.86 19.86 0 0 0 14.14-5.86l51.31-51.31a19.86 19.86 0 0 0 5.86-14.14V48a20 20 0 0 0-20-20ZM52 52h152v92h-48a12 12 0 0 0-12 12v48H52Zm116 139v-23h23Z"/>`),
		g.Group(children),
	)
}