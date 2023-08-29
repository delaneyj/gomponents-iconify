package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func E(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M559 399H69v2c0 113 100 204 215 204c82 0 152-47 186-115l61 29c-44 90-137 152-246 152C133 671 0 551 0 401c0-42 11-82 29-117c47-91 147-153 256-153c130 0 238 89 267 208c4 19 7 40 7 60zM80 339h400c-27-82-104-142-196-142c-82 0-158 47-194 116c-4 8-7 17-10 26z"/>`),
		g.Group(children),
	)
}