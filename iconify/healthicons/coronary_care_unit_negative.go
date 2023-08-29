package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoronaryCareUnitNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsCoronaryCareUnitNegative0)" clip-rule="evenodd"><path d="M48 0H0v48h48V0ZM9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Z"/><path d="M12.19 20.276c0-4.294 2.78-8.276 6.59-8.276c2.643 0 4.604 1.787 5.815 4.32c1.21-2.533 3.171-4.32 5.815-4.32c3.809 0 6.59 3.982 6.59 8.276C37 29.466 24.595 36 24.595 36s-8.265-4.09-11.303-10.805h3.774l1.674-2.927L20.696 29l4.58-5.79h4.685A1.03 1.03 0 0 0 31 22.187a1.03 1.03 0 0 0-1.039-1.022H24.26l-2.691 3.403l-2.2-7.569l-3.518 6.152h-3.287c-.24-.92-.375-1.879-.375-2.876Zm.364 2.876H10v2.043h3.285l-.73-2.043Z"/></g><defs><clipPath id="healthiconsCoronaryCareUnitNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}