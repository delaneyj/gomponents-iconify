package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchDuty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24.017 43.5l5.35-5.104s5.952-5.744 7.097-8.718c1.265-3.286.612-8.061-.886-11.858c-1.498-3.796-4.31-7.549-4.31-7.549s.073 1.932-.368 3.49c-.44 1.556-1.3 2.677-1.3 2.677s-1.07-3.916-3.301-6.9C24.067 6.554 20.676 4.5 20.676 4.5s.318 4.638-2.98 7.879c-5.966 5.864-6.556 8.68-6.821 13.475c-.328 5.913 7.96 12.542 7.96 12.542l5.182 5.104Z"/><path d="M24.063 38.396s5.75-4.351 5.75-7.942s-2.852-6.406-5.75-6.423c-2.899-.017-6.124 2.9-5.84 6.355c.282 3.456 5.84 8.01 5.84 8.01Z"/></g>`),
		g.Group(children),
	)
}