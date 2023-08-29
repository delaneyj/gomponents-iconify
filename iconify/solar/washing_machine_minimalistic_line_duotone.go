package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 10c0-3.771 0-5.657 1.172-6.828C5.343 2 7.229 2 11 2h2c3.771 0 5.657 0 6.828 1.172C21 4.343 21 6.229 21 10v4c0 3.771 0 5.657-1.172 6.828C18.657 22 16.771 22 13 22h-2c-3.771 0-5.657 0-6.828-1.172C3 19.657 3 17.771 3 14v-4Z" opacity=".5"/><path d="M17 14a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/><path stroke-linecap="round" d="M8 6h8m-8.766 8.362c.855.428 1.833 1.159 3.49 1.457c2.362.426 2.126-1.648 4.488-1.223c.72.13 1.206.35 1.55.585"/></g>`),
		g.Group(children),
	)
}