package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buddy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.7 5.307L12.945.253a1.892 1.892 0 0 0-1.891 0L2.299 5.306a1.892 1.892 0 0 0-.945 1.638v10.11c0 .675.36 1.3.945 1.637l8.756 5.056a1.892 1.892 0 0 0 1.89 0l8.756-5.055c.585-.338.945-.962.945-1.638V6.945c0-.675-.36-1.3-.945-1.638zm-7.45 7.753l-3.805 3.804l-1.351-1.351l3.804-3.805l-3.804-3.806l1.35-1.35l3.805 3.805l1.351 1.35l-1.35 1.353z"/>`),
		g.Group(children),
	)
}