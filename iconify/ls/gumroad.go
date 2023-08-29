package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gumroad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M85 566V184c0-20 23-43 42-43h310c14 25 41 43 73 43c47 0 85-39 85-85c0-48-38-85-85-85c-32 0-59 17-73 42H127c-31 0-63 15-88 39c-25 25-39 57-39 89v382c0 32 14 64 39 89c25 24 57 39 88 39h341c31 0 63-15 88-39c25-25 39-57 39-89V354c0-35-14-68-40-92c-24-23-56-36-87-36H298c-34 0-66 14-91 40c-23 24-37 56-37 88v42c0 32 14 64 39 88c25 25 57 40 89 40h11c15 25 42 42 74 42c47 0 85-38 85-85s-38-85-85-85c-32 0-59 17-74 43h-11c-20 0-43-24-43-43v-42c0-20 19-43 43-43h170c19 0 42 19 42 43v212c0 19-23 43-42 43H127c-19 0-42-24-42-43z"/>`),
		g.Group(children),
	)
}