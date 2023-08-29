package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConvoxIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#4F8CE8" d="m2.36 152.515l124.927 65.694l125.91-65.171v29.895l-125.989 65.328L2.36 182.524v-30.009ZM253.197 98.68v29.887l-125.989 65.336L2.36 128.166v-28.99l124.918 64.683l125.92-65.18ZM127.025 0L256 72.258l-64.727 33.455l-64.222 33.762L0 72.188L127.025 0Zm.087 30.496L55.194 71.378l71.892 38.062l52.103-27.406l20.79-10.717l-72.867-40.82Z"/>`),
		g.Group(children),
	)
}