package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.3 0-.563-.163t-.387-.512L6 6.075l3.5.85l1.5-1.15V3.4q0-1 .725-1.7T13.5 1q1.05 0 1.775.7T16 3.4V4h-2v-.6q0-.2-.15-.338t-.35-.137q-.2 0-.35.137T13 3.4v2.375l1.5 1.15l3.5-.85L12.95 21.3q-.125.35-.388.525T12 22Zm1-7.2l1.95-5.95l-.9.225L13 8.3v6.5Zm-2 0V8.3l-1.05.8l-.925-.25L11 14.8Z"/>`),
		g.Group(children),
	)
}