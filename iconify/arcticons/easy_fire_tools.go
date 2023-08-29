package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EasyFireTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="36.884" height="24.954" x="5.558" y="8.275" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.906 33.228c-.248 2.407-.138 4.814-1.624 6.497h13.436c-1.753-1.873-1.863-4.195-2.185-6.497m3.381-18.007v11.225a1.887 1.887 0 0 0 1.887 1.887h.565m-8.636.001V15.88a2.641 2.641 0 0 1 2.641-2.641h0a5.15 5.15 0 0 1 .855.066m-5.707 5.029h10.376m-13.28 8.114a3.652 3.652 0 0 1-3.208 1.886a3.785 3.785 0 0 1-3.773-3.773v-2.453a3.774 3.774 0 1 1 7.547 0v1.32h-7.547"/>`),
		g.Group(children),
	)
}