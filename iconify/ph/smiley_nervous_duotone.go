package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyNervousDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 128a96 96 0 1 1-96-96a96 96 0 0 1 96 96Z" opacity=".2"/><path d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm53.66-53.66a8 8 0 0 1-11.32 11.32L160 163.31l-10.34 10.35a8 8 0 0 1-11.32 0L128 163.31l-10.34 10.35a8 8 0 0 1-11.32 0L96 163.31l-10.34 10.35a8 8 0 0 1-11.32-11.32l16-16a8 8 0 0 1 11.32 0L112 156.69l10.34-10.35a8 8 0 0 1 11.32 0L144 156.69l10.34-10.35a8 8 0 0 1 11.32 0ZM80 108a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm72 0a12 12 0 1 1 12 12a12 12 0 0 1-12-12Z"/></g>`),
		g.Group(children),
	)
}