package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artemis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.567 16.059l8.865-1.77l-1.463 8.93c.831 1.149 1.394 2.488 1.689 4.016l3.15-16.44l-16.313 3.72c1.535.241 2.892.756 4.072 1.544Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.117 28.12l2.715-11.159l-11.06 3.098c4.037 1.388 6.818 4.075 8.345 8.061ZM8.398 40.208c-.018-1.077 3.292-6.15 9.93-15.22c2.005-.609 3.774-.792 5.309-.551c.295 1.522.19 3.277-.316 5.263c-8.872 6.987-13.846 10.49-14.923 10.508Zm26.574-28.994C23.953 2.012 10.772 5.921 5.5 13.162c6.68-3.765 14.273-4.805 21.812-.2m9.168-.45c9.943 11.822 6.097 24.148-.945 29.655c3.693-6.933 4.103-15.828-.563-21.789"/>`),
		g.Group(children),
	)
}