package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlcoholNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsAlcoholNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18.5 4a1 1 0 0 1 1 1v12.199c0 .193.077.378.213.514A14.637 14.637 0 0 1 24 28.063v13.58A2.356 2.356 0 0 1 21.643 44h-9.435A2.208 2.208 0 0 1 10 41.792V28.408a14.39 14.39 0 0 1 4.763-10.695a.715.715 0 0 0 .237-.531V5a1 1 0 0 1 1-1h2.5ZM22 29H12v8h10v-8Zm6.999-11l-.638 6h7.278L35 18h-6Zm7.801-2h-9.6l-1.013 9.536A5.848 5.848 0 0 0 31 31.916V42h-3a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-3V31.915a5.848 5.848 0 0 0 4.813-6.379L36.8 16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAlcoholNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}