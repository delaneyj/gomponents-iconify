package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EfaPublish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.222 26.14l-5.477-3.096a.647.647 0 0 1 .026-1.141l23.312-11.808l-4.02 23.413a.706.706 0 0 1-1.064.483l-8.146-4.975l-3.762 4.715a.596.596 0 0 1-1.061-.388Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.222 26.14l17.861-16.045l-13.23 18.921M38.836 40.73v-7.155m0 11.925c7.065-2.715 6.66-8.154 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c0 3.246-.405 8.685 6.66 11.4Zm-3.578-8.347h7.155"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.54 42.743a21.503 21.503 0 1 1 10.959-18.97m.001.173a21.46 21.46 0 0 1-1.265 7.338"/>`),
		g.Group(children),
	)
}