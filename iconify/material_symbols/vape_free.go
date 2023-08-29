package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VapeFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19q-.975-.25-1.988-.375T3 18.5H2v-2h1q1 0 2.013-.125T7 16v3Zm12.775 3.6l-3.6-3.6H8v-3h5.175L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4Zm2.05-3.6l-3-3H22v3h-.175ZM10.5 18q.2 0 .35-.15t.15-.35q0-.2-.15-.35T10.5 17q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm10-3v-2.225q0-1.7-1.15-2.888T16.5 8.7V7.2q.775 0 1.313-.537t.537-1.313q0-.775-.537-1.312T16.5 3.5V2q1.4 0 2.375.975t.975 2.375q0 .7-.263 1.313t-.737 1.062q1.425.675 2.288 2.013T22 12.75V15h-1.5ZM18 15v-1.3q0-.975-.588-1.513t-1.387-.537h-1.55L11.15 8.325V8.3q0-1.4.975-2.375T14.5 4.95v1.5q-.775 0-1.313.488T12.65 8.2q0 .775.537 1.388t1.313.612h1.525q1.425 0 2.45.9t1.025 2.25V15H18Z"/>`),
		g.Group(children),
	)
}