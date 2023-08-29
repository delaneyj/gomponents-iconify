package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fbdislike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 0H576l-64 64H320V0H0v576h320v-64h64l96 96l160 352h192V800l-64-96v-64h160l96-64V64zM256 512H64V64h192v448zm704 21l-64 43H704v160l64 96v64h-79L544 576L416 448h-96V128h224l64-64h288l64 43v426z"/>`),
		g.Group(children),
	)
}