package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricityCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopElectricityCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M18 11h-3.055l1.974-4.606A1 1 0 0 0 16 5h-5a1 1 0 0 0-.92.606l-3 7A1 1 0 0 0 8 14h1.734l-2.662 6.627c-.4.995.835 1.836 1.614 1.1l5.024-4.742l4.94-4.225c.706-.604.279-1.76-.65-1.76Zm-5.49.606A1 1 0 0 0 13.428 13h1.864l-2.92 2.498l-1.748 1.65l1.517-3.775A1 1 0 0 0 11.214 12H9.517l2.142-5h2.824l-1.974 4.606Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopElectricityCircleFilled0)"/></g>`),
		g.Group(children),
	)
}