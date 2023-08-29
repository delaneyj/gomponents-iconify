package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChefHat0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M12 34h24v8H12z"/><path d="M29 34V20m-7 14v-8m-10-1v9h24v-9s5-3 5-8s-4-7-9-7c0-2-3-6-8-6s-8 4-8 6c-4 0-9 2-9 7s5 8 5 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChefHat0)"/>`),
		g.Group(children),
	)
}