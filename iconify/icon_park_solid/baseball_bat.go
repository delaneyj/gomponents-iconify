package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseballBat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBaseballBat0"><g fill="#fff" stroke="#fff" stroke-width="4"><circle cx="40" cy="40" r="3"/><path stroke-linejoin="round" d="M16.502 9.43S26.5 22 37.5 37.501C21.5 26 9.43 16.502 9.43 16.502S3.111 10.89 7 7c3.89-3.889 9.502 2.43 9.502 2.43Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBaseballBat0)"/>`),
		g.Group(children),
	)
}