package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity="0" d="M9 7L12 5L15 7L21 15L18 19H6L3 15L9 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`),
		g.Group(children),
	)
}