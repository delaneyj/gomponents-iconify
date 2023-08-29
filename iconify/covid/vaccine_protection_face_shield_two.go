package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionFaceShieldTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.167 17.107V15.06H19.7a1.02 1.02 0 0 0 .935-1.435c-1-2.239-2.476-4.994-2.476-4.994V6.45a2.024 2.024 0 0 0-.18-.847c-1.111-2.409-2.617-4.074-5.5-4.652C6.611-.224.807 3.887.763 9.871a9.19 9.19 0 0 0 3.071 6.937v6.442m10.238 0v-1.34M.777 9.433h10.418"/><path d="M11.2 4.628h8.042a4 4 0 0 1 4 4v11.577H18.2a7 7 0 0 1-7-7V4.628Zm-.005 0H2.528"/></g>`),
		g.Group(children),
	)
}