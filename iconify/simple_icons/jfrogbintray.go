package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jfrogbintray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.617 22.316h18.766V24H2.617zm15.88-12.632l-5.655 5.655V3.249l1.744 1.743L15.79 3.79L12 0L8.21 3.79l1.204 1.203l1.744-1.804v12.15L5.504 9.686H7.97V8H2.617v5.354H4.3v-2.527l7.7 7.699l7.698-7.699v2.527h1.685V8H16.03v1.684z"/>`),
		g.Group(children),
	)
}