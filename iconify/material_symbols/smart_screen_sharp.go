package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartScreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h22v14H1Zm5-2h12V7H6v10Zm7.25-4.25q-.325 0-.537-.213T12.5 12q0-.325.213-.537t.537-.213q.325 0 .537.213T14 12q0 .325-.213.537t-.537.213Zm-5 0q-.325 0-.537-.213T7.5 12q0-.325.213-.537t.537-.213q.325 0 .537.213T9 12q0 .325-.213.537t-.537.213Zm7.5 0q-.325 0-.537-.213T15 12q0-.325.213-.537t.537-.213q.325 0 .537.213T16.5 12q0 .325-.213.537t-.537.213Zm-5 0q-.325 0-.537-.213T10 12q0-.325.213-.537t.537-.213q.325 0 .537.213T11.5 12q0 .325-.213.537t-.537.213Z"/>`),
		g.Group(children),
	)
}