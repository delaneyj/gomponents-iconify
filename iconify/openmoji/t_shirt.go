package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61B2E4" d="m23.513 14.309l-2.69 1.712V56h30.354V16.021l-2.871-1.632l-3.207-.861l.633-2.404l-8.993-3.693l-9.302 2.208l-1.801 3.815zM61 31.329l-9.823-3V16.021l12.823 6z"/><path fill="#92D3F5" d="M45.732 11.124s2.451 2.43-4.707 11.283l-3.209-5.165l4.505-4.846l2.32-2.662M27.437 9.64l2.008 2.749l4.512 4.853l-3.209 5.165c-7.158-8.854-4.707-11.283-4.707-11.283"/><path fill="#61B2E4" d="m11 31.329l9.823-3V16.021L8 22.021z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m48.306 15.24l2.871.781V56H20.823V16.021l2.69-.861M61 31.329l-9.823-3V16.021l12.823 6zm-50 0l9.823-3V16.021L8 22.021z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m42.322 12.396l-4.506 4.846l3.209 5.165c7.158-8.854 4.707-11.283 4.707-11.283c-10.443-8.361-19.69 0-19.69 0s-2.452 2.43 4.706 11.283l3.21-5.165l-4.513-4.853"/><circle cx="35.887" cy="21.675" r=".853"/><circle cx="35.887" cy="24.942" r=".853"/><circle cx="35.887" cy="28.208" r=".853"/>`),
		g.Group(children),
	)
}