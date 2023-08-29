package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm8.396 212.462c112.097-4.539 194.092 93.25 198.101 195.351H689.894c-4.339-45.75-36.726-79.855-81.497-78.062c-44.771 1.793-80.5 37.062-81.459 81.497v91.983l363.831-.002v366.644H309.194V503.229l100.457.002v-91.983c9.478-113.423 86.65-194.247 198.745-198.786z"/>`),
		g.Group(children),
	)
}