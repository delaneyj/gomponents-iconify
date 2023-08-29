package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M152 138V48a32 32 0 0 0-64 0v90a56 56 0 1 0 64 0Zm-32 70a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z" opacity=".2"/><path d="M212 56a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm-84 57V88a8 8 0 0 0-16 0v65a32 32 0 1 0 16 0Zm-8 47a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm40-66V48a40 40 0 0 0-80 0v86a64 64 0 1 0 80 0Zm-40 98a48 48 0 0 1-27.42-87.4A8 8 0 0 0 96 138V48a24 24 0 0 1 48 0v90a8 8 0 0 0 3.42 6.56A48 48 0 0 1 120 232Z"/></g>`),
		g.Group(children),
	)
}