package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimalisticMagniferZoomInBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.157 20.313a9.157 9.157 0 1 0 0-18.313a9.157 9.157 0 0 0 0 18.313Z" opacity=".5"/><path fill-rule="evenodd" d="M11.156 8.024c.4 0 .723.324.723.723v1.687h1.687a.723.723 0 1 1 0 1.446h-1.687v1.687a.723.723 0 0 1-1.446 0V11.88H8.746a.723.723 0 1 1 0-1.446h1.687V8.747c0-.399.324-.723.723-.723Zm7.683 10.815a.723.723 0 0 1 1.022 0l1.928 1.927a.723.723 0 0 1-1.023 1.023L18.84 19.86a.723.723 0 0 1 0-1.022Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}