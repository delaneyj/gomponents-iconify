package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageResizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M32.79 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V14.062"/><path d="M42.892 3.492h-9.32a.565.565 0 0 0-.565.565v9.32c0 .312.253.565.565.565h9.32a.565.565 0 0 0 .565-.565v-9.32a.565.565 0 0 0-.565-.565ZM5.515 35.52s4.037-2.855 6.236-3.966c1.535-.775 3.098-1.996 4.812-1.86c1.85.146 3.05 2.216 4.813 2.8c.831.275 1.743.416 2.611.295c1.037-.143 2.02-.633 2.907-1.188c2.228-1.396 3.66-3.847 5.874-5.266c.791-.506 1.648-1.075 2.586-1.12c2.413-.113 4.519 1.987 6.892 2.24"/><circle cx="13.98" cy="13.253" r="3.966"/><path d="m36.412 8.21l1.818-1.818v4.66m1.822-2.853L38.234 6.38v4.66"/></g>`),
		g.Group(children),
	)
}