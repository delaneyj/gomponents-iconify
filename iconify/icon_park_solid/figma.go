package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFigma0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path fill="#000" stroke="#000" d="M18 18a3 3 0 0 1 3-3h3v6h-3a3 3 0 0 1-3-3Zm0 6a3 3 0 0 1 3-3h3v6h-3a3 3 0 0 1-3-3Zm0 6a3 3 0 0 1 3-3h3v3a3 3 0 1 1-6 0Zm6-15h3a3 3 0 1 1 0 6h-3v-6Z"/><path fill="#000" stroke="#000" d="M24 24a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFigma0)"/>`),
		g.Group(children),
	)
}