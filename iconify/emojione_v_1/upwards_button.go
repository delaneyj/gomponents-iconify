package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpwardsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.792 56.913a6.877 6.877 0 0 1-6.878 6.882H6.874A6.878 6.878 0 0 1 0 56.913V6.877A6.876 6.876 0 0 1 6.874 0h50.041a6.876 6.876 0 0 1 6.878 6.877v50.036z"/><path fill="#fff" d="M12.938 42.533c.194 2.843 1.437 5.216 3.137 6.257H50.01c1.704-1.045 2.95-3.426 3.139-6.272l-20.03-27.51l-20.18 27.525"/>`),
		g.Group(children),
	)
}