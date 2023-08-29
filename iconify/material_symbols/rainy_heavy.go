package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainyHeavy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.45 18.9q-.375.2-.762.063T7.1 18.45l-6-12q-.2-.375-.062-.763T1.55 5.1q.375-.2.763-.062t.587.512l6 12q.2.375.063.763t-.513.587Zm4.675 0q-.375.2-.763.063t-.587-.513l-6-12q-.2-.375-.063-.763t.513-.587q.375-.2.75-.062t.575.512l6.025 12q.2.375.063.763t-.513.587Zm9.325 0q-.375.2-.762.063t-.588-.513l-6-12q-.2-.375-.062-.763t.512-.587q.375-.2.763-.062t.587.512l6 12q.2.375.063.763t-.513.587Zm-4.65-.025q-.375.2-.762.075t-.588-.5l-6-12q-.2-.375-.062-.763T10.9 5.1q.375-.2.75-.062t.575.512l6.025 12q.2.375.063.75t-.513.575Z"/>`),
		g.Group(children),
	)
}