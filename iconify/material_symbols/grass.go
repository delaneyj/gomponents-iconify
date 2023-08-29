package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-2h5.75q-.55-2.125-2.087-3.663T2 12.25q.5-.125.988-.188T4 12q3.35 0 5.675 2.325T12 20H2Zm12 0q0-1.05-.225-2.087t-.65-1.988q1.05-1.775 2.863-2.85T20 12q.525 0 1.012.063t.988.187q-2.125.55-3.65 2.087T16.25 18H22v2h-8Zm-2-5.975q0-1.625.6-3.05t1.65-2.513q1.05-1.087 2.463-1.737t3.012-.7q-1.4.875-2.45 2.15t-1.625 2.85q-1.1.525-2.012 1.288T12 14.024Zm-1.825-1.875q-.3-.225-.6-.425t-.625-.4q0-.15.025-.313T9 10.7q0-1.9-.6-3.6T6.7 4q1.65.675 2.863 1.937t1.862 2.913q-.45.75-.775 1.588t-.475 1.712Z"/>`),
		g.Group(children),
	)
}