package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPharmacyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.625 21q-.675 0-1.15-.475T3 19.375q0-.15.037-.363t.088-.362L5 13L3.125 7.35q-.05-.15-.088-.35T3 6.625q0-.675.475-1.15T4.625 5H15.7l1.025-2.825q.175-.475.638-.712t.962-.038q.5.2.713.65t.037.95L18.35 5h1.025q.675 0 1.15.475T21 6.625q0 .15-.038.363t-.087.362L19 13l1.875 5.65q.05.15.088.35t.037.375q0 .675-.475 1.15t-1.15.475H4.625ZM12 17q.425 0 .713-.288T13 16v-2h2q.425 0 .713-.288T16 13q0-.425-.288-.713T15 12h-2v-2q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10v2H9q-.425 0-.713.288T8 13q0 .425.288.713T9 14h2v2q0 .425.288.713T12 17Zm-6.9 2h13.8l-2-6l2-6H5.1l2 6l-2 6Zm6.9-6Z"/>`),
		g.Group(children),
	)
}