package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseHereButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="48.984" height="48.984" x="11.805" y="11.602" fill="#d0cfce" rx="1.699"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-miterlimit="5" d="M17.368 26.219h16.53a.43.43 0 0 1 .43.43v22.296m-.001-2.662H17.368m20.107-20.064h16.53a.43.43 0 0 1 .43.43v22.296m-.001-2.662H37.475"/><rect width="48.984" height="48.984" x="11.805" y="11.602" stroke-miterlimit="10" rx="1.699"/></g>`),
		g.Group(children),
	)
}