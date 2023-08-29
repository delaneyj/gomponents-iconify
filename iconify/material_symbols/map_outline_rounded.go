package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.35 20.775L9 18.9l-4.65 1.8q-.5.2-.925-.112T3 19.75v-14q0-.325.188-.575T3.7 4.8l4.65-1.575q.325-.1.65-.113t.65.113L15 5.1l4.65-1.8q.5-.2.925.113T21 4.25v14q0 .325-.187.575t-.513.375l-4.65 1.575q-.325.1-.65.113t-.65-.113ZM14 18.55V6.85l-4-1.4v11.7l4 1.4Zm2 0l3-1V5.7l-3 1.15v11.7ZM5 18.3l3-1.15V5.45l-3 1V18.3ZM16 6.85v11.7v-11.7Zm-8-1.4v11.7v-11.7Z"/>`),
		g.Group(children),
	)
}