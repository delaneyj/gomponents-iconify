package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForNorthKorea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#024fa2" d="M54 11H10c-5.491 0-8.74 3.385-9.695 8h63.39c-.956-4.615-4.205-8-9.696-8M10 55h44c5.881 0 9.193-3.881 9.865-9H.135c.672 5.119 3.984 9 9.865 9"/><path fill="#ed1c27" d="M0 22v22h64V22c0-.338-.015-.67-.035-1H.035c-.02.33-.035.662-.035 1"/><path fill="#e6e7e8" d="M0 44c0 .684.049 1.351.135 2h63.73c.086-.649.135-1.316.135-2H0m63.965-23a14.188 14.188 0 0 0-.27-2H.305a14.208 14.208 0 0 0-.27 2h63.93"/><circle cx="19.393" cy="32.393" r="8.893" fill="#e6e7e8"/><path fill="#ed1c27" d="m26.803 30.08l-5.693.01l-1.765-5.757l-1.754 5.757l-5.696-.01l4.615 3.51l-1.791 5.72l4.635-3.554l4.635 3.554l-1.789-5.72z"/>`),
		g.Group(children),
	)
}