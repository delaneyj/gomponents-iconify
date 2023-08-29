package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.11 1.88a1.27 1.27 0 0 0-.9-.38h-3v-.3a1.25 1.25 0 0 0-2.5 0v.3h-3a1.25 1.25 0 0 0-1.17 1.29l.41 12A1.26 1.26 0 0 0 4.2 16h7.59A1.24 1.24 0 0 0 13 14.79l.42-12a1.28 1.28 0 0 0-.31-.91zM4.2 14.75l-.32-9.5h8.24l-.33 9.5zM12.16 4H3.84V2.75h8.42z"/><path fill="currentColor" d="M5 7h1.25v5H5zm2.37 0h1.25v5H7.37zm2.38 0H11v5H9.75z"/>`),
		g.Group(children),
	)
}