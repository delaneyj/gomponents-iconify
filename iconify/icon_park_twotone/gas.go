package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGas0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M23.049 9.789c1.626-.436 3.291.508 3.72 2.109c.429 1.6-.541 3.25-2.167 3.686C22.976 16.02 6.708 17 6.708 17s14.715-6.776 16.34-7.211Zm.058 28.658c1.625.435 3.291-.509 3.72-2.11c.429-1.6-.542-3.25-2.168-3.686c-1.625-.436-17.893-1.647-17.893-1.647s14.715 7.007 16.34 7.443Z"/><path stroke-linecap="round" d="M34 16.004a5 5 0 1 1 4 8H16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGas0)"/>`),
		g.Group(children),
	)
}