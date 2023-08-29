package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimalisticMagniferZoomOutBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.157 20.313a9.157 9.157 0 1 0 0-18.313a9.157 9.157 0 0 0 0 18.313Z" opacity=".5"/><path fill-rule="evenodd" d="M8.023 11.156c0-.399.324-.722.723-.722h4.82a.723.723 0 1 1 0 1.445h-4.82a.723.723 0 0 1-.723-.723Zm10.816 7.683a.723.723 0 0 1 1.022 0l1.928 1.927a.723.723 0 0 1-1.023 1.023L18.84 19.86a.723.723 0 0 1 0-1.022Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}