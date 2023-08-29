package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NinePointsConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNinePointsConnected0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m26.829 37.172l10.343-10.343m-16 0L10.829 37.172m16-16l10.343-10.343M10.829 21.172l10.343-10.343"/><path fill="#fff" d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0-16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNinePointsConnected0)"/>`),
		g.Group(children),
	)
}