package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandHolding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.516 18a6.428 6.428 0 0 0-1.817.266l-.02.007l-5.671 2.176l1.984 5.57l4.93-1.89l7.137 3.93l12.324-5.106l-.766-1.844l-11.437 4.735l-7.102-3.91l-3.89 1.488l-.641-1.807l3.797-1.457A4.023 4.023 0 0 1 10.516 20c.703 0 1.522.156 2.222.79l.014.007l.004.004c1.03.895 1.808 1.52 2.89 1.86c1.082.34 2.31.378 4.36.37l-.01-2c-2.012.008-3.063-.063-3.75-.281c-.688-.211-1.176-.59-2.168-1.45l-.012-.007c-1.157-1.039-2.531-1.297-3.55-1.293z"/>`),
		g.Group(children),
	)
}