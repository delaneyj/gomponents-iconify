package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGuineaBissau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#00785e" d="M22 54h32c6.627 0 10-4.925 10-11V33H22v21z"/><path fill="#f9cb38" d="M54 10H22v23h42V21c0-6.075-3.373-11-10-11"/><path fill="#ec1c24" d="M22 10H10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h12V10z"/><path fill="#25333a" d="m19.06 29.357l-6.303.016l-1.95-6.379l-1.945 6.379l-6.308-.016l5.109 3.883l-1.976 6.334l5.124-3.936l5.132 3.936l-1.982-6.334z"/>`),
		g.Group(children),
	)
}