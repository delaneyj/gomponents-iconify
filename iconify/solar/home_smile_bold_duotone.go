package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSmileBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12.204c0-2.289 0-3.433.52-4.381c.518-.949 1.467-1.537 3.364-2.715l2-1.241C9.889 2.622 10.892 2 12 2c1.108 0 2.11.622 4.116 1.867l2 1.241c1.897 1.178 2.846 1.766 3.365 2.715c.519.948.519 2.092.519 4.38v1.522c0 3.9 0 5.851-1.172 7.063C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.212C2 19.576 2 17.626 2 13.725v-1.521Z" opacity=".5"/><path d="M9.447 15.398a.75.75 0 0 0-.894 1.205A5.766 5.766 0 0 0 12 17.75a5.766 5.766 0 0 0 3.447-1.147a.75.75 0 0 0-.894-1.206A4.266 4.266 0 0 1 12 16.25a4.266 4.266 0 0 1-2.553-.852Z"/></g>`),
		g.Group(children),
	)
}