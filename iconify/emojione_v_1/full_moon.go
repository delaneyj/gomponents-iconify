package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32.12" r="31.875" fill="#f5eb35"/><g fill="#e0cf35"><circle cx="29.32" cy="53.02" r="9.226"/><path d="M41.904 24.487a3.918 3.918 0 1 1-7.836-.002a3.918 3.918 0 0 1 7.836.002"/><circle cx="5.967" cy="36.54" r="3.845"/><circle cx="6.313" cy="18.917" r="2.195"/><path d="M20.967 19.656a3.433 3.433 0 1 1-6.866 0a3.433 3.433 0 0 1 6.866 0"/><circle cx="42.896" cy="11.07" r="4.835"/></g>`),
		g.Group(children),
	)
}