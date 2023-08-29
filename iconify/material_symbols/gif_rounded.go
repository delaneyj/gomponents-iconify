package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.25 15q-.325 0-.537-.213t-.213-.537v-4.5q0-.325.213-.537T12.25 9q.325 0 .537.213T13 9.75v4.5q0 .325-.213.537T12.25 15ZM6 15q-.45 0-.725-.313T5 14v-4q0-.375.275-.688T6 9h3.25q.325 0 .537.213T10 9.75q0 .325-.213.537t-.537.213H6.5v3h2v-.75q0-.325.213-.537T9.25 12q.325 0 .537.213t.213.537V14q0 .375-.275.688T9 15H6Zm9.25 0q-.325 0-.537-.213t-.213-.537v-4.5q0-.325.213-.537T15.25 9h3q.325 0 .537.213T19 9.75q0 .325-.213.537t-.537.213H16v1h1.25q.325 0 .537.213t.213.537q0 .325-.213.537T17.25 13H16v1.25q0 .325-.213.537T15.25 15Z"/>`),
		g.Group(children),
	)
}