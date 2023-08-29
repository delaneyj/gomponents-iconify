package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reason(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 512H0V0h512v512zm-198-48h155v-38H362.5v-39h93v-39h-93v-38.5H466v-40H314V464zm-203.333 1h49.666v-55.38h38.535L228 465h57l-41.333-65c35.351-14.895 45.084-62.175 28.506-95.263C261.638 283.712 239.078 269.904 208 269h-97.667l.334 196zM206.5 308.632c35.152 2.931 36.077 62.136-7.864 62.136H160.71v-62.136h45.79z"/>`),
		g.Group(children),
	)
}