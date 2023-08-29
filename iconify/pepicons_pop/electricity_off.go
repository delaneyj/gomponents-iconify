package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricityOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15 8h-3.055l1.974-4.606A1 1 0 0 0 13 2H8a1 1 0 0 0-.92.606l-3 7A1 1 0 0 0 5 11h1.734l-2.662 6.627c-.4.995.835 1.836 1.614 1.1l5.024-4.742l4.94-4.225C16.356 9.156 15.93 8 15 8Zm-5.49.606A1 1 0 0 0 10.428 10h1.864l-2.92 2.498l-1.748 1.65l1.517-3.775A1 1 0 0 0 8.214 9H6.517l2.142-5h2.824L9.51 8.606Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}