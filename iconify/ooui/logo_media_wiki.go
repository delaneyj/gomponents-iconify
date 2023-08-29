package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoMediaWiki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g transform="translate(10 10)"><g id="oouiLogoMediaWiki0"><path id="oouiLogoMediaWiki1" fill="currentColor" d="M0 10c-2.9-3.3-.8-5.9 0-5.9S2.9 6.7 0 10z"/><use href="#oouiLogoMediaWiki1" transform="rotate(15)"/><use href="#oouiLogoMediaWiki1" transform="rotate(30)"/><use href="#oouiLogoMediaWiki1" transform="rotate(45)"/><use href="#oouiLogoMediaWiki1" transform="rotate(60)"/><use href="#oouiLogoMediaWiki1" transform="rotate(75)"/></g><use href="#oouiLogoMediaWiki0" transform="rotate(90)"/><use href="#oouiLogoMediaWiki0" transform="rotate(180)"/><use href="#oouiLogoMediaWiki0" transform="rotate(270)"/></g>`),
		g.Group(children),
	)
}