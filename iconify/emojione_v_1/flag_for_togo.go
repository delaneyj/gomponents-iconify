package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#009543" d="M63.865 19c-.672-5.119-3.984-9-9.865-9H26v9h37.865"/><path fill="#f9cb38" d="M63.865 19H26v9h38v-7c0-.684-.049-1.351-.135-2"/><path fill="#009543" d="M26 28h38v8H26z"/><path fill="#f9cb38" d="M0 36v7c0 .684.049 1.351.135 2h63.73c.086-.649.135-1.316.135-2v-7H0"/><path fill="#009543" d="M10 54h44c5.881 0 9.193-3.881 9.865-9H.135c.672 5.119 3.984 9 9.865 9"/><path fill="#c32129" d="M10 10C4.119 10 .807 13.881.135 19A15.207 15.207 0 0 0 0 21v15h26V10H10"/><path fill="#e6e7e8" d="m21.574 20.393l-6.661.009l-2.066-6.738l-2.053 6.738l-6.666-.009l5.399 4.103l-2.093 6.694l5.42-4.16l5.424 4.16l-2.095-6.694z"/>`),
		g.Group(children),
	)
}