package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12.088 4.82a10 10 0 0 1 9.412.314a1 1 0 0 1 .493.748L22 6v13a1 1 0 0 1-1.5.866a8 8 0 0 0-8 0a1 1 0 0 1-1 0a8 8 0 0 0-7.733-.148l-.327.18l-.103.044l-.049.016l-.11.026l-.061.01L3 20h-.042l-.11-.012l-.077-.014l-.108-.032l-.126-.056l-.095-.056l-.089-.067l-.06-.056l-.073-.082l-.064-.089l-.022-.036l-.032-.06l-.044-.103l-.016-.049l-.026-.11l-.01-.061l-.004-.049L2 19V6a1 1 0 0 1 .5-.866a10 10 0 0 1 9.412-.314l.088.044l.088-.044z"/></g>`),
		g.Group(children),
	)
}