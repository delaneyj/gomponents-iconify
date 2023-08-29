package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cittamobi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.09 24.386v15.127H10.114V13.612a2.072 2.072 0 0 1 2.073-2.073h7.767m-6.331 28.039h3.551v2.886a1.036 1.036 0 0 1-1.036 1.036H14.66a1.036 1.036 0 0 1-1.036-1.036v-2.886h0Zm17.386 0h3.551v2.886a1.036 1.036 0 0 1-1.036 1.036h-1.48a1.036 1.036 0 0 1-1.035-1.036v-2.886h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.05 20.472h-6.405v10.109H34.56v-4.039m-20.915 8.57h5.185m10.545 0h5.185M15.464 15.897h3.627"/><circle cx="30.488" cy="15.897" r="11.397" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.666" cy="13.326" r="1.45" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.956 20.472c-1.016 1.481-2.197 2.876-4.468 2.876s-3.452-1.395-4.469-2.876m8.647-1.872s3.515-2.725 3.515-5.332a3.515 3.515 0 0 0-7.03 0c0 2.607 3.515 5.332 3.515 5.332"/><circle cx="26.31" cy="13.326" r="1.45" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}