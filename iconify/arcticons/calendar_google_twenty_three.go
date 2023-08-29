package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleTwentyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.419 22.5a2.966 2.966 0 0 1 2.889-3a2.986 2.986 0 0 1 2.11 5.111c-1.221 1-5 3.889-5 3.889h5.89m5.051-4.556a2.229 2.229 0 0 1 2.222 2.223h0a2.229 2.229 0 0 1-2.222 2.222h-.889a3.615 3.615 0 0 1-2.778-.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.692 20.278a3.78 3.78 0 0 1 2.778-.778h.889a2.229 2.229 0 0 1 2.222 2.222h0a2.229 2.229 0 0 1-2.222 2.222h-2.222"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}