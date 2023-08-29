package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.875 21.425L11.1 14.6q-.5.2-1.012.3T9 15q-2.5 0-4.25-1.75T3 9q0-.9.25-1.712t.7-1.538L7.6 9.4l1.8-1.8l-3.65-3.65q.725-.45 1.538-.7T9 3q2.5 0 4.25 1.75T15 9q0 .575-.1 1.088t-.3 1.012l6.825 6.775l-3.55 3.55Zm0-2.85l.675-.675l-6.4-6.4q.45-.5.65-1.163T13 9q0-1.5-.963-2.613T9.65 5.05L11.5 6.9q.3.3.3.7t-.3.7l-3.2 3.2q-.3.3-.7.3t-.7-.3L5.05 9.65q.225 1.425 1.338 2.388T9 13q.65 0 1.3-.2t1.175-.625l6.4 6.4ZM11.8 11.8Z"/>`),
		g.Group(children),
	)
}