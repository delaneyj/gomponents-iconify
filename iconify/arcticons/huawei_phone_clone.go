package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiPhoneClone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.77 4.5a2.503 2.503 0 0 0-2.51 2.509v24.465a2.503 2.503 0 0 0 2.509 2.51H20.59v7.007a2.503 2.503 0 0 0 2.509 2.509h15.132a2.504 2.504 0 0 0 2.509-2.509V16.525a2.503 2.503 0 0 0-2.51-2.508H27.416V7.009a2.504 2.504 0 0 0-2.51-2.509Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.589 33.983l-.001-17.458a2.503 2.503 0 0 1 2.509-2.509h4.317M25.28 38.888l10.182-.054m-22.23-26.603l17.113 19.246m.533-5.971l-.533 5.971m-5.704.586l5.704-.586"/>`),
		g.Group(children),
	)
}