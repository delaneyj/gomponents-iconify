package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adguard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C8.249 0 3.725.861 0 2.755C0 6.845-.051 17.037 12 24C24.051 17.037 24 6.845 24 2.755C20.275.861 15.751 0 12 0zm-.106 15.429L6.857 9.612c.331-.239 1.75-1.143 2.794.042l2.187 2.588c.009-.001 5.801-5.948 5.815-5.938c.246-.22.694-.503 1.204-.101l-6.963 9.226z"/>`),
		g.Group(children),
	)
}