package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckedOutByYouTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M939 171q129 0 248 33t224 95t190 147t147 190t95 224t34 249q0 129-33 248t-95 224t-147 191t-190 147t-224 95t-249 34q-129 0-249-33t-224-95t-190-147t-147-190t-95-225t-34-249q0-129 33-248t95-224t148-190t190-147t224-95t249-34zm426 1365V853h-170v361L673 693L522 843l522 522H683v171h682z"/>`),
		g.Group(children),
	)
}