package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.814 20.738l2.25-1.238m0 0v9m5.9-4.5a2.229 2.229 0 0 1 2.223 2.222h0a2.229 2.229 0 0 1-2.223 2.222h-.889a3.615 3.615 0 0 1-2.777-.777m0-7.334a3.78 3.78 0 0 1 2.777-.777h.89a2.229 2.229 0 0 1 2.221 2.222h0A2.229 2.229 0 0 1 26.964 24h-2.222"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}