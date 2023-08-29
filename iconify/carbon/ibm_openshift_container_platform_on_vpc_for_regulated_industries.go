package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmOpenshiftContainerPlatformOnVpcForRegulatedIndustries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="19" cy="27" r="1" fill="currentColor"/><path fill="currentColor" d="M29 31H16c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h13c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2zm-13-6v4h13v-4H16zm13-4H16c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h13c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2zm-13-6v4h13v-4H16z"/><path fill="currentColor" d="M12 27.3C7.347 25.648 4 21.213 4 16C4 9.383 9.383 4 16 4c4.831 0 8.994 2.876 10.895 7h2.166C27.042 5.746 21.957 2 16 2C8.28 2 2 8.28 2 16c0 6.33 4.225 11.685 10 13.41V27.3Z"/>`),
		g.Group(children),
	)
}