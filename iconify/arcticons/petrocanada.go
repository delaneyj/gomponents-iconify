package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Petrocanada(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.585 29.785l5.934-.975l-1.314-5.382l4.365.635l.763-2.034l4.79 4.959l-1.611-9.748l2.797 1.4l2.543-5.172l2.925 5.382l2.457-1.61l-1.484 9.705l4.663-4.958l.848 2.076l4.153-.593l-1.271 5.382l6.146.89m-26.393 9.21h3.127m-3.127-6.242h3.127m-3.127 3.126h2.034m-2.034-3.126v6.242m-6.427.011V32.71H11.5a2.11 2.11 0 0 1 0 4.222H9.469m12.053-4.22h4.142m-2.031 6.251V32.71m4.337 6.253V32.71H30a2.11 2.11 0 0 1 .003 4.222H27.97m2.182-.007l1.958 2.038m2.28-2.033a2.071 2.071 0 1 0 4.142.018v-2.129a2.102 2.102 0 0 0-2.11-2.109a2.034 2.034 0 0 0-2.032 2.109Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.299 5.5H40.7a1.799 1.799 0 0 1 1.8 1.799V40.7a1.799 1.799 0 0 1-1.799 1.799H7.3a1.799 1.799 0 0 1-1.8-1.798V7.3a1.799 1.799 0 0 1 1.799-1.8Z"/>`),
		g.Group(children),
	)
}