package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm-95.801 261.841h191.602v242.358h242.358v191.602H695.801v242.358H504.199V695.801H261.841V504.199h242.358V261.841z"/>`),
		g.Group(children),
	)
}