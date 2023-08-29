package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M381 128.6H132.1c-12.1 0-19.5 0-19.5 20.4v28.1h288V149c0-20.4-7.4-20.4-19.6-20.4z" fill="currentColor"/><path d="M333 96.5H180c-13.1 0-19.5.3-19.5 18.7h192c-.1-18.4-6.4-18.7-19.5-18.7z" fill="currentColor"/><path d="M432.4 169.6l-15.9-9.4v32.3h-321v-32.3l-15.2 9.4c-14.3 8.9-17.8 15.3-15 40.9l17.5 184.8c3.7 20.7 15.9 21.2 24 21.2h299.9c8.1 0 20.2-.5 23.9-21.2l17.2-184.4c2.3-24.4-2-32.8-15.4-41.3z" fill="currentColor"/>`),
		g.Group(children),
	)
}