package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTrinidadAndTobago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M0 21v22c0 6.075 3.373 11 10 11h32.955L3.047 12.654C1.026 14.672 0 17.659 0 21zm64 22V21c0-6.075-3.373-11-10-11H20.728l40.068 41.509C62.924 49.491 64 46.43 64 43z"/><path fill="#e6e7e8" d="M58.963 52.824a8.669 8.669 0 0 0 1.833-1.315L20.726 10h-3.104l41.34 42.824M4.92 11.237a8.636 8.636 0 0 0-1.873 1.417l39.908 41.35h3.242L4.917 11.241"/><path fill="#25333a" d="M10 10c-1.99 0-3.68.451-5.08 1.237L46.197 54H54c1.936 0 3.585-.43 4.963-1.176L17.624 10H10z"/>`),
		g.Group(children),
	)
}