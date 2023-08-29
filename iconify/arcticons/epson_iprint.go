package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EpsonIprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.495 20.75v7.029l-5.887-.015a2.108 2.108 0 0 1-2.108-2.11V14.377c0-1.164.944-2.108 2.108-2.108h34.784c1.164 0 2.108.944 2.108 2.108v11.278a2.108 2.108 0 0 1-2.108 2.108l-6.278.022l-.019-7.035"/><path d="M30.635 22.628h-13.43a2.108 2.108 0 0 0-2.107 2.108v8.887c0 1.164.944 2.108 2.108 2.108h13.429a2.108 2.108 0 0 0 2.108-2.108v-8.887a2.108 2.108 0 0 0-2.108-2.108Zm-26.108-2.02h38.75m-25.452 5.355h12.087m-12.087 3.186h12.087m-12.087 3.186h12.087m-3.659-18.472h-4.665a2.108 2.108 0 0 0-2.108 2.108v.92c0 1.164.944 2.108 2.108 2.108h4.664a2.108 2.108 0 0 0 2.109-2.108v-.92a2.108 2.108 0 0 0-2.108-2.108Z"/></g>`),
		g.Group(children),
	)
}