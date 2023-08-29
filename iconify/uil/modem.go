package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.48 13.13a.65.65 0 0 0-.05-.2a.89.89 0 0 0-.08-.17a.86.86 0 0 0-.1-.16l-.16-.13l-.09-.09l-14.72-8.5a1 1 0 0 0-1 1.74l11.49 6.63H3.5a1 1 0 0 0-1 1v4a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3v-4s-.02-.08-.02-.12Zm-2 4.12a1 1 0 0 1-1 1H5.5a1 1 0 0 1-1-1v-3h15Zm-3 0a1 1 0 1 0-1-1a1 1 0 0 0 1.02 1Z"/>`),
		g.Group(children),
	)
}