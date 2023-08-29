package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareIosCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M12 6a1 1 0 0 1 2 0v10.5a1 1 0 1 1-2 0V6Z"/><path d="M13.64 5.768a1 1 0 1 0-1.28-1.536l-3 2.5a1 1 0 0 0 1.28 1.536l3-2.5Z"/><path d="M12.36 5.768a1 1 0 1 1 1.28-1.536l3 2.5a1 1 0 1 1-1.28 1.536l-3-2.5ZM16 12a1 1 0 1 1 0-2h1c1.623 0 3 1.165 3 2.692v7.616C20 21.835 18.623 23 17 23H9c-1.623 0-3-1.165-3-2.692v-7.616C6 11.165 7.377 10 9 10h1a1 1 0 0 1 0 2H9c-.586 0-1 .35-1 .692v7.616c0 .342.414.692 1 .692h8c.586 0 1-.35 1-.692v-7.616c0-.342-.414-.692-1-.692h-1Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}