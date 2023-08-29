package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RamenDiningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.975 22q-.425 0-.713-.288T7.975 21v-.75q-2.4-.95-3.987-2.813T2.05 13.15q-.05-.45.263-.8t.812-.35h.85V4.9q0-.375.263-.663t.637-.337L21.15 2.1q.35-.05.588.163t.237.562q0 .275-.188.488t-.462.262L10.475 4.8v1.7h10.75q.325 0 .537.213t.213.537q0 .325-.213.537T21.226 8h-10.75v4h10.35q.5 0 .813.35t.262.8q-.35 2.425-1.937 4.288t-3.988 2.812V21q0 .425-.287.713t-.713.287h-6Zm-1-15.5h1V4.95l-1 .1V6.5Zm-2.5 0h1V5.25l-1 .1V6.5Zm2.5 5.5h1V8h-1v4Zm-2.5 0h1V8h-1v4Z"/>`),
		g.Group(children),
	)
}