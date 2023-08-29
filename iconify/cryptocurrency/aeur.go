package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aeur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.048-13.014a7.144 7.144 0 0 0 .093 0a.055.055 0 0 1 .004.014a.049.049 0 1 1-.097-.014zm-6.014-3.493a6.993 6.993 0 0 0 6.055 10.493a7.002 7.002 0 0 0 6.993-6.993a6.96 6.96 0 0 0-.94-3.503a6.993 6.993 0 1 0-12.11.003zM16.048 12a7.138 7.138 0 0 0-.097 0a.049.049 0 0 1 .049-.05a.055.055 0 0 1 .049.05z"/>`),
		g.Group(children),
	)
}