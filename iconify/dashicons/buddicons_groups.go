package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsGroups(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.45 6.25c1.83.94 1.98 3.18.7 4.98c-.8 1.12-2.33 1.88-3.46 1.78L10.05 18H9l-2.65-4.99c-1.13.16-2.73-.63-3.55-1.79c-1.28-1.8-1.13-4.04.71-4.97c.48-.24.96-.33 1.43-.31c-.01.4.01.8.07 1.21c.26 1.69 1.41 3.53 2.86 4.37c-.19.55-.49.99-.88 1.25L9 16.58v-5.66C7.64 10.55 6.26 8.76 6 7c-.4-2.65 1-5 3.5-5s3.9 2.35 3.5 5c-.26 1.76-1.64 3.55-3 3.92v5.77l2.07-3.84c-.44-.23-.77-.71-.99-1.3c1.48-.83 2.65-2.69 2.91-4.4c.06-.41.08-.82.07-1.22c.46-.01.92.08 1.39.32z"/>`),
		g.Group(children),
	)
}