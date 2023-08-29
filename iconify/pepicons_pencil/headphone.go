package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M17.405 14.911a.5.5 0 1 1-.949-.315c.358-1.076.544-2.236.544-3.43c0-4.725-2.972-7.912-7-7.912c-4.029 0-7 3.187-7 7.912a10.9 10.9 0 0 0 .544 3.43a.5.5 0 1 1-.95.315A11.9 11.9 0 0 1 2 11.166c0-5.262 3.403-8.912 8-8.912c4.597 0 8 3.65 8 8.912c0 1.3-.203 2.567-.595 3.745Z"/><path d="M5.977 17.034a3 3 0 0 1-2.942-3.04v-.022a2.978 2.978 0 0 1 3.035-2.937a1 1 0 0 1 .98 1.013l-.054 4a1 1 0 0 1-1.019.986ZM14.089 11a3 3 0 0 1 2.942 3.04v.022A2.978 2.978 0 0 1 14.013 17h-.016a1 1 0 0 1-.981-1.014l.054-4A1 1 0 0 1 14.089 11Z"/></g>`),
		g.Group(children),
	)
}