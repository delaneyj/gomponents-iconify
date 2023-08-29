package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaningGibbousMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiWaningGibbousMoon0" d="M55 35a28.001 28.001 0 0 1-26.547 27.963A28 28 0 1 0 36 8q-.731 0-1.454.037A28.006 28.006 0 0 1 55 35Z"/></defs><use href="#openmojiWaningGibbousMoon0"/><circle cx="36" cy="36" r="28" fill="#fcea2b" stroke="#3f3f3f" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="#3f3f3f" d="M55 35a28.001 28.001 0 0 1-26.547 27.963A28 28 0 1 0 36 8q-.731 0-1.454.037A28.006 28.006 0 0 1 55 35Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiWaningGibbousMoon0"/></g>`),
		g.Group(children),
	)
}