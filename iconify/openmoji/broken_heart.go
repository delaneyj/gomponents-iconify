package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiBrokenHeart0" d="M49.566 15.61c-4.322 0-8.187 2.36-10.189 6.18L36 29.679l7 6.017l-6 8.424L42 51l-3.585 9.162l20.462-26.237a11.425 11.425 0 0 0 2.213-6.793c0-6.354-5.17-11.523-11.524-11.523z"/><path id="openmojiBrokenHeart1" d="m37.1 50.984l-5.054-6.961l5.939-8.046l-7.046-5.946s1.622-7.549 1.455-7.894c-1.944-4.026-5.918-6.528-10.371-6.528c-6.354 0-11.523 5.17-11.523 11.523c0 2.116.578 4.185 1.672 5.981l.66.974l20.462 26.23l.35-.45l3.455-8.883z"/></defs><g fill="#EA5A47"><use href="#openmojiBrokenHeart0"/><use href="#openmojiBrokenHeart1"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiBrokenHeart0" stroke-linecap="round" stroke-linejoin="round"/><use href="#openmojiBrokenHeart1"/></g>`),
		g.Group(children),
	)
}