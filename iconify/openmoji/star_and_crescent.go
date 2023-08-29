package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarAndCrescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiStarAndCrescent0" d="M45 57c1.975 0 3.9-.275 5.713-.8A24.787 24.787 0 0 1 36 61c-13.812 0-25-11.188-25-25s11.188-25 25-25c5.5 0 10.588 1.775 14.713 4.8A20.544 20.544 0 0 0 45 15c-11.6 0-21 9.4-21 21s9.4 21 21 21z"/><path id="openmojiStarAndCrescent1" d="m50.652 27.896l2.372 4.807l5.306.771l-3.839 3.742l.906 5.284l-4.745-2.495l-4.745 2.495l.906-5.284l-3.839-3.742l5.305-.771z"/></defs><g fill="#B399C8" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiStarAndCrescent0"/><use href="#openmojiStarAndCrescent1"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiStarAndCrescent0"/><use href="#openmojiStarAndCrescent1"/></g>`),
		g.Group(children),
	)
}