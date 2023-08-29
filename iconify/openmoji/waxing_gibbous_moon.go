package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaxingGibbousMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiWaxingGibbousMoon0" d="M17 37A28.001 28.001 0 0 1 43.547 9.037A28 28 0 1 0 36 64q.731 0 1.454-.037A28.006 28.006 0 0 1 17 37Z"/></defs><use href="#openmojiWaxingGibbousMoon0"/><circle cx="36" cy="36" r="28" fill="#fcea2b" stroke="#3f3f3f" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="#3f3f3f" d="M17 37A28.001 28.001 0 0 1 43.547 9.037A28 28 0 1 0 36 64q.731 0 1.454-.037A28.006 28.006 0 0 1 17 37Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiWaxingGibbousMoon0"/></g>`),
		g.Group(children),
	)
}