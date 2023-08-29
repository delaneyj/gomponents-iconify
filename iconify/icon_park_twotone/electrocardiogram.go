package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electrocardiogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectrocardiogram0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" d="M11 28.132h5.684L21.224 13l3.671 22l4.553-10.383l3.465 3.515H37"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectrocardiogram0)"/>`),
		g.Group(children),
	)
}