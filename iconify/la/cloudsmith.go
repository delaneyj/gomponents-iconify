package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudsmith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.854 3c-2.41 0-4.563 1.95-4.563 4.35c0 2.41-1.531 3.941-3.941 3.941c-2.4 0-4.35 2.152-4.35 4.563C6 18.264 7.95 20 10.35 20c2.41 0 4.357-1.734 4.357-4.145c0-2.41 1.737-4.146 4.147-4.146c2.4 0 4.146-1.95 4.146-4.36C23 4.952 21.254 3 18.854 3zm.652 17a4.501 4.501 0 1 0-.01 9.002a4.501 4.501 0 0 0 .01-9.002z"/>`),
		g.Group(children),
	)
}