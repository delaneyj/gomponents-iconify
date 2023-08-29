package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M768 103L384 432L0 103V84h768v19zm0 464L514 384l254-220v403zM0 164l254 220L0 567V164zm384 333l105-92l279 201v4H0v-4l279-201z"/>`),
		g.Group(children),
	)
}