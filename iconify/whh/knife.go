package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M990.824 989.5q-25.5 25.5-58.5 32t-70.5-9.5t-72.5-54l-224-263l-30 28q-13 12-31 12t-30-12q-63-57-120.5-117.5t-104.5-117t-87.5-111.5t-70.5-103.5t-51-92t-31-77t-9-57.5t13.5-35.5t31-11.5t30.5 11l372 343l511 434q38 34 54.5 71.5t10 71t-32 59z"/>`),
		g.Group(children),
	)
}