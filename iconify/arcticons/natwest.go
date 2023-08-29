package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Natwest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.892 17.029H13.634l5.184-9.018h10.258l-5.184 9.018zm-.597 9.017h5.781l-5.184-9.017H13.634l2.831 4.925m12.611 4.092l5.185-9.017l-5.185-9.018"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.758 30.971H5.5l5.184-9.017h10.259l-5.185 9.017zm0 0H5.5l5.184 9.018h10.259l-5.185-9.018zM24 27.272l-3.057-5.318m0 18.035L24 34.67m7.429-12.716h5.887l-5.184 9.017H21.873l2.832-4.925m7.427 4.925H21.873l5.184 9.018h10.259l-5.184-9.018zm5.184 9.018l5.184-9.018l-5.184-9.017"/>`),
		g.Group(children),
	)
}