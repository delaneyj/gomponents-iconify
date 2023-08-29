package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyframesFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M8 4a2.599 2.599 0 0 0-2 .957l-4.355 5.24a2.847 2.847 0 0 0-.007 3.598l4.368 5.256C6.505 19.651 7.23 20 8 20a2.599 2.599 0 0 0 2-.957l4.355-5.24a2.847 2.847 0 0 0 .007-3.598L9.994 4.949A2.593 2.593 0 0 0 8 4zm8.382.214a1 1 0 0 1 1.32.074l.084.094l4.576 5.823c.808.993.848 2.396.13 3.419l-.12.158l-4.586 5.836a1 1 0 0 1-1.644-1.132l.072-.104l4.596-5.85a.845.845 0 0 0 .06-.978l-.07-.1l-4.586-5.836a1 1 0 0 1 .168-1.404z"/><path fill="currentColor" d="M12.382 4.214a1 1 0 0 1 1.32.074l.084.094l4.576 5.823c.808.993.848 2.396.13 3.419l-.12.158l-4.586 5.836a1 1 0 0 1-1.644-1.132l.072-.104l4.596-5.85a.845.845 0 0 0 .06-.978l-.07-.1l-4.586-5.836a1 1 0 0 1 .168-1.404z"/></g>`),
		g.Group(children),
	)
}