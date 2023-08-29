package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M20.405 17.911a.5.5 0 1 1-.949-.315c.358-1.076.544-2.236.544-3.43c0-4.725-2.971-7.912-7-7.912S6 9.44 6 14.166a10.9 10.9 0 0 0 .544 3.43a.5.5 0 1 1-.95.315A11.9 11.9 0 0 1 5 14.166c0-5.262 3.403-8.912 8-8.912c4.597 0 8 3.65 8 8.912c0 1.3-.203 2.567-.595 3.745Z"/><path d="M8.977 20.034a3 3 0 0 1-2.942-3.04v-.022a2.978 2.978 0 0 1 3.018-2.938h.017a1 1 0 0 1 .98 1.014l-.054 4a1 1 0 0 1-1.019.986ZM17.089 14a3 3 0 0 1 2.942 3.04v.022A2.978 2.978 0 0 1 17.013 20h-.016a1 1 0 0 1-.981-1.014l.054-4A1 1 0 0 1 17.089 14Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}