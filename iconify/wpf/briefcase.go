package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M11 1C8.802 1 7 2.803 7 5v1h2V5c0-1.115.884-2 2-2h4c1.116 0 2 .885 2 2v1h2V5c0-2.197-1.802-4-4-4h-4zM4.875 7A4.874 4.874 0 0 0 0 11.875v9.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125v-9.25A4.874 4.874 0 0 0 21.125 7H4.875zM13 13.906c1.161 0 2.094.933 2.094 2.094A2.088 2.088 0 0 1 13 18.094A2.088 2.088 0 0 1 10.906 16c0-1.161.933-2.094 2.094-2.094z"/>`),
		g.Group(children),
	)
}