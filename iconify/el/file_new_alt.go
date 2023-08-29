package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileNewAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM448.975 221.924H814.6v244.409h-61.45V283.008H493.799v110.083H378.076v470.581h293.408v61.084H316.553V349.731l132.422-127.807zM725.83 519.873h113.013v145.898h145.972v113.086H838.843v145.898H725.83V778.857H579.932V665.771H725.83V519.873z"/>`),
		g.Group(children),
	)
}