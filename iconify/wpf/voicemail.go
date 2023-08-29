package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M6 7c-3.302 0-6 2.698-6 6s2.698 6 6 6h14c3.302 0 6-2.698 6-6s-2.698-6-6-6s-6 2.698-6 6a5.98 5.98 0 0 0 1.531 4H10.47A5.98 5.98 0 0 0 12 13c0-3.302-2.698-6-6-6zm0 2c2.221 0 4 1.779 4 4s-1.779 4-4 4s-4-1.779-4-4s1.779-4 4-4zm14 0c2.221 0 4 1.779 4 4s-1.779 4-4 4s-4-1.779-4-4s1.779-4 4-4z"/>`),
		g.Group(children),
	)
}