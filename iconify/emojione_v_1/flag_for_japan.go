package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForJapan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M64 43c0 6.075-3.373 11-10 11H10C3.373 54 0 49.075 0 43V21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v22"/><path fill="#ec1c24" d="M43.971 32.569c0 6.294-5.104 11.4-11.4 11.4c-6.292 0-11.395-5.105-11.395-11.4s5.103-11.398 11.395-11.398c6.295 0 11.4 5.103 11.4 11.398"/>`),
		g.Group(children),
	)
}