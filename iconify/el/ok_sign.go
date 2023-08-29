package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm257.446 281.03l124.657 124.658l-389.354 389.43L468.823 918.97L344.165 794.312l-126.27-126.344l123.853-123.853l126.27 126.343L857.446 281.03z"/>`),
		g.Group(children),
	)
}