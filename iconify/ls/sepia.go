package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sepia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M49.064 397c-60 42-24 95 26 98c50 4 92 38 80 79c-22 70-8 128 106 70c58-29 97-101 172 30c44 77 156 26 141-62c-7-43 17-69 101-94c153-45 88-172 12-204s-87-87-60-143c26-54-48-141-113-83c-52 46-117 4-118-12s-57-133-90-34c-22 63 13 172-181 59c-123-72-183 113-49 178c63 31 33 77-27 118z"/>`),
		g.Group(children),
	)
}