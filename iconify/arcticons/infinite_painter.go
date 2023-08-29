package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfinitePainter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.571 31.833c.417 0 .833-.277 1.25-.555c1.665-.971 3.053-2.36 4.302-3.747c5.829-5.83 10.964-12.353 15.822-19.014c.139-.278.416-.694.555-.972c0-.139 0-.555-.139-.694s-.416-.139-.694-.139c-.139 0-.416.278-.694.417c-1.943 1.388-4.025 2.775-5.968 4.163c-4.996 3.609-9.993 7.356-14.434 11.659c-1.249 1.249-2.36 2.637-3.47 3.886c-.416.416-.555.971-.139 1.527c.694 1.11 1.666 2.22 2.776 3.053c.278.278.417.278.833.416Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.685 27.947c-.416-.139-1.527-.139-1.943 0c-1.527.278-2.776 1.25-3.886 2.36c-.833.971-1.388 2.081-1.943 3.33c-.555 1.666-1.388 3.054-2.915 4.025c-.971.417-1.665.972-2.498 1.527l1.25.833c2.636 1.249 5.55 1.665 8.326.833c3.748-.972 6.524-5.274 7.356-9.16"/>`),
		g.Group(children),
	)
}