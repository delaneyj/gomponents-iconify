package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.3 0-.563-.163t-.387-.512l-4.775-14.4q-.1-.275.113-.487t.487-.138l2.625.625l1.5-1.15V3.4q0-1 .725-1.7T13.5 1q.9 0 1.513.475t.812 1.2q.15.55-.125.938T14.85 4q-.35 0-.575-.163t-.3-.487q-.05-.2-.175-.313t-.3-.112q-.2 0-.35.138T13 3.4v2.375l1.5 1.15l2.625-.625q.275-.075.488.138t.112.487L12.95 21.3q-.125.35-.387.525T12 22Zm1-7.2l1.95-5.95l-.9.225L13 8.3v6.5Zm-2 0V8.3l-1.05.8l-.925-.25L11 14.8Z"/>`),
		g.Group(children),
	)
}