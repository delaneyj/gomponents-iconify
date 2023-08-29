package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.544 25.182l2.638-2.638l2.638 2.638l-2.638 2.638zm3.508 5.641l4.771-4.771l2.121 2.12l-4.771 4.772zm8.508-1.035l-4.772 4.772l5.91 7.94l6.802-6.802l-7.94-5.91zm-6.004-9.211h-6.748v-3h6.748zm-6.748-5.283h6.748L29.992 5.5h-9.62l1.436 9.794zm-1.231 6.514v6.748h-3v-6.748zm-5.283 6.748v-6.748L5.5 20.372v9.62l9.794-1.436z"/>`),
		g.Group(children),
	)
}