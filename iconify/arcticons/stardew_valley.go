package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StardewValley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.368 9.926c-7.157.938-6.894 8.364-6.894 12.526c-4.087 2.465-4.568 4.81-4.568 7.635s3.066 9.137 9.017 9.137h17.403c4.208 0 9.347-9.717 9.347-13.054c0-3.727-2.825-3.958-2.825-3.958s3.246-1.383 3.246-4.088s-2.885-3.186-7.033-3.186s-5.771.962-7.815 6.132c0-2.628.315-8.37-4.223-10.46"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.705 26.84c1.082 1.743 1.984 2.946 3.908 2.946a4.901 4.901 0 0 0 4.147-2.946m-12.263-7.213v-5.05m-5.47 5.05H8.25m6.973-5.05v5.05m2.145-9.701a9.158 9.158 0 0 1 5.17-3.87c2.378-.7 4.203-.14 4.767.365a1.096 1.096 0 0 1-.412 1.886c-2.569.69-3.28.427-4.103 2.806m-9.833 27.895s-2.798 1.388-2.798 3.492h20.138a30.654 30.654 0 0 0-.195-3.276m-10.009 0V42.5m-9.619-20.048c.91-.609 3.186-.36 3.186-.36"/>`),
		g.Group(children),
	)
}