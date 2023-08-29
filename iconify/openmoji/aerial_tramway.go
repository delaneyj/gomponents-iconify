package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AerialTramway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="M33.429 22.514c.056.261-.474.898-.165.593a1.544 1.544 0 0 0 .43-.945a1.603 1.603 0 0 0-.807-1.61a1.5 1.5 0 1 0-1.514 2.59l.064.037l-.689-.896l.026.078v-.798l-.017.084l.385-.662a2.564 2.564 0 0 0-.64 1.032a2.824 2.824 0 0 0 .034 1.294a1.5 1.5 0 1 0 2.893-.797Zm8.568 2.912a1.5 1.5 0 0 0 0-3a1.5 1.5 0 0 0 0 3Zm-4.801-1.004a1.5 1.5 0 0 0 0-3a1.5 1.5 0 0 0 0 3Zm-10.969-4.388a1.61 1.61 0 0 0 1.567 2.701a1.515 1.515 0 0 0 1.048-1.845a2.72 2.72 0 0 0-1.08-1.438a1.51 1.51 0 0 0-2.053.538a1.535 1.535 0 0 0 .538 2.052a.632.632 0 0 1-.298-.355l1.047-1.845l-.094.025l1.156.15l-.08-.053l.742 1.295l-.003-.097l-.439 1.06l.07-.067a1.5 1.5 0 1 0-2.12-2.12Z"/><path fill="#fcea2b" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M27 30h15v23H27z"/><path fill="none" stroke="#61b2e4" stroke-linejoin="round" stroke-width="2" d="M28 50h14"/><path fill="#d0cfce" stroke="#d0cfce" stroke-linejoin="round" stroke-width="1.633" d="M27 42h6v4h-6z"/><path fill="#3f3f3f" stroke="#3f3f3f" stroke-linejoin="round" stroke-width="1.472" d="M27 34h6v6.5h-6z"/><path fill="#d0cfce" stroke="#d0cfce" stroke-linejoin="round" stroke-width="2" d="M36 34h6v11h-6z"/><g fill="none" stroke="#000" stroke-width="2"><rect width="4" height="4" x="30" y="20" stroke-miterlimit="10" rx="1.732" transform="rotate(9.671 32 22)"/><rect width="4" height="4" x="35" y="21" stroke-miterlimit="10" rx="1.732" transform="rotate(9.671 37 23)"/><rect width="4" height="4" x="40" y="22" stroke-miterlimit="10" rx="1.732" transform="rotate(9.667 41.998 24.003)"/><rect width="4" height="4" x="25" y="19" stroke-miterlimit="10" rx="1.732" transform="rotate(9.667 26.998 21.002)"/><path stroke-miterlimit="10" d="m12 18l48 10m-25-5v8"/><path stroke-linejoin="round" d="M36 34h6v11h-6zm-9 0h6v13h-6z"/><path stroke-linejoin="round" d="M27 30h15v23H27z"/></g>`),
		g.Group(children),
	)
}