package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThaiPbs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.388 20.999C25.784 27.379 20.858 35.26 27.245 43c-1.99-4.293-.701-14.25 7.143-22.001Zm5.223-9.71c-7.24 3.326-28.338 16.215-16.258 28.984c-5.8-10.51 6.268-21.719 16.258-28.985ZM40.956 5c-10.773 7.98-27.239 16.12-23.098 30.925c-10.37-10.39 7.59-24.864 23.098-30.924V5ZM12.63 29.971c.124-3.28.02-6.644 2.667-8.994c-3.624-1.76-6.052-2.227-8.254.416c5.905 1.142 4.044 5.502 5.586 8.578h.001Z"/>`),
		g.Group(children),
	)
}