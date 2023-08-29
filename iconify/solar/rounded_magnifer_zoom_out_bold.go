package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedMagniferZoomOutBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.157 20.313a9.157 9.157 0 1 0 0-18.313a9.157 9.157 0 0 0 0 18.313Zm8.704-1.475a.723.723 0 0 0-1.023 1.023l1.928 1.927a.723.723 0 0 0 1.022-1.022l-1.927-1.928ZM8.747 10.434a.723.723 0 1 0 0 1.445h4.82a.723.723 0 0 0 0-1.445h-4.82Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}