package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 10.5h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm0-4h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm19.345 2.852a.666.666 0 0 0-.939-.086l-2.666 2.222a.665.665 0 0 0 0 1.023l2.666 2.223a.661.661 0 0 0 .938-.084a.667.667 0 0 0-.084-.94L19.708 12l2.052-1.71a.666.666 0 0 0 .085-.938zM2.5 14.5h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}