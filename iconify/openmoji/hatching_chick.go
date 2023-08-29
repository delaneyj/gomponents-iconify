package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HatchingChick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" stroke="#FCEA2B" stroke-miterlimit="10" stroke-width="1.8" d="M36.109 44.124s21.337-.711 10.668-20.625c0 0-2.845-7.824-10.668-7.824c-2.208 0-4.019.51-5.497 1.273c-3.762 1.943-5.372 5.53-5.883 6.55c-.711 1.423-11.38 19.204 11.38 20.626z"/><path fill="#F1B31C" d="M35.372 30.603s5.925-.29 0 5.233c0 0-5.925-5.233 0-5.233z"/><path fill="#FFF" d="M61.694 37.95c0 14.387-11.663 26.05-26.05 26.05S9.594 52.337 9.594 37.95l7.443 3.721l6.513-4.651l8.373 5.582l6.202-5.582l8.176 4.736l7.02-5.667l8.373 1.861z"/><path fill="#D0CFCE" d="M53.321 36.09c-1.98 11.638-7.1 21.403-17.677 27.91c14.387 0 26.05-11.663 26.05-26.05l-8.373-1.86z"/><circle cx="30.908" cy="26.363" r="2.368"/><circle cx="40.381" cy="26.363" r="2.368"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m34.796 14.903l-1.363-5.117m3.395 5.92v-3.552m2.171 3.749l1.382-6.117m-5.009 20.817s5.925-.29 0 5.233c0 0-5.925-5.233 0-5.233z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M49.764 33.42c.028-2.618-.818-5.873-2.987-9.921c0 0-2.845-7.824-10.668-7.824c-2.208 0-4.019.51-5.497 1.273c-3.762 1.943-5.372 5.53-5.883 6.55c-.348.697-3.078 5.307-3.113 10.002m40.078 4.45c0 14.387-11.663 26.05-26.05 26.05S9.594 52.337 9.594 37.95l7.443 3.721l6.513-4.651l8.373 5.582l6.202-5.582l8.176 4.736l7.02-5.667l8.373 1.861z"/>`),
		g.Group(children),
	)
}