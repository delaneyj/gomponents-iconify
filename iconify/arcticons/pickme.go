package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pickme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 5.5h33c1.1 0 2 .9 2 2v33c0 1.1-.9 2-2 2H27.502l-.034-12.719c0-1.541.407-1.906 1.543-1.541l5.395 5.395c2.082 1.992 2.943-.592 2.313-1.541c-1.9-2.71-4.47-5.536-7.708-8.479c-2.475-1.524-4.416-.349-7.164-1.768c-2.75-1.42-5.672-6.852-6.712-9.024s-3.623-2.172-2.762.412c.862 2.583 2.443 8.695 5.845 11.92c.704 1.052.68 1.063.771 2.313V42.5H7.5c-1.1 0-2-.9-2-2v-33c0-1.1.9-2 2-2Z"/><circle cx="25.157" cy="15.135" r="3.854"/></g>`),
		g.Group(children),
	)
}