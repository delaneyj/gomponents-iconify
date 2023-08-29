package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RattleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRattleOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="29.463" cy="16.989" r="11" fill="#fff" transform="rotate(40 29.463 16.989)"/><path d="M38.533 23.293s-5.636.493-10.998-4.006c-5.362-4.5-5.855-10.135-5.855-10.135m5.186 18.711c-3.46 1.012-5.068 2.928-6.675 4.843c-1.607 1.915-1.682 5.115-3.61 7.414c-1.928 2.298-4.746 2.544-7.044.616c-2.298-1.928-2.545-4.746-.617-7.044c1.929-2.298 5.068-2.928 6.675-4.843c1.607-1.915 3.214-3.83 3.61-7.414M38.832 10.49a3 3 0 0 0-4.596-3.856"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRattleOne0)"/>`),
		g.Group(children),
	)
}