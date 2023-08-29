package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v1200h1200V0H0zm354.932 373.828c124.93 0 226.172 101.242 226.172 226.172S479.861 826.245 354.932 826.245S128.76 724.93 128.76 600s101.242-226.172 226.172-226.172zm490.136 0c124.93 0 226.172 101.242 226.172 226.172S969.998 826.245 845.068 826.245S618.896 724.93 618.896 600s101.243-226.172 226.172-226.172z"/>`),
		g.Group(children),
	)
}