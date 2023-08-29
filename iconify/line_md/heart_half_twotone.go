package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartHalfTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`),
		g.Group(children),
	)
}