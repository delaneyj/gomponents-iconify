package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoronaryCareUnitOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M18.78 12c-3.81 0-6.59 3.982-6.59 8.276c0 .997.133 1.957.374 2.877h3.287L19.37 17l2.199 7.569l2.69-3.403h5.702A1.03 1.03 0 0 1 31 22.188a1.03 1.03 0 0 1-1.039 1.021h-4.685L20.696 29l-1.956-6.732l-1.674 2.927h-3.774C16.33 31.911 24.595 36 24.595 36S37 29.467 37 20.276C37 15.983 34.219 12 30.41 12c-2.644 0-4.605 1.787-5.815 4.32C23.384 13.787 21.423 12 18.78 12Z"/><path fill-rule="evenodd" d="M6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm2 0a1 1 0 0 1 1-1h30a1 1 0 0 1 1 1v30a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V25.195h5.285l-.73-2.043H8V9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}