package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M441 99h99c18 0 32 14 32 32v23c0 15-11 29-25 32v475c0 31-26 56-58 56H83c-32 0-57-25-57-56V186c-14-3-26-17-26-32v-23c0-18 15-32 33-32h98V57c0-31 25-57 57-57h196c32 0 57 26 57 57v42zm-57-49H188c-5 0-8 3-8 7v42h212V57c0-4-3-7-8-7zM99 653h375c4 0 7-5 7-9V187H91v457c0 4 3 9 8 9zm77-37h-6c-18 0-33-15-33-33V256c0-17 15-31 33-31h6c17 0 32 14 32 31v327c0 18-15 33-32 33zm113 0h-5c-18 0-33-15-33-33V256c0-17 15-31 33-31h5c18 0 33 14 33 31v327c0 18-15 33-33 33zm113 0h-5c-18 0-32-15-32-33V256c0-17 14-31 32-31h5c18 0 33 14 33 31v327c0 18-15 33-33 33z"/>`),
		g.Group(children),
	)
}