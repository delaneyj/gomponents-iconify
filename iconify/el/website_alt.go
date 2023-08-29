package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebsiteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM281.154 301.634h637.653v596.732H281.154V301.634zm78.611 77.963V820.44h480.471V379.597H359.765zm52.393 56.231H786.55v195.479H412.158V435.828zm0 229.369h106.041v106.041H412.158V665.197zm134.156 0h106.078v106.041H546.314V665.197zm134.157 0H786.55v106.041H680.471V665.197z"/>`),
		g.Group(children),
	)
}