package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vrc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#418bca"/><path fill="#fff" fill-rule="nonzero" d="M20.265 8H25l-9 18L7 8h4.704l4.327 9.113zM16 12.959c-.788 0-1.427-.656-1.427-1.465s.639-1.466 1.427-1.466s1.427.657 1.427 1.466s-.64 1.465-1.427 1.465zm6.465 4.04c.788 0 1.427.657 1.426 1.466c0 .81-.638 1.465-1.426 1.465c-.788 0-1.427-.656-1.427-1.465S21.677 17 22.465 17z"/></g>`),
		g.Group(children),
	)
}