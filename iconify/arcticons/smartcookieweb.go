package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartcookieweb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.53a42.639 42.639 0 0 1-17.233 7.695s.133 5.788.488 8.699a30.683 30.683 0 0 0 5.497 13.678C15.613 38.424 22.31 43.47 24 43.47s8.387-5.046 11.248-8.868a30.688 30.688 0 0 0 5.497-13.678c.355-2.91.488-8.699.488-8.699A42.636 42.636 0 0 1 24 4.53Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 10.523a29.514 29.514 0 0 1-11.928 5.327s.092 4.007.338 6.021a21.239 21.239 0 0 0 3.804 9.468c1.981 2.645 6.616 6.138 7.786 6.138s5.805-3.493 7.786-6.139a21.242 21.242 0 0 0 3.804-9.467c.246-2.015.338-6.021.338-6.021A29.511 29.511 0 0 1 24 10.523Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.134 23.266a7.28 7.28 0 0 1 .037.734A7.171 7.171 0 1 1 24 16.828c.247 0 .492.013.785.024a2.258 2.258 0 0 0-.506 1.502a2.094 2.094 0 0 0 .681 1.442a2.182 2.182 0 0 0-.558 1.533a2.11 2.11 0 0 0 2.192 2.024a2.182 2.182 0 0 0 1.484-.678a2.064 2.064 0 0 0 2.953-.116c.043.229.079.466.104.707Z"/><circle cx="21.28" cy="21.201" r=".75" fill="currentColor"/><circle cx="27.322" cy="26.696" r=".75" fill="currentColor"/><circle cx="21.21" cy="27.159" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}