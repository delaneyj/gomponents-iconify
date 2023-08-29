package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSouthKorea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><circle cx="36" cy="36" r="9" fill="#d22f27"/><path fill="#1e50a0" d="M28.127 31.676A4.492 4.492 0 0 0 36 36c.023-.04.034-.083.055-.123l.024.014a4.493 4.493 0 0 1 7.724 4.59l.003.002a8.992 8.992 0 0 1-15.68-8.807Zm.204-.389l.02.011c-.03.046-.067.085-.095.133c.027-.047.047-.098.075-.144Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m24.232 41.902l3 5.196m-6.464-3.196l3 5.196M22.5 42.902l1 1.732m1 1.732l1 1.732m20 0l1-1.732m1-1.732l1-1.732m-1.268 6.196l1-1.732m1-1.732l1-1.732m-6.464 3.196l1-1.732m1-1.732l1-1.732m-26-13.804l3-5.196M22.5 29.098l3-5.196m-1.268 6.196l3-5.196m17.536 0l1 1.732m1 1.732l1 1.732m.464-7.196l1 1.732m1 1.732l1 1.732M46.5 23.902l3 5.196"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}