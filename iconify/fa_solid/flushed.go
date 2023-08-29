package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flushed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M344 200c-13.3 0-24 10.7-24 24s10.7 24 24 24s24-10.7 24-24s-10.7-24-24-24zm-192 0c-13.3 0-24 10.7-24 24s10.7 24 24 24s24-10.7 24-24s-10.7-24-24-24zM248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8zM80 224c0-39.8 32.2-72 72-72s72 32.2 72 72s-32.2 72-72 72s-72-32.2-72-72zm232 176H184c-21.2 0-21.2-32 0-32h128c21.2 0 21.2 32 0 32zm32-104c-39.8 0-72-32.2-72-72s32.2-72 72-72s72 32.2 72 72s-32.2 72-72 72z"/>`),
		g.Group(children),
	)
}