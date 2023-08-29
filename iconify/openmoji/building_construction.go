package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingConstruction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiBuildingConstruction0" d="M58 25.833L12 27v-7l46 1.167z"/></defs><path fill="#FCEA2B" d="M58 25.833L12 27v-7l46 1.167z"/><path fill="#FCEA2B" d="M28 59H18l2-46h6z"/><path fill="#D0CFCE" d="M40.5 37h19v4h-19z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiBuildingConstruction0"/><path d="M28 59H18l2-46h6zm-9-9h8m-8-9h8m-7.5-7h7m7.5-8v-5m8 5v-5m8 16V21m-9.5 16h19v4h-19z"/><use href="#openmojiBuildingConstruction0"/></g>`),
		g.Group(children),
	)
}