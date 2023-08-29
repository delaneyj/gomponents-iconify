package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Captainamerica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-960q-91 0-174 35.5T195 195T99.5 338T64 512t35.5 174T195 829t143 95.5T512 960t174-35.5T829 829t95.5-143T960 512t-35.5-174T829 195T686 99.5T512 64zm188 707l-65-219l181-139q16 49 16 99q0 78-35.5 146T700 771zM588 407l-76-215q105 0 189 62t115 159zm-380 6q31-97 115-159t189-62l-76 215zm181 139l-65 219q-61-45-96.5-113T192 512q0-50 16-99zm123 89l188 130q-84 61-188 61t-188-61z"/>`),
		g.Group(children),
	)
}