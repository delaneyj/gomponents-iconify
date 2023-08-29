package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignFoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.798 14.262c0 3.231-2.339.253-5.294.253c-2.951 0-5.403 2.979-5.403-.253c0-1.509 1.232-3.097 2.09-4.453c.978-1.547 1.685-2.766 3.259-2.766c1.58 0 2.364 1.294 3.345 2.852c.849 1.351 2.003 2.865 2.003 4.367z"/><ellipse cx="5.91" cy="2.881" rx="1.91" ry="2.881"/><ellipse cx="10.936" cy="2.926" rx="1.936" ry="2.926"/><ellipse cx="1.871" cy="7.371" rx="1.885" ry="2.436" transform="rotate(-10.51 1.986 7.435)"/><path d="M12.115 7.305c-.345 1.326.201 2.504 1.214 2.627c1.014.126 2.116-.849 2.463-2.175c.345-1.328-.2-2.506-1.214-2.63c-1.016-.126-2.118.85-2.463 2.178z"/></g>`),
		g.Group(children),
	)
}